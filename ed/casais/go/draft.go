package main

import "fmt"

func main() {

	var n int
	fmt.Scanln(&n)
	var vet []int
	var usado []bool
	var qtdPares int

	for i := 0; i < n; i++ {

		var animal int
		fmt.Scanln(&animal)

		pos := -1

		for j := 0; j < len(vet); j++ {
			if !usado[j] && vet[j] == -animal {
				pos = j
				break
			}
		}

		if pos != -1 {
			usado[pos] = true
			qtdPares++
		} else {
			vet = append(vet, animal)
			usado = append(usado, false)
		}
	}

	fmt.Println(qtdPares)
}
