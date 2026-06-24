package main
import (
    "fmt"
    "bufio"
    "os"
)

type Pos struct {
    l, c int 
}

func (p *Pos) getNeig() []Pos {
    return []Pos{
        {p.l - 1, p.c},
        {p.l + 1, p.c},
        {p.l, p.c - 1},
        {p.l, p.c + 1},
    }
}

func inside (grid [][]int, pos Pos) bool {
    nrows := len(grid)
    ncols := len(grid[0])
    return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func orangesRotting(grid [][]int) int {
    nl := len(grid)
    nc := len(grid[0])

    queue := []Pos{}
    fresh := 0

    for l := 0; l < nl; l++ {
        for c := 0; c < nc; c++ {
            switch grid[l][c] {
            case 2: 
                queue = append(queue, Pos{l, c})
            case 1: 
                fresh++
            }
        }
    }

    if fresh == 0 {
        return 0
    }

    minutes := 0

    for len(queue) > 0 {
        next := []Pos{}

        for _, cur := range queue {
            for _, neig := range cur.getNeig() {
                if inside(grid, neig) && grid[neig.l][neig.c] == 1 {
                    grid[neig.l][neig.c] = 2
                    fresh--
                    next = append(next, neig)
                }
            }
        }

        queue = next
        if len(next) > 0 {
            minutes++
        }
    }

    if fresh > 0 {
        return -1
    }

    return minutes
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

    var nl, nc int
    scanner.Scan()
    fmt.Sscanf(scanner.Text(), "%d %d", &nl, &nc)

    grid := make([][]int, nl)
    for i := 0; i < nl; i++ {
        grid[i] = make([]int, nc)
        scanner.Scan()
        line := scanner.Text()
        pos := 0
        for _, tok := range splitSpaces(line) {
            fmt.Sscanf(tok, "%d", &grid[i][pos])
            pos++
        }
    }

    fmt.Println(orangesRotting(grid))
}

func splitSpaces(s string) []string {
    var tokens []string
    cur := ""
    for _, r := range s {
        if r == ' ' || r == '\t' {
            if cur != "" {
                tokens = append(tokens, cur)
                cur = ""
            }
        } else {
            cur += string(r)
        }
    }
    if cur != "" {
        tokens = append(tokens, cur)
    }
    return tokens
}