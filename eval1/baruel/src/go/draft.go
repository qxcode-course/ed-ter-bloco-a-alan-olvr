package main

import "fmt"

type Colecao struct {
	TotalAlbum int
	Possuidas  []int
}

func main() {

	var c Colecao
	fmt.Scan(&c.TotalAlbum)

	var n int
	fmt.Scan(&n)
	c.Possuidas = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&c.Possuidas[i])
	}

	teveRepetida := false
	primeiraRepetida := true

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if c.Possuidas[i] == c.Possuidas[j] {
				if !primeiraRepetida {
					fmt.Print(" ")
				}
                fmt.Print(c.Possuidas[i])
                teveRepetida = true
				primeiraRepetida = false
				break
			}
		}
	}

    if !teveRepetida {
	    fmt.Print("N")
    }


	temFaltante := false
	primeiraFaltante := true
	for i := 1; i <= c.TotalAlbum; i++ {
		temFaltante = true
		for j := 0; j < n; j++ {
			if c.Possuidas[j] == i {
				temFaltante = false
				break
			}
		}
		if temFaltante {
			if !primeiraFaltante {
				fmt.Println(" ")
			}
			primeiraFaltante = false
		}
	}

	if !temFaltante {
		fmt.Println("N")
	}

}
