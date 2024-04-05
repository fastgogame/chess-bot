package internal

import "fmt"

type Piece struct {
	Killed bool
	White  bool
	Icon   rune
}

type King struct {
	Piece
	Castled bool
	// ♚ ♔
}

type Pawn struct {
	Piece
	Moved bool
	// ♙ ♟
}

type Rook struct {
	Piece
	Moved bool
	// ♖ ♜
}

type Knight struct {
	Piece
	// ♘ ♞
}

type Bishop struct {
	Piece
	// ♗ ♝
}

type Queen struct {
	Piece
	// ♕ ♛
}

type Spot struct {
	Piece      *Piece
	file, rank int
	Icon       rune
	// □ ■
}

const BoardSize = 8

type Board struct {
	Spots [BoardSize][BoardSize]Spot
}

func (b *Board) Init() {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			b.Spots[i][j].file = i
			b.Spots[i][j].rank = j
			if (i+j)%2 == 0 {
				b.Spots[i][j].Icon = '■'
			} else {
				b.Spots[i][j].Icon = '□'
			}
		}
	}
}

func (b *Board) TerminalPrint() {
	println("a b c d e f g h")

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if b.Spots[i][j].Piece != nil {
				print(string(b.Spots[i][j].Piece.Icon) + " ")
			} else {
				print(string(b.Spots[i][j].Icon) + " ")
			}
		}
		print(BoardSize - i)
		println()
	}
}

func (b *Board) addPiece(p *Piece, file, rank int) {
	b.Spots[file][rank].Piece = p
}

func (b *Board) DefaultSetup() {
	fmt.Println("Arranging pieces")
	kingW := King{Piece{White: true, Icon: '♚'}, false}
	b.addPiece(&kingW.Piece, 7, 4)
	kingB := King{Piece{White: false, Icon: '♔'}, false}
	b.addPiece(&kingB.Piece, 0, 4)

	queenW := Queen{Piece{White: true, Icon: '♛'}}
	b.addPiece(&queenW.Piece, 7, 3)
	queenB := Queen{Piece{White: false, Icon: '♕'}}
	b.addPiece(&queenB.Piece, 0, 3)

	rookWL := Rook{Piece{White: true, Icon: '♜'}, false}
	b.addPiece(&rookWL.Piece, 7, 0)
	rookWR := Rook{Piece{White: true, Icon: '♜'}, false}
	b.addPiece(&rookWR.Piece, 7, 7)
	rookBL := Rook{Piece{White: false, Icon: '♖'}, false}
	b.addPiece(&rookBL.Piece, 0, 0)
	rookBR := Rook{Piece{White: false, Icon: '♖'}, false}
	b.addPiece(&rookBR.Piece, 0, 7)

	fmt.Println("Complete")
}
