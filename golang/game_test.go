package main

import (
    "testing"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
    if a != b {
        t.Errorf("Got: %v, Wanted: %v.", a, b)
    }
}

func TestCellIndex(t *testing.T) {
    game := Game{4, 4, make([]bool, 9)}

    t.Run("CellIndex finds right cell", func(t *testing.T) {
        index := game.CellIndex(1, 1)

        AssertEqual(t, index, 5)
    })

    t.Run("CellIndex achieves wraparound", func(t *testing.T) {
        index := game.CellIndex(-1, 1)

        AssertEqual(t, index, 7)
    })
}

func TestNeighborIndices(t *testing.T) {
    game := Game{4, 4, make([]bool, 9)}

    t.Run("can find surrounding neighbors", func(t *testing.T) {
        neighbors := game.NeighborIndices(1, 1)

        AssertEqual(t, neighbors, [8]int {0, 1, 2, 4, 6, 8, 9, 10})
    })

    t.Run("can find surround neighbors with wraparound", func(t *testing.T) {
        neighbors := game.NeighborIndices(0, 0)

        AssertEqual(t, neighbors, [8]int {15, 12, 13, 3, 1, 7, 4, 5})
    })
}

func TestNeighborValues(t *testing.T) {
    game := Game{4, 4, []bool {
        true, true, false, false,
        false, true, true, true,
        true, true, false, false,
        true, false, true, false,
    }}

    t.Run("can find surround neighbor values", func(t *testing.T) {
        values := game.NeighborValues(1, 1)

        AssertEqual(t, values, [8]bool {true, true, false, false, true, true, true, false})
    })

    t.Run("can find surrounding neighbor values with wraparound", func(t *testing.T) {
        values := game.NeighborValues(0, 0)

        AssertEqual(t, values, [8]bool {false, true, false, false, true, true, false, true})
    })
}

func CreateGameWithNNeighbors(n int, self bool) Game {
    game := Game{3, 3, make([]bool, 9)}

    for i := 0; i < n; i++ {
        game.board[i] = true
    }

    copy(game.board[5:], game.board[4:])
    game.board[4] = self

    return game
}

func TestCreateGameWithNNeighbors(t *testing.T) {
    game := CreateGameWithNNeighbors(5, false)
    expected := []bool {true, true, true, true, false ,true, false, false, false}
    for index, value := range game.board {
        AssertEqual(t, value, expected[index])
    }
}

func TestTransition0(t *testing.T) {
    game := CreateGameWithNNeighbors(0, true)
    nextGame := game.Transition()

    AssertEqual(t, nextGame.board[4], false)
}

func TestTransition1(t *testing.T) {
    game := CreateGameWithNNeighbors(1, true)
    nextGame := game.Transition()

    AssertEqual(t, nextGame.board[4], false)
}

func TestTransition2(t *testing.T) {
    game := CreateGameWithNNeighbors(2, true)
    nextGame := game.Transition()

    AssertEqual(t, nextGame.board[4], true)
}

func TestTransition3(t *testing.T) {
    game := CreateGameWithNNeighbors(3, false)
    nextGame := game.Transition()

    AssertEqual(t, nextGame.board[4], true)
}

func TestTransition4(t *testing.T) {
    game := CreateGameWithNNeighbors(4, true)
    nextGame := game.Transition()

    AssertEqual(t, nextGame.board[4], false)
}

func TestTransition5(t *testing.T) {
    game := CreateGameWithNNeighbors(5, true)
    nextGame := game.Transition()

    AssertEqual(t, nextGame.board[4], false)
}

func TestTransition6(t *testing.T) {
    game := CreateGameWithNNeighbors(6, true)
    nextGame := game.Transition()

    AssertEqual(t, nextGame.board[4], false)
}

func TestTransition7(t *testing.T) {
    game := CreateGameWithNNeighbors(7, true)
    nextGame := game.Transition()

    AssertEqual(t, nextGame.board[4], false)
}

func TestTransition8(t *testing.T) {
    game := CreateGameWithNNeighbors(8, true)
    nextGame := game.Transition()

    AssertEqual(t, nextGame.board[4], false)
}
