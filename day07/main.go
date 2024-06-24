package main

import (
    "fmt"
)

var units = map[int]*ChessPieceUnit{
    1: {
        ID:    1,
        Name:  "車",
        Color: "red",
    },
    2: {
        ID:    2,
        Name:  "炮",
        Color: "red",
    },
    // 添加其他棋子
    // 3: {ID: 3, Name: "馬", Color: "red"},
    // ...
}


// MoveRule 移动规则接口
type MoveRule interface {
    CanMove(currentX, currentY, targetX, targetY int) bool
}

// RookMoveRule "車" 的移动规则
type RookMoveRule struct{}

func (r RookMoveRule) CanMove(currentX, currentY, targetX, targetY int) bool {
    // "車" 可以直线移动
    return currentX == targetX || currentY == targetY
}

// CannonMoveRule "炮" 的移动规则
type CannonMoveRule struct{}

func (c CannonMoveRule) CanMove(currentX, currentY, targetX, targetY int) bool {
    // "炮" 可以直线移动，但实际规则更复杂
    // 这里只实现简单的直线移动规则
    return currentX == targetX || currentY == targetY
}

type ChessPieceUnit struct {
    ID       uint
    Name     string
    Color    string
    MoveRule MoveRule
}

// NewChessPieceUnit 工厂
func NewChessPieceUnit(id int) *ChessPieceUnit {
    unit := units[id]
    switch unit.Name {
    case "車":
        unit.MoveRule = RookMoveRule{}
    case "炮":
        unit.MoveRule = CannonMoveRule{}
    // 其他棋子的规则可以在这里初始化
    }
    return unit
}

type ChessPiece struct {
    Unit *ChessPieceUnit
    X    int
    Y    int
}

// ChessBoard 棋局
type ChessBoard struct {
    chessPieces map[int]*ChessPiece
}

// NewChessBoard 初始化棋盘
func NewChessBoard() *ChessBoard {
    board := &ChessBoard{chessPieces: map[int]*ChessPiece{}}
    initialPositions := map[int][2]int{
        1: {0, 0}, // 红车1在(0, 0)
        2: {2, 3}, // 红炮2在(2, 3)
    }

    for id, pos := range initialPositions {
        board.chessPieces[id] = &ChessPiece{
            Unit: NewChessPieceUnit(id),
            X:    pos[0],
            Y:    pos[1],
        }
    }
    return board
}

// Move 移动棋子
func (c *ChessBoard) Move(id, x, y int) {
    if piece, exists := c.chessPieces[id]; exists {
        if piece.Unit.MoveRule.CanMove(piece.X, piece.Y, x, y) {
            piece.X = x
            piece.Y = y
            fmt.Printf("棋子 %s 移动到 (%d, %d)\n", piece.Unit.Name, x, y)
        } else {
            fmt.Printf("棋子 %s 无法移动到 (%d, %d)\n", piece.Unit.Name, x, y)
        }
    } else {
        fmt.Printf("没有找到ID为 %d 的棋子\n", id)
    }
}
