package main

import "fmt"

func valido(seq []byte, pos, L int) bool {
	if pos == len(seq) {
		return true
	}

	if seq[pos] != '.' {
		return valido(seq, pos+1, L)
	}

	for k := 0; k <= L; k++ {
		char := byte('0' + k)

		ok := true

		for j := 0; j < len(seq); j++ {
			if seq[j] == char {
				dist := pos - j
				if dist < 0 {
					dist = -dist
				}

				if dist <= L {
    				ok = false
   					break
				}
			}
		}

		if ok {
			seq[pos] = char

			if valido(seq, pos+1, L) {
				return true
			}

			seq[pos] = '.'
		}
	}

	return false
}

func main() {
	var sequence string
	var L int

	fmt.Scan(&sequence)
	fmt.Scan(&L)

	seq := []byte(sequence)

	valido(seq, 0, L)

	fmt.Println(string(seq))
}