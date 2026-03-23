package main
import "fmt"

type Colecao struct {
	TotalAlbum int
	Possuidas  []int
}

func main() {
	var c Colecao
	var qtdPossui int

	fmt.Scan(&c.TotalAlbum)
	fmt.Scan(&qtdPossui)

	for i := 0; i < qtdPossui; i++ {
		var fig int
		fmt.Scan(&fig)
		c.Possuidas = append(c.Possuidas, fig)
	}

	teveRepetida := false
	primeiraRepetida := true

	for i := 1; i < len(c.Possuidas); i++ {
		if c.Possuidas[i] == c.Possuidas[i-1] {
			if !primeiraRepetida {
				fmt.Print(" ")
			}
			fmt.Print(c.Possuidas[i])
			teveRepetida = true
			primeiraRepetida = false
		}
	}

	if !teveRepetida {
		fmt.Print("N")
	}
	fmt.Println() 

	teveFaltante := false
	primeiraFaltante := true

	for i := 1; i <= c.TotalAlbum; i++ {
		encontrou := false
		for _, fig := range c.Possuidas {
			if fig == i {
				encontrou = true
				break
			}
		}

		if !encontrou {
			if !primeiraFaltante {
				fmt.Print(" ")
			}
			fmt.Print(i)
			teveFaltante = true
			primeiraFaltante = false
		}
	}

	if !teveFaltante {
		fmt.Print("N")
	}
	fmt.Println()
}

