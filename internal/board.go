package internal

/*
	WHITE_PAWN : '♙',
    WHITE_ROOK : '♖',
    WHITE_KNIGHT : '♘',
    WHITE_BISHOP : '♗',
    WHITE_QUEEN : '♕',
    WHITE_KING: '♔',
	WHITE_SPOT: '□',

    BLACK_PAWN : '♟',
    BLACK_ROOK : '♜',
    BLACK_KNIGHT : '♞',
    BLACK_BISHOP : '♝',
    BLACK_QUEEN : '♛',
    BLACK_KING: '♚',
	BLACK_SPOT: '■'
*/

type Piece struct {
	Killed bool
	White  bool
	Icon   rune
}

type King struct {
	Piece
	Castled bool
}

type Pawn struct {
	Piece
	Moved bool
}

type Rook struct {
	Piece
	Moved bool
}

type Knight struct {
	Piece
}

type Bishop struct {
	Piece
}

type Queen struct {
	Piece
}

type Spot struct {
	Piece      *Piece
	file, rank int
	Icon       rune
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

func (b *Board) Print() {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if b.Spots[i][j].Piece != nil {
				if b.Spots[i][j].Piece.White {
					b.Spots[i][j].Icon = b.Spots[i][j].Piece.Icon
				} else {
					b.Spots[i][j].Icon = b.Spots[i][j].Piece.Icon
				}
			}
			print(string(b.Spots[i][j].Icon) + " ")
		}
		println()
	}
}
