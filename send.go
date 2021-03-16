package main

import (
	"os"
	"path"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Sends the png at path to the chat.
func sendPng(s *discordgo.Session, channelID string, pngPath string) {
	file, err := os.Open(pngPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Join([]string{path.Base(pngPath), ".png"}, "")
	discordFile := discordgo.File{Name: name, ContentType: "image/png", Reader: file}
	data := discordgo.MessageSend{
		Content: "", Embed: nil, Tts: false, Files: []*discordgo.File{&discordFile},
	}
	s.ChannelMessageSendComplex(channelID, &data)
}

func createAndSendMaymay(s *discordgo.Session, maymayText string, channelID string, recipe ImageAnnotationRecipe) {
	// TODO: Spawn a goroutine w/ channel to tell users that the bot is busy until we're done processing the image
	err := s.ChannelTyping(channelID)
	if err != nil {
		panic(err)
	}
	// 2500 character limit
	if len(maymayText) > 2500 {
		s.ChannelMessageSend(channelID, "Please use no more than 2500 characters!")
	}
	maymayPath := recipe.Annotate()
	defer os.Remove(maymayPath)
	sendPng(s, channelID, maymayPath)
}
