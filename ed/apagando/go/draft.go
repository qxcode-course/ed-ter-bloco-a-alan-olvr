package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	fila := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	var m int
	fmt.Scan(&m)

	saindo := make(map[string]bool) 
	for i := 0; i < m; i++ {
		var id string
		fmt.Scan(&id)
		saindo[id] = true
	}

	resultado  := []string{}
	for _, id := range fila {
		if !saindo[id] {
			resultado = append(resultado, id)
		}
	}

	for i, id := range resultado {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(id)
	}
	fmt.Println(" ")
}
