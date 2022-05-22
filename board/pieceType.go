package board

import "errors"

type pieceType uint8

const (
	PiecePawn pieceType = iota
	PieceKnight
	PieceBishop
	PieceRook
	PieceQueen
	PieceKing
)

func (pt pieceType) typeError() (err error) {
	if pt >= 6 {
		err = errors.New("piece type must be between 0 (pawn) and 5 (king)")
	}

	return
}
