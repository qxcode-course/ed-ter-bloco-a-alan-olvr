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

func (ms *MultiSet) expand() {
	newCap := ms.capacity * 2
	ms.capacity = newCap
}

func (ms *MultiSet) search(value int) (bool, int) {
	low := 0
	high := ms.size - 1
	found := false
	result := 0

	for low <= high {
		mid := low + (high-low)/2

		if ms.data[mid] == value {
			found = true
			result = mid
			low = mid + 1
		} else if ms.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found {
		return true, result
	}

	return false, low
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
	_, index := ms.search(value)
	ms.insert(index, value)
}

func (ms *MultiSet) Contains(value int) bool {
	found, _ := ms.search(value)
	return found
}

func (ms *MultiSet) Erase(value int) bool {
	found, index := ms.search(value)
	if found {
		ms.erase(index)
		return true
	}
	return false
}

func (ms *MultiSet) Count(value int) int {
	qtdOcurrence := 0
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			qtdOcurrence++
		}
	}

	return qtdOcurrence
}

func (ms *MultiSet) Unique() int {
	if ms.size == 0 {
		return 0
	}

	cnt := 1
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			cnt++
		}
	}

	return cnt
}

func (ms *MultiSet) Clear() {
	ms.size = 0
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
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
