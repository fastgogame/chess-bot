package internal

import (
	"errors"
	"fmt"
)

const (
	minFile = 'a'
	maxFile = 'h'
	minRank = '1'
	maxRank = '8'
)

func (b *Board) Move(commandMove string) string {
	err := b.parseMove(commandMove)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
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

func isValidCommandMove(r []rune) bool {
	if len(r) != 4 {
		return false
	}

	if r[0] < minFile || r[0] > maxFile || r[2] < minFile || r[2] > maxFile {
		return false
	}

	if r[1] < minRank || r[1] > maxRank || r[3] < minRank || r[3] > maxRank {
		return false
	}
	return true
}

func (b *Board) parseMove(commandMove string) (err error) {
	moveRunes := []rune(commandMove)
	if !isValidCommandMove(moveRunes) {
		return errors.New("invalid move string")
	}

	fromFile := int(moveRunes[0] - minFile)
	fromRank := 7 - int(moveRunes[1]-minRank)
	toFile := int(moveRunes[2] - minFile)
	toRank := 7 - int(moveRunes[3]-minRank)

	fromSpot := &b.Spots[fromRank][fromFile]
	toSpot := &b.Spots[toRank][toFile]

	if fromSpot.Piece != nil && fromSpot.Piece.White == b.WhiteTurn {
		println("Switching... ")
		switch fromSpot.Piece.Icon {
		case whitePawn, blackPawn:
			println("Pawn\n")
			b.movePawn(fromSpot, toSpot)
		default:
			return errors.New("invalid move switch")
		}
	} else {
		return errors.New("missing piece")
	}
	return nil
}

func (b *Board) movePawn(fromSpot, toSpot *Spot) {
	toSpot.Piece = fromSpot.Piece
	fromSpot.Piece = nil
}
