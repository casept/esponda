package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/jedib0t/go-pretty/table"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

// BingoBoard represents the current bingo state
type BingoBoard struct {
	boardState [5][5]bool
	board      [5][5]string
}

// Create a new bingo board with default entries
func newBingoBoard() *BingoBoard {
	var b BingoBoard
	b.boardState[0] = [5]bool{false, false, false, false, false}
	b.boardState[1] = [5]bool{false, false, false, false, false}
	b.boardState[2] = [5]bool{false, false, false, false, false}
	b.boardState[3] = [5]bool{false, false, false, false, false}
	b.boardState[4] = [5]bool{false, false, false, false, false}

	b.board[0] = [5]string{"test0", "test1", "test2", "test3", "test4"}
	b.board[1] = [5]string{"test0", "test1", "test2", "test3", "test4"}
	b.board[2] = [5]string{"test0", "test1", "test2", "test3", "test4"}
	b.board[3] = [5]string{"test0", "test1", "test2", "test3", "test4"}
	b.board[4] = [5]string{"test0", "test1", "test2", "test3", "test4"}
	return &b
}

// Mark a bingo field off
func (b *BingoBoard) markField(x uint, y uint) {
	b.boardState[x][y] = true
}

// Converts the bingo board to ASCII art
func (b *BingoBoard) boardToASCII() string {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"1", "2", "3", "4", "5"})
	t.AppendRows([]table.Row{
		{b.boardState[0][0], b.boardState[0][1], b.boardState[0][2], b.boardState[0][3], b.boardState[0][4]},
	})
	//t.AppendSeparator()
	return t.RenderMarkdown()
}
func main() {

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

	b := newBingoBoard()
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "!bingo" {
		s.ChannelMessageSend(m.ChannelID, "```\n"+b.boardToASCII()+"```\n")
	}
}
