package board_test

import (
	"testing"

	"github.com/nolwn/chess-engine/board"
)

func TestBoardSetPieceMethodAcceptsValidInput(t *testing.T) {
	b := board.New()

	if err := b.SetPiece(board.PieceBishop, board.PieceBlack, 4, 6); err != nil {
		t.Fatalf("expected board.SetPiece to accept input, but recieved: %s", err)
	}
}

func TestBoardSetPieceMethodRejectsBadInput(t *testing.T) {
	b := board.New()

	if err := b.SetPiece(200, board.PieceWhite, 5, 2); err == nil {
		t.Fatal("expected an error when piece was invalid")
	}

	if err := b.SetPiece(board.PieceBishop, 200, 6, 3); err == nil {
		t.Fatal("expected an error when color was invalid")
	}

	if err := b.SetPiece(board.PiecePawn, board.PieceWhite, 200, 4); err == nil {
		t.Fatal("expected an error when rank was invalid")
	}

	if err := b.SetPiece(board.PiecePawn, board.PieceWhite, 4, 200); err == nil {
		t.Fatal("expected an error when file was invalid")
	}
}
