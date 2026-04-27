package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (set *Set) String() string {
	str := "["
	for i := 0; i < set.size; i++ {
		str += fmt.Sprintf("%d", set.data[i])
		if i < set.size-1 {
			str += ", "
		}
	}

	str += "]"
	return str
}

func (set *Set) reserve(newCapacity int) {
	if newCapacity <= set.capacity {
		return
	}

	newData := make([]int, set.size, newCapacity)
	for i := 0; i < set.size; i++ {
		newData[i] = set.data[i]
	}

	set.data = newData
	set.capacity = newCapacity
}

func (set *Set) binarySearch(value int) int {
	left, right := 0, set.size-1

	for left <= right {
		mid := (left + right) / 2

		if set.data[mid] == value {
			return mid
		} else if set.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func (set *Set) insert(value int, index int) error {
	if index < 0 || index > set.size {
		return fmt.Errorf("index out of range")
	}

	if set.size == set.capacity {
		newCap := set.capacity * 2
		if newCap == 0 {
			newCap = 1
		}
		set.reserve(newCap)
	}

	set.data = append(set.data, 0)
	for i := set.size; i > index; i-- {
		set.data[i] = set.data[i-1]
	}

	set.data[index] = value
	set.size++

	return nil
}

func (set *Set) erase(index int) error {
	if index < 0 || index >= set.size {
		return fmt.Errorf("index out of range")
	}

	for i := index; i < set.size-1; i++ {
		set.data[i] = set.data[i+1]
	}

	set.data = set.data[:set.size-1]
	set.size--

	return nil
}

func (set *Set) Insert(value int) {
	if set.binarySearch(value) != -1 {
		return
	}

	index := 0
	for index < set.size && set.data[index] < value {
		index++
	}

	set.insert(value, index)
}

func (set *Set) Contains(value int) bool {
	return set.binarySearch(value) != -1
}

func (set *Set) Erase(value int) bool {
	index := set.binarySearch(value)
	if index == -1 {
		return false
	}

	set.erase(index)
	return true
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !v.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
