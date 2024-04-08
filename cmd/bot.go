package main

import (
	"bufio"
	"fmt"
	"main/internal"
	"os"
	"strings"
)

func main() {
	b := new(internal.Board)
	b.Init()
	b.DefaultSetup()
	b.TerminalPrint()

	reader := bufio.NewReader(os.Stdin)
	//Move example: d2d4 or a1b2
	for {
		fmt.Print("Enter your move: ")
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}
		s = strings.TrimSpace(s)
		if s == "qq" {
			fmt.Println("Exiting...")
			break
		}
		println(b.Move(s))
	}
}
