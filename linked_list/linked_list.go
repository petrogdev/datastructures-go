package linked_list

import "fmt"

// LinkedList
type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Next *Node
	Data interface{}
}

// New
func New() *LinkedList {
	newNode := &Node{
		Next: nil,
		Data: nil,
	}
	return &LinkedList{
		Head: newNode,
		Tail: newNode,
	}
}

// IsEmpty
func (ll *LinkedList) IsEmpty() bool {
	if ll.Head == nil {
		return true
	}
	return false
}

// Append
func (ll *LinkedList) Append(d interface{}) *LinkedList {
	nextNode := &Node {
		Next: nil,
		Data: d,
	}

	if ll.Head.Data == nil {
		ll.Head = nextNode
	} else {
		ll.Tail.Next = nextNode
	}

	ll.Tail = nextNode
	return ll
}

func (ll *LinkedList) DeleteWithValue(v interface{}) *LinkedList {
	var node = ll.Head
	if node.Data == v {
		ll.Head = ll.Head.Next
		return ll
	}
	for {
		if v == node.Next.Data {
			if node.Next.Next != nil {
				node.Next = node.Next.Next
				break
			}
			node.Next = nil
			break
		}
		node = node.Next
	}
	return ll
}

// PrintAll : Prints all elements of the Linked List
func (ll *LinkedList) PrintAll() {
	var node = ll.Head
	for {
		fmt.Println(node.Data)
		if node.Next == nil {
			return
		}
		node = node.Next
	}
}