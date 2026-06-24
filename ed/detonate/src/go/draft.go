package main
import (
    "bufio"
    "fmt"
    "os"
)

type Bomb struct {
    x, y, r int64
}

func canDetonate(a, b Bomb) bool {
    dx := a.x - b.x
    dy := a.y - b.y
    distSq := dx*dx + dy*dy
    rSq := a.r * a.r
    return distSq <= rSq
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    var n, d int
    fmt.Fscan(reader, &n, &d)
    _ = d

    bombs := make([]Bomb, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(reader, &bombs[i].x, &bombs[i].y, &bombs[i].r)
    }

    adj := make([][]bool, n)
    for i := 0; i < n; i++ {
        adj[i] = make([]bool, n)
        for j := 0; j < n; j++ {
            if i != j && canDetonate(bombs[i], bombs[j]) {
                adj[i][j] = true
            }
        }
    }

    maxCount := 0
    for start := 0; start < n; start++ {
        visited := make([]bool, n)
        visited[start] = true
        queue := []int{start}
        count := 1

        for len(queue) > 0 {
            cur := queue[0]
            queue = queue[1:]

            for next := 0; next < n; next++ {
                if !visited[next] && adj[cur][next] {
                    visited[next] = true
                    count++
                    queue = append(queue, next)
                }
            }
        }
        if count > maxCount {
            maxCount = count
        }
    }

    fmt.Println(maxCount)
}