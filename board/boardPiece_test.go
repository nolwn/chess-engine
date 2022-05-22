package board_test

import (
	"testing"

	"github.com/nolwn/chess-engine/board"
)

func TestSetPieceMethodAcceptsValidInput(t *testing.T) {
	p := board.NewPiece()

	if err := p.SetPiece(board.PieceKnight); err != nil {
		t.Fatalf("boardPiece.SetPiece failed to set a valid piece: %s", err)
	}

	if err := p.SetPiece(board.PieceKnight); err != nil {
		t.Fatalf("boardPiece.SetPiece failed to set a valid piece: %s", err)
	}

	if err := p.SetPiece(board.PiecePawn); err != nil {
		t.Fatalf("boardPiece.SetPiece failed to set a valid piece: %s", err)
	}
}

func TestSetPieceMethodRejectsInvalidInput(t *testing.T) {
	p := board.NewPiece()

	if err := p.SetPiece(12); err == nil {
		t.Fatal("boardPiece.SetPiece didn't return an error when it received and " +
			"invalid input")
	}
}

func TestPieceMethodShouldReturnASetPiece(t *testing.T) {
	p := board.NewPiece()

	p.SetPiece(board.PieceBishop)
	if pt := p.Piece(); pt != board.PieceBishop {
		t.Fatalf("expected boardPiece.Piece to return 2 (bishop), but received: %d", pt)
	}

	p.SetPiece(board.PieceBishop)
	if pt := p.Piece(); pt != board.PieceBishop {
		t.Fatalf("expected board.Piece to return 2 (bishop) after 2 was set a second "+
			"time, but received: %d", pt)
	}

	p.SetPiece(board.PiecePawn)
	if pt := p.Piece(); pt != board.PiecePawn {
		t.Fatalf("expected boardPiece.Piece to return 0 (pawn), but received: %d", pt)
	}

	p.SetPiece(board.PieceKing)
	if pt := p.Piece(); pt != board.PieceKing {
		t.Fatalf("expected boardPiece.Piece to return 5 (king), but received: %d", pt)
	}
}

func TestSetColorMethodAcceptsValidInput(t *testing.T) {
	p := board.NewPiece()

	if err := p.SetColor(board.PieceBlack); err != nil {
		t.Fatal("boardPiece.SetColor should accept a valid color")
		return
	}
}

func TestSetColorMethodShouldRejectInvalidInput(t *testing.T) {
	p := board.NewPiece()

	if err := p.SetColor(7); err == nil {
		t.Fatal("boardPiece.SetColor should return an error when it recieves a number " +
			"larger than 1")
		return
	}

	if p.Color() == 7 {
		t.Fatal("boardPiece.SetColor should not set an invalid color")
		return
	}
}

func TestColorMethodShouldReturnASetColor(t *testing.T) {
	p := board.NewPiece()

	p.SetColor(board.PieceWhite)
	if p.Color() != board.PieceWhite {
		t.Fatal("boardPiece.Color did not return 0 (white) after that color was set")
	}

	p.SetColor(board.PieceBlack)
	if p.Color() != board.PieceBlack {
		t.Fatal("boardPiece.Color did not return 1 (black) after that color was set")
	}

	p.SetColor(board.PieceBlack) // to rule out a toggling behavior
	if p.Color() != board.PieceBlack {
		t.Fatal("boardPiece.Color did not return 1 (black) after that color was reset")
	}

	p.SetColor(board.PieceWhite)
	if c := p.Color(); c != board.PieceWhite {
		t.Fatalf("boardPiece.Color did not return 0 (white) after that color was set " +
			"again.")
	}
}

// Because color and type are combined internally, it's important to make sure you can
// set one without changing the other.
func TestSetColorMethodShouldNotChangePiece(t *testing.T) {
	p := board.NewPiece()

	p.SetPiece(board.PieceQueen)
	p.SetColor(board.PieceBlack)

	if p.Color() != board.PieceBlack {
		t.Fatalf("expected boardPiece.SetColor to be able to set the color of a queen " +
			"to black")
	}

	if pt := p.Piece(); pt != board.PieceQueen {
		t.Fatalf("expected boardPiece.SetColor to only affect the color, but the piece "+
			"became: %d", pt)
	}
}

// Because color and type are combined internally, it's important to make sure you can
// set one without changing the other.
func TestSetPieceMethodShouldNotChangeColor(t *testing.T) {
	p := board.NewPiece()

	p.SetColor(board.PieceBlack)
	p.SetPiece(board.PieceKnight)

	if p.Color() != board.PieceBlack {
		t.Fatalf("expected p.SetPiece to only affect to piece type, but the color " +
			"changed")
	}

	if pt := p.Piece(); pt != board.PieceKnight {
		t.Fatalf("expected boardPiece.SetPiece to change a black piece to a knight, "+
			"but received: %d", pt)
	}
}

func TestSetPositionMethodAcceptsValidInput(t *testing.T) {
	p := board.NewPiece()

	if err := p.SetPosition(4, 7); err != nil {
		t.Fatalf("expected board.SetPosition to accept valid inputs, but received: %s",
			err)
	}
}

func TestSetPositionMethodRejectsInvalidRank(t *testing.T) {
	p := board.NewPiece()

	if err := p.SetPosition(8, 1); err == nil {
		t.Fatal("expected board.SetPosition to reject an invalid rank")
	}

	rank, file := p.Position()

	if rank == 8 {
		t.Fatal("expected a rejected rank not to be set")
	}

	if file == 1 {
		t.Fatal("expected file not to be set if the rank was rejected")
	}
}

func TestSetPositionMethodRejectsInvalidFile(t *testing.T) {
	p := board.NewPiece()

	if err := p.SetPosition(4, 12); err == nil {
		t.Fatal("expected board.SetPosition to reject an invalid file")
	}

	rank, file := p.Position()

	if file == 12 {
		t.Fatal("expected a rejected file not to be set")
	}

	if rank == 4 {
		t.Fatal("expected rank not to be set if the file was rejected")
	}
}

func TestPositionMethodShouldReturnASetPosition(t *testing.T) {
	p := board.NewPiece()

	p.SetPosition(3, 5)

	rank, file := p.Position()

	if rank != 3 {
		t.Fatalf("expected boardPiece.Position to return (3, 5), but received: (%d, %d)",
			rank, file)
	}

	if file != 5 {
		t.Fatalf("expected boardPiece.Position to return (3, 5), but received: (%d, %d)",
			rank, file)
	}
}
