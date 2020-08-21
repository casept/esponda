package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Send a meme with the result of the given haskell code
func createAndSendHaskellMaymay(s *discordgo.Session, code string, channelID string) {
	// TODO: Spawn a goroutine w/ channel to tell users that the bot is busy until we're done processing the image
	err := s.ChannelTyping(channelID)
	if err != nil {
		panic(err)
	}

	evalResult := evalHaskell(code)
	maymayPath := createEspondaMaymay(evalResult)
	sendPng(s, channelID, maymayPath)
	defer os.Remove(maymayPath)
}

// Call the script to generate an image in a temporary location and return the path
func createEspondaMaymay(text string) string {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "hascc-")
	defer tmpFile.Close()
	if err != nil {
		panic(err)
	}
	assetsPath := extractAssetsIfNeeded()

	cmd := exec.Command("convert", path.Join(assetsPath, "EspondaBlank.png"), "-size", "534x235", "-background", "none", "-font", path.Join(assetsPath, "Raleway-ExtraBold.ttf"), "-fill", "#042B66", strings.Join([]string{"caption:", text}, ""), "-geometry", "+420+120", "-composite", tmpFile.Name())
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		panic(err)
	}
	return tmpFile.Name()
}

// Evaluate the given Haskell code.
// FIXME: This is literally RCE, but for a meme project it doesn't matter that much
// TODO: Do this in a container or so
func evalHaskell(code string) string {
	cmd := exec.Command("/bin/sh", "echo", "'"+code+"'", "|", "ghci")
	out, _ := cmd.CombinedOutput()
	return string(out)
}
