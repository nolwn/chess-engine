package board

import (
	"fmt"
)

type board struct {
	pieces   []*boardPiece
	pieceMap map[uint8]*boardPiece
}

// New returns a new board
func New() board {
	return board{pieces: make([]*boardPiece, 16), pieceMap: make(map[uint8]*boardPiece)}
}

// SetPiece takes a piece value, a color, and a position and puts a piece into that
// position on the board. If there is already a piece there, or if any of the
// arguments provided aren't valid, an error is returned.
//
// pieceType is a uint8 representing what kind of piece it is. Use const PiecePawn,
// pieceKnight
func (b *board) SetPiece(p pieceType, pc pieceColor, rank coord, file coord) error {
	n := NewPiece()

	if err := n.SetPiece(p); err != nil {
		return err
	}

	if err := n.SetColor(pc); err != nil {
		return err
	}

	if err := n.SetPosition(rank, file); err != nil {
		return err
	}

	if _, ok := b.pieceMap[n.position]; ok {
		return fmt.Errorf("there is already a piece at rank %d, file %d", rank, file)
	}

	b.pieces = append(b.pieces, &n)
	b.pieceMap[n.position] = b.pieces[len(b.pieces)-1]

	return nil
}

// Piece takes a rank and a file and returns a pointer to whatever piece is at that
// location. If there is no piece, it returns nil.
func (b *board) Piece(rank coord, file coord) *boardPiece {
	pos := coordsToPos(rank, file)
	p, ok := b.pieceMap[pos]

	if !ok {
		return nil
	}

	return p
}
