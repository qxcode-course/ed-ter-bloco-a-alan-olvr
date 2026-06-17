package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	nrows := len(board)
	ncols := len(board[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= nrows || c < 0 || c >= ncols || board[r][c] != 'O' {
			return
		}
		board[r][c] = 'S'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r := 0; r < nrows; r++ {
		dfs(r, 0)
		dfs(r, ncols-1)
	}
	for c := 0; c < ncols; c++ {
		dfs(0, c)
		dfs(nrows-1, c)
	}

	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == 'S' {
				board[r][c] = 'O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
