package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func burnTrees(grid [][]rune, l, c int) {
	nl := len(grid)
	nc := len(grid[0])

	dirs := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	stack := NewStack[Pos]()
	stack.Push(Pos{l, c})

	for !stack.IsEmpty() {
		pos := stack.Pop()
		r, col := pos.l, pos.c

		if r < 0 || r >= nl || col < 0 || col >= nc {
			continue
		}
		if grid[r][col] != '#' {
			continue
		}

		grid[r][col] = 'o'

		for _, d := range dirs {
			stack.Push(Pos{r + d.l, col + d.c})
		}
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
