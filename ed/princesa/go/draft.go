package main

import "fmt"

func proximoVivo(vivos []bool, pos, n int) int {
	pos = (pos + 1) % n
	for !vivos[pos] {
		pos = (pos + 1) % n
	}
	return pos
}

func imprimirVivos(elementos []int, vivos []bool, espada int) {
	fmt.Print("[ ")
	for i, v := range elementos {
		if !vivos[i] {
			continue
		}
		fmt.Print(v)
		if i == espada {
			fmt.Print(">")
		}
		fmt.Print(" ")
	}
	fmt.Println("]")
}

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
		imprimirVivos(elementos, vivos, espada)

		alvo := proximoVivo(vivos, espada, n)
		vivos[alvo] = false
		count--
		espada = proximoVivo(vivos, alvo, n)
	}

	for i, v := range elementos {
		if vivos[i] {
			fmt.Printf("[ %d> ]\n", v)
			break
		}
	}
}

func main() {
	var n, e int
	fmt.Scan(&n, &e)
	josephus(n, e)

}
