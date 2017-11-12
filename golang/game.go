package main

import (
    "fmt"
    "bytes"
    "math/rand"
    "time"
)


type Game struct {
    width int
    height int
    board []bool
}

type Pair struct {
    x int
    y int
}


func (game Game) CellIndex(x int, y int) int {
    modx := (x % game.width + game.width) % game.width
    mody := (y % game.height + game.height) % game.height

    return modx + mody * game.height
}

func (game Game) NeighborIndices(x int, y int) [8]int {
    neighbors := []Pair {
        Pair{x - 1, y - 1}, Pair{x, y - 1}, Pair{x + 1, y - 1},
        Pair{x - 1, y}, Pair{x + 1, y},
        Pair{x - 1, y + 1}, Pair{x, y + 1}, Pair{x + 1, y + 1},
    }

    result := [8]int{}

    for index, pair := range neighbors {
        result[index] = game.CellIndex(pair.x, pair.y)
    }

    return result
}

func (game Game) NeighborValues(x int, y int) [8]bool {
    neighbor_indices := game.NeighborIndices(x, y)
    result := [8]bool {}

    for index, neighbor_index := range neighbor_indices {
        result[index] = game.board[neighbor_index]
    }

    return result
}

func (game Game) Transition() Game {
    nextGame := Game{game.width, game.height, make([]bool, game.width * game.height)}

    for y := 0; y < game.height; y++ {
        for x:= 0; x < game.width; x++ {
            index := game.CellIndex(x, y)
            neighbor_values := game.NeighborValues(x, y)
            number_alive := 0

            for _, is_neighbor_alive := range neighbor_values {
                if is_neighbor_alive {
                    number_alive++
                }
            }

            if number_alive == 3 {
                nextGame.board[index] = true
            } else if number_alive == 2 {
                nextGame.board[index] = game.board[index]
            }
        }
    }

    return nextGame
}

func (game Game) String() string {
    var buffer bytes.Buffer

    for y := 0; y < game.height; y++ {
        for x:= 0; x < game.height; x++ {
            index := game.CellIndex(x, y)
            is_alive := game.board[index]

            if is_alive {
                buffer.WriteString("o ")
            } else {
                buffer.WriteString(". ")
            }
        }

        buffer.WriteString("\n")
    }

    return buffer.String()
}

func RandomGame(width int, height int) Game {
    rand.Seed(time.Now().UnixNano())
    game := Game {width, height, make([]bool, width * height)}

    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            index := game.CellIndex(x, y)

            game.board[index] = rand.Float32() < 0.5
        }
    }

    return game
}

func main() {
    game := RandomGame(50, 50)

    for {
        game = game.Transition()
        fmt.Println("\033[H")
        fmt.Println("\033[2J")
        fmt.Println(game.String())
        time.Sleep(time.Nanosecond * 5e7)
    }
}
