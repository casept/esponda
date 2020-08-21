package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/jedib0t/go-pretty/table"
)

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

// Send a bingo board to the chat.
func sendBingo(s *discordgo.Session, channelID string) {
	b := newBingoBoard()
	s.ChannelMessageSend(channelID, "```\n"+b.boardToASCII()+"```\n")

}
