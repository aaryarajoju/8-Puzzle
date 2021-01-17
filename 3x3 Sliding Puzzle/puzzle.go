package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var stepCounter int
var totalTimeTaken time.Duration

func main() {

	startTime := time.Now()

	puzzle := [3][3]int{}
	puzzle = initBoard(puzzle)

	//boardTimeTaken := time.Since(startTime)
	//fmt.Println("Time taken to generate the board:", boardTimeTaken)

	fmt.Print("\nRules: \n",
		"\t`U` or `u`  for UP\n",
		"\t`D` or `d`  for DOWN\n",
		"\t`R` or `r`  for right\n",
		"\t`L` or `l`  for left\n")

	originalBoard := puzzle

start:
	fmt.Print("\n\n")
	printBoard(puzzle)

	fmt.Print("\n> ")

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	if char != 76 && char != 85 && char != 68 && char != 82 &&
		char != 108 && char != 117 && char != 100 && char != 114 {
		fmt.Println("INVALID MOVE")
		goto start
	}

	var dir int

	switch char {
	case 76, 108:
		dir = 1 //left
	case 85, 117:
		dir = 2 //up
	case 68, 100:
		dir = 3 //down
	case 82, 114:
		dir = 4 //right
	}

	positionGapI, positionGapJ := findGap(puzzle)

	if positionGapI == 0 && dir == 3 {
		fmt.Println("INVALID MOVE")
		goto start
	}
	if positionGapI == 2 && dir == 2 {
		fmt.Println("INVALID MOVE")
		goto start
	}
	if positionGapJ == 0 && dir == 4 {
		fmt.Println("INVALID MOVE")
		goto start
	}
	if positionGapJ == 2 && dir == 1 {
		fmt.Println("INVALID MOVE")
		goto start
	}

	if dir == 1 {
		i := positionGapI
		j := positionGapJ + 1
		num := getNum(puzzle, i, j)
		puzzle[positionGapI][positionGapJ], err = strconv.Atoi(num)
		puzzle[i][j] = 9
	} else if dir == 2 {
		i := positionGapI + 1
		j := positionGapJ
		num := getNum(puzzle, i, j)
		puzzle[positionGapI][positionGapJ], err = strconv.Atoi(num)
		puzzle[i][j] = 9
	} else if dir == 3 {
		i := positionGapI - 1
		j := positionGapJ
		num := getNum(puzzle, i, j)
		puzzle[positionGapI][positionGapJ], err = strconv.Atoi(num)
		puzzle[i][j] = 9
	} else if dir == 4 {
		i := positionGapI
		j := positionGapJ - 1
		num := getNum(puzzle, i, j)
		puzzle[positionGapI][positionGapJ], err = strconv.Atoi(num)
		puzzle[i][j] = 9
	}

	stepCounter++

	if !isBoardFinished(puzzle) {
		goto start
	}

	totalTimeTaken = time.Since(startTime)
	boardFinished(puzzle, originalBoard)
}

func isAlreadyPresent(board [3][3]int, x int) bool {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == x {
				return true
			}
		}
	}
	return false
}

func getUniqueAndRandomNum(board [3][3]int) int {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(9) + 1

	if !isAlreadyPresent(board, x) {
		return x
	}
	return getUniqueAndRandomNum(board)
}

func printBoard(board [3][3]int) {

	fmt.Println("| ", getNum(board, 0, 0), " | ", getNum(board, 0, 1), " | ", getNum(board, 0, 2), " |")
	fmt.Println("| ", getNum(board, 1, 0), " | ", getNum(board, 1, 1), " | ", getNum(board, 1, 2), " |")
	fmt.Println("| ", getNum(board, 2, 0), " | ", getNum(board, 2, 1), " | ", getNum(board, 2, 2), " |")
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

func findGap(board [3][3]int) (int, int) { return findPositionOfNum(board, 9) }

func findPositionOfNum(board [3][3]int, num int) (int, int) {

	var positionI, positionJ int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == num {
				positionI = i
				positionJ = j
				goto returning
			}
		}
	}

returning:
	return positionI, positionJ
}

func boardFinished(board [3][3]int, originalBoard [3][3]int) {
	fmt.Println("\n\nThe board is solved.")
	fmt.Println("It took and", totalTimeTaken, "and", stepCounter, "steps to solve")
	fmt.Println("\nThe original board was: ")
	printBoard(originalBoard)
	fmt.Println("\nThe final board is: ")
	printBoard(board)
}
