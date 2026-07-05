package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	cima := Pos{p.l - 1, p.c}
	baixo := Pos{p.l + 1, p.c}
	esquerda := Pos{p.l, p.c - 1}
	direita := Pos{p.l, p.c + 1}

	return []Pos{cima, baixo, esquerda, direita}
}

func inside(grid [][]rune, pos Pos) bool {
	qtdLinhas := len(grid)
	qtdColunas := len(grid[0])

	if pos.l < 0 || pos.l >= qtdLinhas {
		return false
	}
	if pos.c < 0 || pos.c >= qtdColunas {
		return false
	}
	return true
}

func match(grid [][]rune, pos Pos, char rune) bool {
	if !inside(grid, pos) {
		return false
	}
	return grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos) {
	veioDe := map[Pos]Pos{}
	explorado := map[Pos]bool{}

	fila := NewQueue[Pos]()
	fila.Enqueue(startPos)
	explorado[startPos] = true

	for !fila.IsEmpty() {
		noAtual, _ := fila.Dequeue()

		if noAtual == endPos {
			trilha := make([]Pos, 0)
			no := noAtual
			for no != startPos {
				trilha = append(trilha, no)
				no = veioDe[no]
			}
			trilha = append(trilha, startPos)
			for _, ponto := range trilha {
				grid[ponto.l][ponto.c] = '.'
			}

			return
		}

		vizinhos := noAtual.getNeig()
		for _, v := range vizinhos {
			if explorado[v] {
				continue
			}
			ehCaminhoLivre := match(grid, v, ' ')
			ehDestino := v == endPos
			if ehCaminhoLivre || ehDestino {
				explorado[v] = true
				veioDe[v] = noAtual
				fila.Enqueue(v)
			}
		}
	}
}

func voltar() {}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
