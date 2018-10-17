package lib

import (
	"fmt"
)

type Board struct {
	cells         [10][10]Cell
	reveal        bool
	occupiedCount int
}

func NewBoard(reveal bool) Board {
	board := Board{}

	board.reveal = reveal

	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			board.cells[row][col] = Cell{Occupied: false, Content: EMPTY}
		}
	}

	return board
}

func (b *Board) GameOver() bool {
	return b.occupiedCount == 0
}

func (b *Board) Render() {
	fmt.Print(CLEAR)
	fmt.Println(HEADERS)

	for row := 0; row < 10; row++ {
		fmt.Printf("%02d", row+1)
		for col := 0; col < 10; col++ {
			cell := b.cells[col][row]
			switch {
			case cell.Occupied && b.reveal:
				fmt.Print(OCCUPIED)
			case cell.Occupied:
				fmt.Print(EMPTY)
			default:
				fmt.Print(cell.Content)
			}
		}
		fmt.Println()
	}
}

func (b *Board) Guess(col rune, row int) {

	cell := b.Get(col, row)

	if cell.Occupied {
		cell.Content = HIT
		cell.Occupied = false
		b.occupiedCount--
	} else {
		cell.Content = MISS
	}
}

func (b *Board) Get(col rune, row int) *Cell {
	return &b.cells[int(col)-65][row-1]
}

func (b *Board) Put(col rune, row int, cell Cell) {
	if cell.Occupied {
		b.occupiedCount++
	}

	b.cells[int(col)-65][row-1] = cell
}
