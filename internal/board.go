package internal

type Piece struct {
	Killed bool
	White  bool
	Icon   rune
	Moved  bool
}

type Spot struct {
	Piece      *Piece
	file, rank int
	Square     rune
	// □ ■
}

type Board struct {
	Spots     [8][8]Spot
	WhiteTurn bool
}

func (b *Board) Init() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b.Spots[i][j].rank = i
			b.Spots[i][j].file = j
			if (i+j)%2 == 0 {
				b.Spots[i][j].Square = '■'
			} else {
				b.Spots[i][j].Square = '□'
			}
		}
	}
}

func (b *Board) TerminalPrint() {
	println("a b c d e f g h")

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if b.Spots[i][j].Piece != nil {
				print(string(b.Spots[i][j].Piece.Icon) + " ")
			} else {
				print(string(b.Spots[i][j].Square) + " ")
			}
		}
		print(8 - i)
		println()
	}
}

func (b *Board) addPiece(p *Piece, file, rank int) {
	b.Spots[file][rank].Piece = p
}

func (b *Board) DefaultSetup() {
	b.WhiteTurn = true

	kingW := Piece{White: true, Icon: '♚'}
	b.addPiece(&kingW, 7, 4)
	kingB := Piece{White: false, Icon: '♔'}
	b.addPiece(&kingB, 0, 4)

	queenW := Piece{White: true, Icon: '♛'}
	b.addPiece(&queenW, 7, 3)
	queenB := Piece{White: false, Icon: '♕'}
	b.addPiece(&queenB, 0, 3)

	rookWL := Piece{White: true, Icon: '♜'}
	b.addPiece(&rookWL, 7, 0)
	rookWR := Piece{White: true, Icon: '♜'}
	b.addPiece(&rookWR, 7, 7)
	rookBL := Piece{White: false, Icon: '♖'}
	b.addPiece(&rookBL, 0, 0)
	rookBR := Piece{White: false, Icon: '♖'}
	b.addPiece(&rookBR, 0, 7)

	bishopWL := Piece{White: true, Icon: '♝'}
	b.addPiece(&bishopWL, 7, 2)
	bishopWR := Piece{White: true, Icon: '♝'}
	b.addPiece(&bishopWR, 7, 5)
	bishopBL := Piece{White: false, Icon: '♗'}
	b.addPiece(&bishopBL, 0, 2)
	bishopBR := Piece{White: false, Icon: '♗'}
	b.addPiece(&bishopBR, 0, 5)

	knightWL := Piece{White: true, Icon: '♞'}
	b.addPiece(&knightWL, 7, 1)
	knightWR := Piece{White: true, Icon: '♞'}
	b.addPiece(&knightWR, 7, 6)
	knightBL := Piece{White: false, Icon: '♘'}
	b.addPiece(&knightBL, 0, 1)
	knightBR := Piece{White: false, Icon: '♘'}
	b.addPiece(&knightBR, 0, 6)

	var pawnsW, pawnsB [8]Piece
	for i := 0; i < 8; i++ {
		pawnsW[i] = Piece{White: true, Icon: '♟'}
		b.addPiece(&pawnsW[i], 6, i)
		pawnsB[i] = Piece{White: false, Icon: '♙'}
		b.addPiece(&pawnsB[i], 1, i)
	}
}
