package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	puzzle := [3][3]int{}
	puzzle = initBoard(puzzle)
	printBoard(puzzle)

}

func isAlreadyPresent(x int, a [3][3]int) bool {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if a[i][j] == x {
				return true
			}
		}
	}
	return false
}

func getUniqueAndRandomNum(a [3][3]int) int {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(9) + 1

	// fmt.Println(x)

	if !isAlreadyPresent(x, a) {
		return x
	}
	return getUniqueAndRandomNum(a)

}

func printBoard(board [3][3]int) {

	fmt.Println("| ", getNum(board, 0, 0), " | ", getNum(board, 0, 1), " | ", getNum(board, 0, 2), " |")
	fmt.Println("| ", getNum(board, 1, 0), " | ", getNum(board, 1, 1), " | ", getNum(board, 1, 2), " |")
	fmt.Println("| ", getNum(board, 2, 0), " | ", getNum(board, 2, 1), " | ", getNum(board, 2, 2), " |")
	fmt.Println(board)

}

func initBoard(board [3][3]int) [3][3]int {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = 0
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = getUniqueAndRandomNum(board)
		}
	}
	return board

}

func getNum(board [3][3]int, r int, c int) string {

	if board[r][c] == 9 {
		return " "
	}
	return strconv.Itoa(board[r][c])

}

func isBoardFinished(board [3][3]int) bool {
	if board[0][0] == 1 && board[0][1] == 2 && board[0][2] == 3 &&
		board[1][0] == 4 && board[1][1] == 5 && board[1][2] == 6 &&
		board[2][0] == 7 && board[2][1] == 8 && board[2][2] == 9 {
		return true
	}
	return false

}
