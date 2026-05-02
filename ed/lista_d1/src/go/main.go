package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
}

type root struct {
	head *Node
}

func NewList() *root {
	list := &root{}
	list.head = &Node{}
	list.head.next = list.head
	list.head.prev = list.head
	return list
}

func insert(A *Node, value int) {
	B := A.prev
	C := &Node{
		Value: value,
		next:  A,
		prev:  B,
	}
	B.next = C
	A.prev = C
}

func (ll *root) String() string {
	str := "["

	current := ll.head.next
	for current != ll.head {
		str += fmt.Sprintf("%d", current.Value)

		if current.next != ll.head {
			str += ", "
		}

		current = current.next
	}

	str += "]"
	return str
}

func (ll *root) Size() int {
	count := 0
	current := ll.head.next

	for current != ll.head {
		count++
		current = current.next
	}

	return count
}

func (ll *root) Clear() {
	ll.head.next = ll.head
	ll.head.prev = ll.head
} 

func (ll *root) PushFront(value int) {
	insert(ll.head.next, value)
}

func (ll *root) PushBack(value int) {
	insert(ll.head, value)
}

func (ll *root) PopFront() {
	if ll.head.next == ll.head {
		return
	}

	first := ll.head.next
	second := first.next

	ll.head.next = second
	second.prev = ll.head
}

func (ll *root) PopBack() {
	if ll.head.prev == ll.head {
		return
	}

	last := ll.head.prev
	penultimate := last.prev

	ll.head.prev = penultimate
	penultimate.next = ll.head
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushBack(num)
			 }
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
