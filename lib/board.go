package lib

import "fmt"

type Board struct {
	cells [][]string
}

func NewBoard() Board {
	return Board{}
}

func (b *Board) Render() {
	fmt.Println("   A  B  C  D  E  F  G  H  I  J")
	for row := 0; row < 10; row++ {
		fmt.Printf("%02d", row + 1)
		for col := 0; col <10; col++ {
			fmt.Print(" . ")
		}
		fmt.Println()
	}
}

func (b *Board) Miss(row, col int) {

}

func (b *Board) Hit(row, col int) {

}