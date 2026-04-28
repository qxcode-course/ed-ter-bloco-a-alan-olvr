package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (ms *MultiSet) String() string {
	str := "["

	for i := 0; i < ms.size; i++ {
		str += fmt.Sprintf("%d", ms.data[i])
		if i < ms.size-1 {
			str += ", "
		}
	}

	str += "]"
	return str
}

func (ms *MultiSet) insert(index int, value int) error {
	if index < 0 || index > ms.size {
		return fmt.Errorf("index out of range")
	}

	if ms.size == ms.capacity {
		newCap := ms.capacity * 2
		if newCap == 0 {
			newCap = 1
		}

		newData := make([]int, ms.size, newCap)
		for i := 0; i < ms.size; i++ {
			newData[i] = ms.data[i]
		}

		ms.data = newData
		ms.capacity = newCap
	}

	ms.data = append(ms.data, 0)
	for i := ms.size; i > index; i-- {
		ms.data[i] = ms.data[i-1]
	}

	ms.data[index] = value
	ms.size++
	return nil
}

func (ms *MultiSet) erase(index int) error {
	if index < 0 || index >= ms.size {
		return nil
	}

	for i := index; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}

	ms.data = ms.data[:ms.size-1]
	ms.size--

	return nil
}

func (ms *MultiSet) Insert(value int) {
	index := 0
	for index < ms.size && ms.data[index] < value {
		index++
	}

	ms.insert(index, value)
}

func (ms *MultiSet) Contains(value int) bool {
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			return true
		}
	}

	return false
}

func (ms *MultiSet) Erase(value int) bool {
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			ms.erase(i)
			return true
		}

		if ms.data[i] > value {
			return false
		}
	}

	return false
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			if !ms.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println(true)
			} else {
				fmt.Println(false)
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
