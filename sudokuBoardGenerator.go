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

	for i := 0; i < len(board[0]); i++ {
		fmt.Println(board[i])
	}

	fmt.Println("\n\n")
	randomiseBoard(board)

	for i := 0; i < len(board[0]); i++ {
		fmt.Println(board[i])
	}

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

	for i := 0; i < len(board[0]); i++ {

		ranNum := r.Intn(9)
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
