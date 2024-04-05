package internal

import (
	"bufio"
	"fmt"
	"os"
)

func (b *Board) Move() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your move: ")
	move, _ := reader.ReadString('\n')

	fmt.Print(move)
}
