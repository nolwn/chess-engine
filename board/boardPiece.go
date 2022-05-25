package board

// boardPiece stores the pieces type, color and position.
type boardPiece struct {
	// piece is a number between 0 and 15 where 0-7 are the white pieces and 8-15 are the
	// black pieces.
	piece piece

	// position is a number between 0 and 63 where 0 is the lower left corner and 63 is
	// the upper right. The numbering proceeds such that a larger number is always either
	// above or directly to the right of a smaller number (like reading text from bottom
	// to top).
	position uint8
}

const boardSize = 8
const pieceCount = 6

// NewPiece returns a new piece.
func NewPiece() boardPiece {
	return boardPiece{}
}

// Color returns 0 if the piece is a white piece and 1 if it is a black piece.
func (bp *boardPiece) Color() pieceColor {
	if bp.piece < pieceCount {
		return PieceWhite
	} else {
		return PieceBlack
	}
}

// Color returns a uint8 which represents the kind of piece (King, Pawn, Rook etc.).
func (bp *boardPiece) Piece() pieceType {
	return bp.piece.piece()
}

// Position returns the position file.
func (bp *boardPiece) Position() (rank coord, file coord) {
	pos := bp.position // Position is stored as a single number.
	rank = coord(pos / boardSize)
	file = coord(pos % boardSize)

	return
}

// SetColor takes a piece color const value and sets it as the piece's color. If it
// receives and invalid value it will return an error.
func (bp *boardPiece) SetColor(c pieceColor) (err error) {
	if err = checkType(c); err != nil {
		return
	}

	if bp.piece.color() != c {
		if c == PieceWhite {
			// The first six pieces are white...
			bp.piece -= pieceCount
		} else {
			// ...the second six are black.
			bp.piece += pieceCount
		}
	}

	return
}

// SetPiece takes a piece type between 0 (Pawn) and 5 (King). If it receives and invalid
// value, it will return an error.
func (bp *boardPiece) SetPiece(p pieceType) (err error) {
	if err = checkType(p); err != nil {
		return
	}

	// Reset the piece to the first piece of each color...
	if bp.piece.color() == PieceWhite {
		bp.piece = 0
	} else {
		bp.piece = pieceCount
	}

	// ...then add the piece number.
	bp.piece += piece(p)

	return
}

// SetPosition takes a rank and a file as uint8s and sets the piece to the corresponding
// position on the board. The rank and file must be between 0 and 7 or an error will be
// returned.
func (bp *boardPiece) SetPosition(rank coord, file coord) (err error) {
	if err = checkType(rank, file); err != nil {
		return err
	}

	bp.position = uint8(boardSize*rank + file)

	return nil
}

func coordsToPos(rank coord, file coord) uint8 {
	return uint8(boardSize*rank + file)
}
