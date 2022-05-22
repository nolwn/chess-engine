package board

type piece uint8

func (p piece) color() pieceColor {
	if p < 6 {
		return PieceWhite
	} else {
		return PieceBlack
	}
}

func (p piece) piece() pieceType {
	if p.color() == PieceBlack {
		return pieceType(p - 6)
	} else {
		return pieceType(p)
	}
}
