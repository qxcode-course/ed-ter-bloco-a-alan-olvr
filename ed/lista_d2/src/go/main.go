package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	Value int
	next *Node
	prev *Node
	root *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root

	return &LList{
		root: root,
		size: 0,
	}
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}

	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}

	return n.prev
}

func (ll *LList) Front() *Node {
	if ll.root.next == ll.root {
		return nil
	}

	return ll.root.next
}

func (ll *LList) Back() *Node {
	if ll.root.prev == ll.root {
		return nil
	}

	return ll.root.prev
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) Size() int{
	return ll.size
}

func (ll *LList) String() string {
	if ll.size == 0 {
		return "[]"
	}

	str := "["
	for node := ll.Front(); node != nil; node = node.Next() {
		str += fmt.Sprintf("%d", node.Value)

		if node.Next() != nil {
			str += ", "
		}
	}

	str += "]"
	return str
}

func (ll *LList) PopFront() {
	if ll.root.next == ll.root {
		return 
	}

	first := ll.root.next
	second := first.next

	ll.root.next = second
	second.prev = ll.root
}

func (ll *LList) PopBack() {
	if ll.root.prev == ll.root {
		return
	}

	last := ll.root.prev
	penultimate := last.prev

	ll.root.prev = penultimate
	penultimate.next = ll.root
}

func (ll *LList) PushFront(value int) {
	node := &Node{Value: value, root: ll.root}

	first := ll.root.next

	node.next = first
	node.prev = ll.root

	first.prev = node
	ll.root.next = node

	ll.size++
}

func (ll *LList) PushBack(value int) {
	node := &Node{Value: value, root: ll.root}

	last := ll.root.prev

	node.prev = last
	node.next = ll.root

	last.next = node
	ll.root.prev = node

	ll.size++
}

func (ll *LList) Search(value int) *Node {
	for node := ll.Front(); node != nil; node = node.Next() {
		if node.Value == value {
			return node
		}
	}

	return nil
}

func (ll *LList) Insert(node *Node, value int) {
	if node == nil {
		return
	}

	ref := &Node{Value: value, root: ll.root}

	prev := node.prev

	ref.prev = prev
	ref.next = node

	prev.next = ref
	node.prev = ref

	ll.size++
}

func (ll *LList) Remove(node *Node) *Node {
	if node == nil || node == ll.root {
		return nil
	}

	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	ll.size--

	if next == ll.root {
		return nil
	}
	return next
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

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
			// fmt.Println(ll.Size())
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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
			 	fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
			 	fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
			 	node.Value = newvalue
			} else {
			 	fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
			 	ll.Insert(node, newvalue)
			} else {
			 	fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
			 	ll.Remove(node)
			} else {
			 	fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
