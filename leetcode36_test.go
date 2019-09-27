package main

import (
	"fmt"
	"testing"
)

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		hash := make(map[byte]byte, 0)
		for j := 0; j < 9; j++ {
			if board[i][j] == byte('.') {
				continue
			}
			if _, ok := hash[board[i][j]]; ok {
				return false
			}
			hash[board[i][j]] = board[i][j]
		}
	}

	for j := 0; j < 9; j++ {
		hash := make(map[byte]byte, 0)
		for i := 0; i < 9; i++ {
			if board[i][j] == byte('.') {
				continue
			}
			if _, ok := hash[board[i][j]]; ok {
				return false
			}
			hash[board[i][j]] = board[i][j]
		}
	}

	for m := 0; m < 3; m++ {
		for n := 0; n < 3; n++ {
			hash := make(map[byte]byte, 0)
			for i := m * 3; i < m*3+3; i++ {
				for j := n * 3; j < n*3+3; j++ {
					if board[i][j] == byte('.') {
						continue
					}
					if _, ok := hash[board[i][j]]; ok {
						return false
					}
					hash[board[i][j]] = board[i][j]
				}
			}
		}
	}
	return true
}

func TestLeetcode36(t *testing.T) {
	fmt.Println(searchRange([]int{1}, 1))
}
