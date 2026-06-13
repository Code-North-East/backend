// Stack mechanism
// First In Last out
// Insertion should be at head and deletion also in head

package data

import "fmt"

// stack node structure
type Node struct {
	Val  int64
	Next *Node
}

// delcare the head and tail
var HEAD *Node = nil
var TAIL *Node = nil

func PushStackS(data int64) {
	if HEAD == nil {
		HEAD = &Node{
			Val: data,
		}
		// Tail will be nil if head is nil
		TAIL = HEAD
		return
	}

	n := &Node{
		Val:  data,
		Next: HEAD,
	}

	if TAIL == HEAD {
		HEAD.Next = TAIL
	}
	HEAD = n
	fmt.Printf("HEAD : %d and TAIL: %d", HEAD.Val, TAIL.Val)
	fmt.Println("")
}

func PopStackS() {
	// TODO: implement
}
