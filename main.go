package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	var board = [][]int{
		{1,2,3, 4,5,6, 7,8,9},
		{4,5,6, 7,8,9, 1,2,3},
		{7,8,9, 1,2,3, 4,5,6},

		{2,3,1, 5,6,4, 8,9,7},
		{5,6,4, 7,8,9, 2,3,1},
		{8,9,7, 2,3,1, 5,6,4},

		{3,1,2, 6,4,5, 9,7,8},
		{6,4,5, 9,7,8, 3,1,2},
		{9,7,8, 3,1,2, 6,4,5},
	}

	board = randomiseBoard(board)

	board = removeNumbers(board, 50)

	for i := 0; i < len(board[0]); i++ {
		fmt.Println(board[i])
	}

	fmt.Println(solve(board, 0, 0))


}

func randomiseBoard(board [][]int) [][]int {

	board = ShuffleNumbers(board)
	board = ShuffleRows(board)
	board = ShuffleCols(board)
	board = Shuffle3X3Rows(board)
	board = Shuffle3X3Cols(board)

	return board

}

func ShuffleNumbers(board [][]int) [][]int {

	//seed rand
	s := rand.NewSource(time.Now().UnixNano())
    r := rand.New(s)

	for i := 1; i < len(board[0])+1; i++ {

		ranNum := r.Intn(9)+1
		board = swapNumbers(board, i, ranNum)
		
	}

	return board

}

func swapNumbers(board [][]int, n1 int, n2 int) [][]int {

	for y := 0; y < len(board[0]); y++ {
		for x := 0; x < len(board[0]); x++ {
			if (board[x][y] == n1) {
                board[x][y] = n2;
            } else if (board[x][y] == n2) {
                board[x][y] = n1;
            }
		}	
	}

	return board

}

func ShuffleRows(board [][]int) [][]int {

	//seed rand
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var blockNumber int;

	for i := 0; i < len(board[0])/3; i++ {
		ranNum := r.Intn(3);
		blockNumber = i / 3;
		board = swapRows(board, i, blockNumber * 3 + ranNum);
	}

	return board

}

func swapRows(board [][]int, r1 int, r2 int) [][]int {

	row := board[r1]
	board[r1] = board[r2]
	board[r2] = row

	return board

}


func ShuffleCols(board [][]int) [][]int {

	//seed rand
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var blockNumber int;

	for i := 0; i < len(board[0]); i++ {
		ranNum := r.Intn(3);
		blockNumber = i / 3;
		board = swapCols(board, i, blockNumber * 3 + ranNum);
	}

	return board

}

func swapCols(board [][]int, c1 int, c2 int) [][]int {

	var colVal int;
	for i := 0; i < len(board[0]); i++ {
		colVal = board[i][c1];
		board[i][c1] = board[i][c2];
        board[i][c2] = colVal;
	}

	return board

}

func Shuffle3X3Rows(board [][]int) [][]int {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < len(board[0])/3; i++ {
		ranNum := r.Intn(3);
		board = swap3X3Rows(board, i, ranNum) 
	}

	return board

}

func swap3X3Rows(board [][]int, r1 int, r2 int) [][]int {

	for i := 0; i < len(board[0])/3; i++ {
		board = swapCols(board, r1 * 3 + i, r2 * 3 + i);
	}

	return board

}

func Shuffle3X3Cols(board [][]int) [][]int {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < len(board[0])/3; i++ {
		ranNum := r.Intn(3);
		board = swap3X3Cols(board, i, ranNum) 
	}

	return board

}

func swap3X3Cols(board [][]int, c1 int, c2 int) [][]int {

	for i := 0; i < len(board[0])/3; i++ {
		board = swapCols(board, c1 * 3 + i, c2 * 3 + i);
	}

	return board

}

func removeNumbers(board [][]int, n int) [][]int {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	//n is the number of remaining numbers
	activeNumbers := len(board[0]) * len(board[0])

	for n < activeNumbers {

		ranRow := r.Intn(9)
		ranCol := r.Intn(9)

		if (board[ranRow][ranCol] != 0) {
			board[ranRow][ranCol] = 0
			activeNumbers--
		}
		
	}

	return board

}

func isValid(board [][]int, r int, c int, k int) bool {

	notInRow := true
	for i := 0; i < len(board[0]); i++ {
		if (k == board[r][i]) {
			notInRow = false
		}
	}

	notInCol := true
	for i := 0; i < len(board[0]); i++ {
		if (k == board[i][c]) {
			notInCol = false
		}
	}

	notInBox := true
	for i := (r/3)*3; i < (r/3)*3+3; i++ {
		for j := (c/3)*3; j < (c/3)*3+3; j++ {
			if (k == board[i][j]) {
				notInCol = false
			}
		}
	}

	return (notInRow && notInCol && notInBox)

}

func solve(board [][]int, r int, c int) bool {
	if (r == 9) {
		for i := 0; i < len(board[0]); i++ {
			fmt.Println(board[i])
		}
		return true
	} else if (c == 9) {
		return solve(board, r+1, 0)
	} else if (board[r][c] != 0) {
		return solve(board, r, c+1)
	} else {
		for k := 1; k < len(board[0])+1; k++ {
			if isValid(board, r, c, k) {
				board[r][c] = k
				if solve(board, r, c+1) {
					return true
				}
				board[r][c] = 0
			}
		}
		return false
	}
}