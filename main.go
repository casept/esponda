package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	if Token == "" {
		panic("You must provide a bot token!")
	}
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Bingo board
	if m.Content == "!bingo" {
		sendBingo(s, m.ChannelID)
	}

	// Schiller maymays
	if strings.HasPrefix(m.Content, "!schillersay ") {
		availableStyles := make(map[string]func(s *discordgo.Session, maymayText string, channelID string))
		availableStyles["schillerQuote"] = createAndSendSchillerQuoteMaymay
		availableStyles["schillerSpeech"] = createAndSendSchillerSpeechMaymay
		availableStyles["esponda"] = createAndSendEspondaMaymay

		withoutPrefix := strings.TrimPrefix(m.Content, "!schillersay ")
		styleAndText := strings.SplitN(withoutPrefix, " ", 2)
		style := styleAndText[0]
		if senderFunc, ok := availableStyles[style]; ok {
			text := styleAndText[1]
			fmt.Printf("Creating maymay in style '%s' with text '%s'\n", style, text)
			go senderFunc(s, text, m.ChannelID)
		} else {
			// Send a help message
			styles := []string{}
			for style := range availableStyles {
				styles = append(styles, style)
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Unknown format. Available formats are %s", strings.Join(styles, " ")))
		}

	}
}
