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

	if p := b.Piece(5, 2); p != nil {
		t.Fatal("expected piece not to be set it was rejected")
	}
}

func TestBoardPieceMethodReturnsNilWhenEmpty(t *testing.T) {
	b := board.New()

	if p := b.Piece(5, 3); p != nil {
		t.Fatalf("expected to receive nil when a piece hasn't been placed at the given"+
			"spot, but received: %v", p)
	}
}

func TestBoardPieceMethodReturnsPieceWhenSet(t *testing.T) {
	b := board.New()

	b.SetPiece(board.PieceKing, board.PieceBlack, 4, 6)

	p1 := b.Piece(4, 6)
	if p1 == nil {
		t.Fatal("expected to receive a piece after it was set")
	}

	if p := b.Piece(4, 4); p != nil {
		t.Fatal("expected not to receive a piece on a square where a piece wasn't set")
	}

	b.SetPiece(board.PieceKnight, board.PieceWhite, 4, 4)

	p2 := b.Piece(4, 4)
	if p2 == nil {
		t.Fatal("expected to receive a piece after it was set")
	}

	if p1.Piece() == p2.Piece() {
		t.Fatal("expected different pieces on different squares to be different")
	}
}
