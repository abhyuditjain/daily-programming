/*
This problem was asked by Google.

Given two singly linked lists that intersect at some point, find the intersecting node. The lists are non-cyclical.

For example, given A = 3 -> 7 -> 8 -> 10 and B = 99 -> 1 -> 8 -> 10, return the node with value 8.

In this example, assume nodes with the same value are the exact same node objects.

Do this in O(M + N) time (where M and N are the lengths of the lists) and constant space.
*/

package main

import (
	"errors"
	"fmt"
)

// Node is a Linked List node
type Node struct {
	Val  int
	Next *Node
}

func main() {
	head1 := &Node{1, nil}
	node1 := &Node{2, nil}
	node2 := &Node{3, nil}
	node3 := &Node{4, nil}
	node4 := &Node{5, nil}
	node5 := &Node{6, nil}

	head2 := &Node{1, nil}
	node6 := &Node{6, nil}

	head1.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	head2.Next = node6
	node6.Next = node4

	common, err := getIntersection(head1, head2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(common)
}

func getIntersection(head1, head2 *Node) (int, error) {
	length1 := head1.length()
	length2 := head2.length()

	if length1 > length2 {
		return getIntersectionLists(length1-length2, head1, head2)
	}
	return getIntersectionLists(length2-length1, head2, head1)
}

func getIntersectionLists(d int, head1, head2 *Node) (int, error) {
	current1 := head1
	current2 := head2

	for i := 0; i < d; i++ {
		if current1 == nil {
			return 0, errors.New("No common node")
		}
		current1 = current1.Next
	}

	for current1 != nil && current2 != nil {
		if current1 == current2 {
			return current1.Val, nil
		}
		current1 = current1.Next
		current2 = current2.Next
	}
	return 0, errors.New("No common node")
}

func (head *Node) length() int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}
