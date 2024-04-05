package main

import "main/internal"

func init() {
	// TODO
}

func main() {
	b := new(internal.Board)
	b.Init()
	b.TerminalPrint()
	//b.Move()
}
