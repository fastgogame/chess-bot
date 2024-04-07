package internal

import (
	"errors"
	"fmt"
)

func (b *Board) Move(commandMove string) string {
	err := b.parseMove(commandMove)
	if err != nil {
		return fmt.Sprintf("errror: %v", err)
	}
	b.TerminalPrint()
	b.switchSide()
	return "Move successful"
}

func (b *Board) switchSide() {
	b.WhiteTurn = !b.WhiteTurn
	if b.WhiteTurn {
		println("White's turn")
	} else {
		println("Black's turn")
	}
}

func isValidMove(r []rune) bool {
	if len(r) != 4 {
		return false
	}

	if r[0] < 'a' || r[0] > 'h' || r[2] < 'a' || r[2] > 'h' {
		return false
	}

	if r[1] < '1' || r[1] > '8' || r[3] < '1' || r[3] > '8' {
		return false
	}
	return true
}

func (b *Board) parseMove(commandMove string) (err error) {

	moveRunes := []rune(commandMove)
	if !isValidMove(moveRunes) {
		return errors.New("invalid move string")
	}

	fromFile := int(moveRunes[0] - 'a')
	fromRank := 7 - int(moveRunes[1]-'1')
	toFile := int(moveRunes[2] - 'a')
	toRank := 7 - int(moveRunes[3]-'1')

	fromSpot := &b.Spots[fromRank][fromFile]
	toSpot := &b.Spots[toRank][toFile]

	if fromSpot.Piece != nil {
		print("switching ... ")
		switch fromSpot.Piece.Icon {
		case '♟', '♙':
			println("pawn\n")
			b.movePawn(fromSpot, toSpot)
		default:
			return errors.New("invalid move switch")
		}

	}
	return nil
}

func (b *Board) movePawn(fs, ts *Spot) {
	ts.Piece = fs.Piece
	fs.Piece = nil
}
