package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		{"_", "_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	fmt.Println("board", board)

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	fmt.Println("board", len(board[0][1]))
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}
