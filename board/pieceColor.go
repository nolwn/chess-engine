package board

import "errors"

type pieceColor uint8

const (
	PieceWhite pieceColor = iota
	PieceBlack
)

func (pc pieceColor) typeError() error {
	if pc > 1 {
		return errors.New("piece color must be either 0 (white) or 1 (black)")
	}

	return nil
}
