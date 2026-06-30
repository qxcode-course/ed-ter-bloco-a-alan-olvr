package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {

    for i, row := range grid {
        filtered := row[:0]
        for _, b := range row {
            if b != ' ' {
                filtered = append(filtered, b)
            }
        }
        grid[i] = filtered
    }

	m, n := len(grid), len(grid[0])
	dirs := [][2]int{{0,1}, {0, -1}, {1, 0}, {-1, 0}}
	const INF = int(1e18)

	fireTime := make([][]int, m)
	for i := range fireTime {
		fireTime[i] = make([]int, n)
		for j := range fireTime[i] {
			fireTime[i][j] = INF
		}
	}

	queue := [][2]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				fireTime[i][j] = 0
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	for t := 1; len(queue) > 0; t++ {
    	next := [][2]int{}
        for _, cell := range queue {
        	for _, d := range dirs {
                ni, nj := cell[0]+d[0], cell[1]+d[1]
                if ni >= 0 && ni < m && nj >= 0 && nj < n &&
                    grid[ni][nj] != '2' && fireTime[ni][nj] == INF {
                    fireTime[ni][nj] = t
                    next = append(next, [2]int{ni, nj})
                }
            }
        }
        queue = next
    }

	canEscape := func(w int) bool {
        if fireTime[0][0] <= w {
            return false
        }
        visited := make([][]bool, m)
        for i := range visited {
            visited[i] = make([]bool, n)
        }
        visited[0][0] = true
        queue := [][2]int{{0, 0}}
        for step := 1; len(queue) > 0; step++ {
            next := [][2]int{}
            for _, cell := range queue {
                for _, d := range dirs {
                    ni, nj := cell[0]+d[0], cell[1]+d[1]
                    if ni < 0 || ni >= m || nj < 0 || nj >= n {
                        continue
                    }
                    if grid[ni][nj] == '2' || grid[ni][nj] == '1' || visited[ni][nj] {
                        continue
                    }
                    arrivalTime := w + step
                    if ni == m-1 && nj == n-1 {
                        if arrivalTime <= fireTime[ni][nj] {
                            return true
                        }
                        continue
                    }
                    if arrivalTime < fireTime[ni][nj] {
                        visited[ni][nj] = true
                        next = append(next, [2]int{ni, nj})
                    }
                }
            }
            queue = next
        }
        return false
    }

    if !canEscape(0) {
        return -1
    }
    if canEscape(int(1e9)) {
        return int(1e9)
    }
    lo, hi := 0, int(1e9)
    for lo < hi {
        mid := (lo + hi + 1) / 2
        if canEscape(mid) {
            lo = mid
        } else {
            hi = mid - 1
        }
    }
    return lo
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
