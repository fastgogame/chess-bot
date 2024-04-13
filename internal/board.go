package internal

const (
	whitePawn   = '♟'
	blackPawn   = '♙'
	whiteKnight = '♞'
	blackKnight = '♘'
	whiteBishop = '♝'
	blackBishop = '♗'
	whiteRook   = '♜'
	blackRook   = '♖'
	whiteQueen  = '♛'
	blackQueen  = '♕'
	whiteKing   = '♚'
	blackKing   = '♔'
	whiteSquare = '■'
	blackSquare = '□'
)

type Piece struct {
	White bool
	Icon  rune
	Moved bool
}

type Spot struct {
	Piece      *Piece
	file, rank int
	Square     rune
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
				b.Spots[i][j].Square = whiteSquare
			} else {
				b.Spots[i][j].Square = blackSquare
			}
		}
	}
}

func NewBoard() *Board {
	board := &Board{}
	board.Init()
	return board
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

func (b *Board) placePieceOnBoard(p *Piece, file, rank int) {
	b.Spots[file][rank].Piece = p
}

func newPiece(white bool, icon rune) *Piece {
	return &Piece{White: white, Icon: icon}
}

func (b *Board) DefaultSetup() {
	b.WhiteTurn = true

	kingW := newPiece(true, whiteKing)
	b.placePieceOnBoard(kingW, 7, 4)
	kingB := newPiece(false, blackKing)
	b.placePieceOnBoard(kingB, 0, 4)

	queenW := newPiece(true, whiteQueen)
	b.placePieceOnBoard(queenW, 7, 3)
	queenB := newPiece(false, blackQueen)
	b.placePieceOnBoard(queenB, 0, 3)

	rookWL := newPiece(true, whiteRook)
	b.placePieceOnBoard(rookWL, 7, 0)
	rookWR := newPiece(true, whiteRook)
	b.placePieceOnBoard(rookWR, 7, 7)
	rookBL := newPiece(false, blackRook)
	b.placePieceOnBoard(rookBL, 0, 0)
	rookBR := newPiece(false, blackRook)
	b.placePieceOnBoard(rookBR, 0, 7)

	bishopWL := newPiece(true, whiteBishop)
	b.placePieceOnBoard(bishopWL, 7, 2)
	bishopWR := newPiece(true, whiteBishop)
	b.placePieceOnBoard(bishopWR, 7, 5)
	bishopBL := newPiece(false, blackBishop)
	b.placePieceOnBoard(bishopBL, 0, 2)
	bishopBR := newPiece(false, blackBishop)
	b.placePieceOnBoard(bishopBR, 0, 5)

	knightWL := newPiece(true, whiteKnight)
	b.placePieceOnBoard(knightWL, 7, 1)
	knightWR := newPiece(true, whiteKnight)
	b.placePieceOnBoard(knightWR, 7, 6)
	knightBL := newPiece(false, blackKnight)
	b.placePieceOnBoard(knightBL, 0, 1)
	knightBR := newPiece(false, blackKnight)
	b.placePieceOnBoard(knightBR, 0, 6)

	var pawnsW, pawnsB [8]Piece
	for i := 0; i < 8; i++ {
		pawnsW[i] = *newPiece(true, whitePawn)
		b.placePieceOnBoard(&pawnsW[i], 6, i)
		pawnsB[i] = *newPiece(false, blackPawn)
		b.placePieceOnBoard(&pawnsB[i], 1, i)
	}
}
