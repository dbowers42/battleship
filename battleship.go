package main

import (
	"fmt"
	"github.com/dbowers42/battleship/lib"
	"strconv"
	"strings"
)

func setup(board *lib.Board) {
	board.Put('B', 3, lib.Cell{Occupied: true, Content: lib.OCCUPIED})
	board.Put('B', 4, lib.Cell{Occupied: true, Content: lib.OCCUPIED})
	board.Put('B', 5, lib.Cell{Occupied: true, Content: lib.OCCUPIED})
	board.Put('B', 6, lib.Cell{Occupied: true, Content: lib.OCCUPIED})
	board.Put('B', 7, lib.Cell{Occupied: true, Content: lib.OCCUPIED})

	board.Put('I', 1, lib.Cell{Occupied: true, Content: lib.OCCUPIED})
	board.Put('J', 1, lib.Cell{Occupied: true, Content: lib.OCCUPIED})

	board.Put('D', 4, lib.Cell{Occupied: true, Content: lib.OCCUPIED})
	board.Put('E', 4, lib.Cell{Occupied: true, Content: lib.OCCUPIED})
	board.Put('F', 4, lib.Cell{Occupied: true, Content: lib.OCCUPIED})
	board.Put('G', 4, lib.Cell{Occupied: true, Content: lib.OCCUPIED})

	board.Put('B', 10,lib.Cell{Occupied: true, Content: lib.OCCUPIED})
	board.Put('C', 10,lib.Cell{Occupied: true, Content: lib.OCCUPIED})
	board.Put('D', 10,lib.Cell{Occupied: true, Content: lib.OCCUPIED})

	board.Put('H', 6, lib.Cell{Occupied: true, Content: lib.OCCUPIED})
	board.Put('H', 7, lib.Cell{Occupied: true, Content: lib.OCCUPIED})
	board.Put('H', 8, lib.Cell{Occupied: true, Content: lib.OCCUPIED})
}

func main(){
	board := lib.NewBoard(true)

	setup(&board)

	fmt.Println("BATTLESHIP")
	board.Render()

	for !board.GameOver() {
		var row, col string

		fmt.Print("Enter you guess: ")
		fmt.Scanf("%s %s", &col, &row)

		rowGuess, err := strconv.Atoi(row)
		colGuess := ([]rune(strings.ToUpper(col)))[0]

		if err != nil {
			fmt.Println(err)
		}

		board.Guess(colGuess, rowGuess)

		board.Render()
	}

	fmt.Println("Game Over")
}
