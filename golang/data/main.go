// @author - Pratyay Ganguli
// this file will contain all the code being written in Vim for strengthening the DSA and problem solving skills
// Preparing for Google Senior Software Engineer position
// No need of writing the coverage you, can invoke all the functions inside main with proper comments.

package main

import (
	"fmt"
	"strconv"
	"log"
)

// execution/entry point of the program
func main() {
	fmt.Println("Hello, I am prepping for Senior SWE google!")
	if n, e := strToInt("Not a number"); e != nil {
		fmt.Printf("Cannot convert the value into integer: %v\n", e)
	} else {
		fmt.Printf("Converted number is %d", n)
	}
	s := intToStr(123)
	fmt.Printf("Converted string is %s", s)
	fmt.Println()
	// Insert at the head of the singly linked list
	InsertSLLNodeHead(10)
	InsertSLLNodeHead(20)
	DeleteSLLNodeHead()
	InsertSLLNodeHead(30)
	ReadSLLNode()
	// Insert at the tail of the singly linked list to make sure the logic is working as expected.
	InsertSLLNodeTail(50)
	InsertSLLNodeTail(60)
	DeleteSLLNodeTail()
	InsertSLLNodeTail(70)
	// Insert at the head of the doubly linked list
	InsertDLLNodeHead(15)
	InsertDLLNodeHead(25)
	InsertDLLNodeHead(35)
	ReadDLLNodeHeadToTail()
}

// let's write a program to convert a string into a number using golang
func strToInt(val string) (uint64, error) {
	if n, e := strconv.Atoi(val); e != nil {
		return 0, e
	} else {
		return uint64(n), nil
	}
}

// let's write a program to convert a int to a string using golang
func intToStr(num int) string {
	return strconv.Itoa(num)
}

// Write the logic for Singly List Node, it is the best way to start the data structure journey
type SLLNode struct {
	Data uint64
	Next *SLLNode
}

// Write the logic for Doubly List Node
type DLLNode struct {
	Data uint64
	Next *DLLNode
	Prev *DLLNode
}

var headSLLNode *SLLNode
var tailSLLNode *SLLNode

var headDLLNode *DLLNode
var tailDLLNode *DLLNode

// Insert the data at the head of the node
func InsertSLLNodeHead(data uint64) {
	// Check if the head node is null or not
	if headSLLNode == nil {
		headSLLNode = &SLLNode{Data: data}
		tailSLLNode = headSLLNode
	} else {
		newNode := &SLLNode{Data: data}
		newNode.Next = headSLLNode
		headSLLNode = newNode
	}
	// Print the head and the tail
	log.Printf("Head - %d; Tail - %d\n", headSLLNode.Data, tailSLLNode.Data)
}

// Insert the data at the head of the Doubly Linked List Node
func InsertDLLNodeHead(data uint64) {
	// check if the head of the Doubly linked list is null or not
	if headDLLNode == nil {
		headDLLNode = &DLLNode{Data: data}
		tailDLLNode = headDLLNode
		tailDLLNode.Prev = headDLLNode
	} else {
	// if the head is not null which means data exists in the linked list and the logic for insertion should be different
		tempNode := &DLLNode{Data: data}
		tempNode.Next = headDLLNode
		headDLLNode.Prev = tempNode
		headDLLNode = tempNode
	}
	// Enhance the logging, it should not only take the head and the tail data; it should be showing the reference of the prev-> head-> next too
	if headDLLNode.Next != nil && tailDLLNode.Prev != nil {
		log.Printf("Head - %d -> %d; %d <- Tail - %d;", headDLLNode.Data, headDLLNode.Next.Data, tailDLLNode.Prev.Data,tailDLLNode.Data)
		return	
	}
	log.Printf("Head - %d; Tail - %d;", headDLLNode.Data, tailDLLNode.Data)
}	

// Read the data inside the singly linked list
func ReadSLLNode() {
	// Check if the head or tail is empty, both of these should not be empty
	if headSLLNode == nil {
		log.Println("No elements present in the node")
		return
	} else {
		tempNode := headSLLNode
		for tempNode != nil {
			log.Printf("Element - %d\n", tempNode.Data)
			tempNode = tempNode.Next
		}
	}
}

// Read the data inside the doubly linked list
func ReadDLLNodeHeadToTail() {
	if headDLLNode == nil {
		log.Println("No elements present in the node")
	} else {
		tempNode := headDLLNode
		for tempNode != nil {
			if tempNode.Prev != nil && tempNode.Next != nil {
				log.Printf("%d -> %d -> %d", tempNode.Prev.Data, tempNode.Data, tempNode.Next.Data)
			} else if tempNode.Next != nil && tempNode.Prev == nil {
				log.Printf("x -> %d -> %d", tempNode.Data, tempNode.Next.Data)
			} else if tempNode.Prev	!= nil && tempNode.Next == nil {
				log.Printf("%d -> %d -> x", tempNode.Prev.Data, tempNode.Data)
			}
			tempNode = tempNode.Next
		}
	}
}

// Read the data inside the doubly linked list (inverted)
func ReadDLLNodeTailToHead() {
	// To be implemented
}

// Write the function to insert at the tail of the node
func InsertSLLNodeTail(data uint64) {
	if tailSLLNode == nil {
		log.Println("No data is present in the tail")
		tailSLLNode = &SLLNode{Data: data}
		headSLLNode = tailSLLNode
	} else {
		tempNode := &SLLNode{Data: data}
		tailSLLNode.Next = tempNode
		tailSLLNode = tempNode
	}
	log.Printf("Head: %d; Tail: %d\n", headSLLNode.Data, tailSLLNode.Data)
}

// Write the function for deleting from the head
func DeleteSLLNodeHead() {
	// Check if the head is nil
	if headSLLNode == nil {
		log.Println("No elements present in the linked list")
		return
	}
	// Check if the linked list has only one element
	if headSLLNode.Next == nil {
		log.Println("Can't remove the only element present in the linked list")
		return
	}
	// Else just simply remove the element from the head; change the reference of the head
	log.Printf("Popping %d out", headSLLNode.Data)
	headSLLNode = headSLLNode.Next
}

// Write the function for deleting from the tail
func DeleteSLLNodeTail() {
	// Business logic
	if tailSLLNode == nil {
		log.Println("No elements present in the linked list")
		return
	}
	// check if only one element is present in the linked list if that is the case, do not remove anything
	if headSLLNode == tailSLLNode {
		log.Println("Only one element present in the linked list, can't remove the element")
		return
	}
	// else just change the reference
	// intiate the tempNode with head as starting reference
	tempNode := headSLLNode

	// traverse through the entire list
	for tempNode != nil {
		// check if we have reached the node prev to tail node
		if tempNode.Next == tailSLLNode {
			log.Printf("Popping %d out", tailSLLNode.Data)
			tailSLLNode = tempNode
			tailSLLNode.Next = nil
			return
		}
		tempNode = tempNode.Next
	}
}

