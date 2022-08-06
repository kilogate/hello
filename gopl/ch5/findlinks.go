package main

import "fmt"

type Node struct {
	Val      string
	Children []*Node
}

func main() {
	node11 := &Node{
		Val:      "11",
		Children: nil,
	}

	node12 := &Node{
		Val:      "12",
		Children: nil,
	}

	node1 := &Node{
		Val:      "1",
		Children: []*Node{node11, node12},
	}

	outline(nil, node1)
}

// outline：注意只有入栈没有出栈，这和切片的特性有关
func outline(stack []string, n *Node) {
	if n != nil {
		stack = append(stack, n.Val)
		fmt.Println(stack)
	}
	for _, c := range n.Children {
		outline(stack, c)
	}
}
