package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	_ "github.com/casept/esponda/statik"

	"github.com/bwmarrin/discordgo"
	"github.com/rakyll/statik/fs"
)

func createAndSendSchillerMaymay(s *discordgo.Session, maymayText string, channelID string) {
	// TODO: Spawn a goroutine w/ channel to tell users that the bot is busy until we're done processing the image
	err := s.ChannelTyping(channelID)
	if err != nil {
		panic(err)
	}
	maymayPath := createMaymay(maymayText)
	maymayFile, err := os.Open(maymayPath)
	if err != nil {
		panic(err)
	}
	defer maymayFile.Close()
	defer os.Remove(maymayPath)
	file := discordgo.File{Name: "schillersay.png", ContentType: "image/png", Reader: maymayFile}
	data := discordgo.MessageSend{
		Content: "", Embed: nil, Tts: false, Files: []*discordgo.File{&file},
	}
	s.ChannelMessageSendComplex(channelID, &data)
}

// Call the script to generate an image in a temporary location and return the path
func createMaymay(maymayText string) string {
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

// Extracts the image template and font to a temporary directory and return the path to this directory.
func extractAssetsIfNeeded() string {
	// Create dir if it doesn't exist
	dirPath := path.Join(os.TempDir(), "schillersay-assets")
	_, err := os.Stat(dirPath)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("Creating directory and extracting assets")
		err = os.Mkdir(dirPath, 0700)
		if err != nil {
			panic(err)
		}

		// And extract assets into it
		statikFS, err := fs.New()
		if err != nil {
			panic(err)
		}
		files := []string{"SchillerBlank.png", "Raleway-ExtraBold.ttf"}
		for _, file := range files {
			r, err := statikFS.Open(strings.Join([]string{"/", file}, ""))
			if err != nil {
				panic(err)
			}
			defer r.Close()

			tempFile, err := os.Create(path.Join(dirPath, file))
			if err != nil {
				panic(err)
			}
			defer tempFile.Close()
			_, err = io.Copy(tempFile, r)
			if err != nil {
				panic(err)
			}
		}
	}

	return dirPath
}
