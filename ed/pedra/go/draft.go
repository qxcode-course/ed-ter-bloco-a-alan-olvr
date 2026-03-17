package main

import "fmt"

func main() {

	var n int
	fmt.Scanln(&n)

	melhorIndice := -1
	melhorPontuacao := 1<<31 - 1


	for i := 0 ; i < n; i++ {

		var a, b int
		fmt.Scanln(&a, &b)

		if a < 10 ||  b < 10 {
			continue
		}

		pontuacao := a - b
		if pontuacao < 0 {
			pontuacao = -pontuacao
		}

		if pontuacao < melhorPontuacao {
			melhorPontuacao = pontuacao
			melhorIndice = i
		}	
	}

	if melhorIndice == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(melhorIndice)
	}
}
