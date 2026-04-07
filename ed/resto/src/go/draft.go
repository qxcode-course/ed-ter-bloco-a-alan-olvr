package main
import "fmt"

func divisao(n int) {
    if n == 0 {
        return
    } 

    divisao(n / 2) 
    fmt.Println(n/2, n%2)
}

func main() {
    var n int
    fmt.Scan(&n)
    divisao(n)
}
