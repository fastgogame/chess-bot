package main

import (
	"bufio"
	"fmt"
	"main/internal"
	"os"
	"strings"
)

const (
	exitCommand = "qq"
)

func main() {
	gameBoard := internal.NewBoard()
	gameBoard.DefaultSetup()
	gameBoard.TerminalPrint()

	reader := bufio.NewReader(os.Stdin)
	playGame(reader, gameBoard)
}

func playGame(reader *bufio.Reader, board *internal.Board) {
	for {
		fmt.Print("Enter your move: ")
		move, err := readMove(reader)
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		if move == exitCommand {
			fmt.Println("Exiting...")
			break
		}

		println(board.Move(move))
	}
}

func readMove(reader *bufio.Reader) (string, error) {
	s, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(s), nil
}
