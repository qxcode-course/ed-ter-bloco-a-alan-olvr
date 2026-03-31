package main

import "fmt"

func josephus(n int, e int) {

	elementos := make([]int, n)
	vivos := make([]bool, n)
	for i := 0; i < n; i++ {
		elementos[i] = i + 1
		vivos[i] = true
	}

	espada := e - 1
	count := n

	for count > 1 {
		fmt.Print("[ ")
		for i := 0; i < n; i++ {
			if vivos[i] {
				fmt.Printf("%d", elementos[i])
				if i == espada {
					fmt.Print(">")
				}
				fmt.Print(" ")
			}
		}
		fmt.Println("]")

		alvo := (espada + 1) % n
		for !vivos[alvo] {
			alvo = (alvo + 1) % n
		}

		vivos[alvo] = false
		count--

		espada = (alvo + 1) % n
		for !vivos[espada] {
			espada = (espada + 1) % n
		}
	}

	for i := 0; i < n; i++ {
		if vivos[i] {
			fmt.Printf("[ %d> ]\n", elementos[i])
			break
		}
	}
}

func main() {

	var n, e int
	fmt.Scan(&n, &e)

	josephus(n, e)

}
