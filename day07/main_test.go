package main

import (
        "testing"

        "github.com/stretchr/testify/assert"
)

func TestNewChessBoard(t *testing.T) {
        board1 := NewChessBoard()
        board1.Move(1, 1, 0)
        board1.Move(1, 3, 2)

        assert.Equal(t, board1.chessPieces[1].Unit, board1.chessPieces[1].Unit)
}
