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
