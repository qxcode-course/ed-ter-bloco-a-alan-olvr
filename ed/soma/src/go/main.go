package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func rec_sum(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Value + rec_sum(node.Left) + rec_sum(node.Right)
}

func rec_min(node *Node) int {
	if node == nil {
		return math.MaxInt
	}
	min := node.Value
	if l := rec_min(node.Left); l < min {
		min = l
	}
	if r := rec_min(node.Right); r < min {
		min = r
	}
	return min
}

// MyShow imprime a árvore binária de forma formatada.
func MyShow(node *Node, nivel int) {
	if node == nil {
		return
	}

	MyShow(node.Right, nivel+1)
	for i := 0; i < nivel; i++ {
		fmt.Print("    ")
	}
	fmt.Println(node.Value)
	MyShow(node.Left, nivel+1)
}

func BShow(node *Node, heranca string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, heranca+"l")
	}
	for i := 0; i < len(heranca)-1; i++ {
		if heranca[i] != heranca[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if heranca != "" {
		if heranca[len(heranca)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, heranca+"r")
	}
}

func create(parts *[]string) *Node {
	elem := (*parts)[0]
	*parts = (*parts)[1:]
	if elem == "#" {
		return nil
	}
	value, _ := strconv.Atoi(elem)
	node := &Node{Value: value}
	node.Left = create(parts)
	node.Right = create(parts)
	return node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	root := create(&parts)
	fmt.Println("Arvore:")
	BShow(root, "")
	fmt.Printf("Soma: %d, Minimo: %d\n", rec_sum(root), rec_min(root))
}
