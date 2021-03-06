package main

import "fmt"

func main() {
	t := NewBinaryTree(2)
	t.SetLeft(7)
	t.Left().SetLeft(2)
	t.Left().SetRight(6)
	t.Left().Right().SetLeft(5)
	t.Left().Right().SetRight(11)
	t.SetRight(5)
	t.Right().SetRight(9)
	t.Right().Right().SetLeft(4)

	fmt.Println(t.Right().Right().Get()) // 9
	fmt.Println(t.Left().Left().Get())   // 2
}

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(data int) *BinaryTree {
	return &BinaryTree{root: &Node{Data: data}}
}

func (t *BinaryTree) Get() int {
	return t.root.Data
}

func (t *BinaryTree) Left() *BinaryTree {
	return &BinaryTree{root: t.root.Left}
}

func (t *BinaryTree) Right() *BinaryTree {
	return &BinaryTree{root: t.root.Right}
}

func (t *BinaryTree) SetLeft(data int) {
	t.root.Left = &Node{Data: data}
}

func (t *BinaryTree) SetRight(data int) {
	t.root.Right = &Node{Data: data}
}
