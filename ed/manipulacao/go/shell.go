package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	result := []int{}
	for _, v := range vet {
		if v > 0 {
			result = append(result, v)
		}
	}
	return result
}

func getCalmWomen(vet []int) []int {
	result := []int{}
	for _, v := range vet {
		if v < 0 && -v < 10 {
			result = append(result, v)
		}
	}
	return result
}

func sortVet(vet []int) []int {
	result := make([]int, len(vet))
	copy(result, vet)
	sort.Ints(result)
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortStress(vet []int) []int {
	result := make([]int, len(vet))
	copy(result, vet)
	sort.Slice(result, func(i, j int) bool {
		return abs(result[i]) < abs(result[j])
	})
	return result
}

func reverse(vet []int) []int {
	result := make([]int, len(vet))
	for i, v := range vet {
		result[len(vet)-1-i] = v
	}
	return result
}

func unique(vet []int) []int {
	seen := map[int]bool{}
	result := []int{}
	for _, v := range vet {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func repeated(vet []int) []int {
	seen := map[int]bool{}
	result := []int{}
	for _, v := range vet {
		if seen[v] {
			result = append(result, v)
		}
		seen[v] = true
	}
	return result
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, v := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return []int{}
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	vet := make([]int, len(parts))
	for i, part := range parts {
		vet[i], _ = strconv.Atoi(part)
	}
	return vet
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}
