package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fila := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	var m int
	fmt.Scan(&m)

	saindo := make(map[int] bool, m) 
	for i := 0; i < m; i++ {
		var id int
		fmt.Scan(&id)
		saindo[id] = true
	}

	sep := ""
	for _, id := range fila {
		if !saindo[id] {
			fmt.Print(sep, id)
			sep = " "
		}
	}

	fmt.Println(" ")
}
