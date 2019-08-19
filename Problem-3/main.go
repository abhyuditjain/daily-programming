/*
This problem was asked by Google.

Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

type nodes []string

func serialize(root *treeNode) string {
	buffer := nodes([]string{})
	(&buffer).buildString(root)
	return strings.Join(buffer, "")
}

func (n *nodes) buildString(root *treeNode) {
	if root == nil {
		*n = append(*n, "X", ",")
	} else {
		*n = append(*n, strconv.Itoa(root.Val), ",")
		n.buildString(root.Left)
		n.buildString(root.Right)
	}
}

func deserialize(data string) *treeNode {
	dataSlice := strings.Split(data, ",")
	n := nodes(dataSlice)
	return (&n).buildTree()
}

func (n *nodes) buildTree() *treeNode {
	if len(*n) > 0 && (*n)[0] == "X" {
		return nil
	}

	data, _ := strconv.Atoi((*n)[0])
	t := &treeNode{data, nil, nil}
	*n = (*n)[1:]
	t.Left = n.buildTree()
	*n = (*n)[1:]
	t.Right = n.buildTree()
	return t
}

func main() {
	root := &treeNode{1, nil, nil}
	root.Left = &treeNode{2, nil, nil}
	root.Right = &treeNode{3, nil, nil}
	root.Left.Left = &treeNode{4, nil, nil}

	fmt.Println(serialize(root))
	s := serialize(root)
	fmt.Println(serialize(deserialize(s)))
}
