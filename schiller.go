package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	_ "github.com/casept/esponda/statik"

	"github.com/bwmarrin/discordgo"
)

func createAndSendSchillerMaymay(s *discordgo.Session, maymayText string, channelID string) {
	// TODO: Spawn a goroutine w/ channel to tell users that the bot is busy until we're done processing the image
	err := s.ChannelTyping(channelID)
	if err != nil {
		panic(err)
	}
	maymayPath := createSchillerMaymay(maymayText)
	defer os.Remove(maymayPath)
	sendPng(s, channelID, maymayPath)
}

// Call the script to generate an image in a temporary location and return the path
func createSchillerMaymay(maymayText string) string {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "schillersay-")
	defer tmpFile.Close()
	if err != nil {
		panic(err)
	}
	assetsPath := extractAssetsIfNeeded()

	cmd := exec.Command("convert", path.Join(assetsPath, "SchillerBlank.png"), "-size", "534x235", "-background", "none", "-font", path.Join(assetsPath, "Raleway-ExtraBold.ttf"), "-fill", "#042B66", strings.Join([]string{"caption:", maymayText}, ""), "-geometry", "+420+120", "-composite", tmpFile.Name())
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		panic(err)
	}
	return tmpFile.Name()
}
