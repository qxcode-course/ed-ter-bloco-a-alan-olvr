package main

import "fmt"

type Ponto struct {
	X int
	Y int
}

func main() {
	var q int
	var direcao string

	fmt.Scan(&q, &direcao)
	cobra := make([]Ponto, q)

	for i := 0; i < q; i++ {
		fmt.Scan(&cobra[i].X, &cobra[i].Y)
	}

	for i := q - 1; i > 0; i-- {
		cobra[i] = cobra[i-1]
	}

	switch direcao {
	case "L":
		cobra[0].X--
	case "R":
		cobra[0].X++
	case "U":
		cobra[0].Y--
	case "D":
		cobra[0].Y++
	}

	for i := 0; i < q; i++ {
		fmt.Printf("%d %d\n", cobra[i].X, cobra[i].Y)
	}
}