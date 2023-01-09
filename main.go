package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"time"
)

var solvedBoardForward [9][9]int
var solvedBoardBackward [9][9]int

func main() {

	var board = [9][9]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},

		{2, 3, 1, 5, 6, 4, 8, 9, 7},
		{5, 6, 4, 8, 9, 7, 2, 3, 1},
		{8, 9, 7, 2, 3, 1, 5, 6, 4},

		{3, 1, 2, 6, 4, 5, 9, 7, 8},
		{6, 4, 5, 9, 7, 8, 3, 1, 2},
		{9, 7, 8, 3, 1, 2, 6, 4, 5},
	}

	board = randomiseBoard(board)
	board = removeNumbers(board, 26)
	printBoard(board, 0)

}

func randomiseBoard(board [9][9]int) [9][9]int {

	board = ShuffleNumbers(board)
	board = ShuffleRows(board)
	board = ShuffleCols(board)
	board = Shuffle3X3Rows(board)
	board = Shuffle3X3Cols(board)

	return board

}

func ShuffleNumbers(board [9][9]int) [9][9]int {

	//seed rand
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 1; i < len(board[0])+1; i++ {

		ranNum := r.Intn(9) + 1
		board = swapNumbers(board, i, ranNum)

	}

	return board

}

func swapNumbers(board [9][9]int, n1 int, n2 int) [9][9]int {

	for y := 0; y < len(board[0]); y++ {
		for x := 0; x < len(board[0]); x++ {
			if board[x][y] == n1 {
				board[x][y] = n2
			} else if board[x][y] == n2 {
				board[x][y] = n1
			}
		}
	}

	return board

}

func ShuffleRows(board [9][9]int) [9][9]int {

	//seed rand
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var blockNumber int

	for i := 0; i < len(board[0])/3; i++ {
		ranNum := r.Intn(3)
		blockNumber = i / 3
		board = swapRows(board, i, blockNumber*3+ranNum)
	}

	return board

}

func swapRows(board [9][9]int, r1 int, r2 int) [9][9]int {

	row := board[r1]
	board[r1] = board[r2]
	board[r2] = row

	return board

}

func ShuffleCols(board [9][9]int) [9][9]int {

	//seed rand
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var blockNumber int

	for i := 0; i < len(board[0]); i++ {
		ranNum := r.Intn(3)
		blockNumber = i / 3
		board = swapCols(board, i, blockNumber*3+ranNum)
	}

	return board

}

func swapCols(board [9][9]int, c1 int, c2 int) [9][9]int {

	var colVal int
	for i := 0; i < len(board[0]); i++ {
		colVal = board[i][c1]
		board[i][c1] = board[i][c2]
		board[i][c2] = colVal
	}

	return board

}

func Shuffle3X3Rows(board [9][9]int) [9][9]int {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < len(board[0])/3; i++ {
		ranNum := r.Intn(3)
		board = swap3X3Rows(board, i, ranNum)
	}

	return board

}

func swap3X3Rows(board [9][9]int, r1 int, r2 int) [9][9]int {

	for i := 0; i < len(board[0])/3; i++ {
		board = swapCols(board, r1*3+i, r2*3+i)
	}

	return board

}

func Shuffle3X3Cols(board [9][9]int) [9][9]int {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < len(board[0])/3; i++ {
		ranNum := r.Intn(3)
		board = swap3X3Cols(board, i, ranNum)
	}

	return board

}

func swap3X3Cols(board [9][9]int, c1 int, c2 int) [9][9]int {

	for i := 0; i < len(board[0])/3; i++ {
		board = swapCols(board, c1*3+i, c2*3+i)
	}

	return board

}

func sliceCopy(in, out interface{}) {
	buf := new(bytes.Buffer)
	gob.NewEncoder(buf).Encode(in)
	gob.NewDecoder(buf).Decode(out)
}

func removeNumbers(board [9][9]int, n int) [9][9]int {

	var backupBoard [9][9]int
	sliceCopy(board, backupBoard)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	//n is the number of remaining numbers
	activeNumbers := len(board[0]) * len(board[0])

	for n < activeNumbers {

		// 0 if cell untested, 1 if tested
		var testedMap = [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},

			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},

			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}

		var completeMap = [9][9]int{
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},

			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},

			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
		}

		for n < activeNumbers {

			ranRow := r.Intn(9)
			ranCol := r.Intn(9)

			if testedMap == completeMap {
				return removeNumbers(backupBoard, n)
			}

			if board[ranRow][ranCol] != 0 {
				var temp = board[ranRow][ranCol]
				board[ranRow][ranCol] = 0
				if isSolvable(board, 0, 0) && isSolvableReverse(board, 0, 0) {
					if solvedBoardForward == solvedBoardBackward {
						activeNumbers--
					} else {
						board[ranRow][ranCol] = temp
					}
				} else {
					board[ranRow][ranCol] = temp
				}
				testedMap[ranRow][ranCol] = 1
			}

		}

	}
	return board
}

func isValid(board [9][9]int, r int, c int, k int) bool {

	notInRow := true
	for i := 0; i < len(board[0]); i++ {
		if k == board[r][i] {
			notInRow = false
		}
	}

	notInCol := true
	for i := 0; i < len(board[0]); i++ {
		if k == board[i][c] {
			notInCol = false
		}
	}

	notInBox := true
	for i := (r / 3) * 3; i < (r/3)*3+3; i++ {
		for j := (c / 3) * 3; j < (c/3)*3+3; j++ {
			if k == board[i][j] {
				notInBox = false
			}
		}
	}

	return (notInRow && notInCol && notInBox)

}

func isSolvable(board [9][9]int, r int, c int) bool {
	if r == 9 {
		solvedBoardForward = board
		return true
	} else if c == 9 {
		return isSolvable(board, r+1, 0)
	} else if board[r][c] != 0 {
		return isSolvable(board, r, c+1)
	} else {
		for k := 1; k < len(board[0])+1; k++ {
			if isValid(board, r, c, k) {
				board[r][c] = k
				if isSolvable(board, r, c+1) {
					return true
				}
				board[r][c] = 0
			}
		}
		return false
	}
}

func isSolvableReverse(board [9][9]int, r int, c int) bool {
	if r == 9 {
		solvedBoardBackward = board
		return true
	} else if c == 9 {
		return isSolvableReverse(board, r+1, 0)
	} else if board[r][c] != 0 {
		return isSolvableReverse(board, r, c+1)
	} else {
		for k := len(board[0]); k > 0; k-- {
			if isValid(board, r, c, k) {
				board[r][c] = k
				if isSolvableReverse(board, r, c+1) {
					return true
				}
				board[r][c] = 0
			}
		}
		return false
	}
}

func printBoard(board [9][9]int, flag int) {
	//takes board, flag 0 prints sudoku string, flag 1 prints a terminal sudoku

	if flag == 0 {
		fmt.Printf("\n")
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == 0 {
					fmt.Printf(".")
				} else {
					fmt.Printf("%d", board[i][j])
				}
			}
		}
		fmt.Printf("\n")
	}

	if flag == 1 {
		fmt.Printf("\n")
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == 0 {
					fmt.Printf("0")
				} else {
					fmt.Printf("%d", board[i][j])
				}
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}

}
