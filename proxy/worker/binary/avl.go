package binary

import (
	"fmt"
	"math/rand"
)

var dfs func(root *Node)

type Node struct {
	Key    int
	Height int
	Left   *Node
	Right  *Node
}

func NewNode(key int) *Node {
	return &Node{
		Key:    key,
		Height: 1,
	}
}

type AVLTree struct {
	Root *Node
}

func (t *AVLTree) Insert(key int) {
	t.Root = insert(t.Root, key)
}

func (t *AVLTree) InsertBin(key int) {
	t.Root = insertBin(t.Root, key)
}

func (t *AVLTree) ToMermaid() string {
	var result string

	dfs = func(root *Node) {
		if root == nil {
			return
		}
		if root.Left != nil {
			result += fmt.Sprintf("%d --> %d\n", root.Key, root.Left.Key)
			dfs(root.Left)
		}
		if root.Right != nil {
			result += fmt.Sprintf("%d --> %d\n", root.Key, root.Right.Key)
			dfs(root.Right)
		}
	}

	dfs(t.Root)

	return result
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func updateHeight(node *Node) {
	node.Height = max(height(node.Left), height(node.Right)) + 1
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}

	return height(node.Left) - height(node.Right)
}

func leftRotate(x *Node) *Node {
	y := x.Right
	t := y.Left

	y.Left = x
	x.Right = t

	updateHeight(x)
	updateHeight(y)

	return y
}

func rightRotate(y *Node) *Node {
	x := y.Left
	t := x.Right

	x.Right = y
	y.Left = t

	updateHeight(y)
	updateHeight(x)

	return x
}

func insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}

	if key < node.Key {
		node.Left = insert(node.Left, key)
	} else if key > node.Key {
		node.Right = insert(node.Right, key)
	} else {
		return node
	}

	updateHeight(node)

	balance := getBalance(node)

	if balance > 1 { 
        if key < node.Left.Key { 
            return rightRotate(node) 
        } else if key > node.Left.Key { 
            node.Left = leftRotate(node.Left) 
            return rightRotate(node) 
        } 
    } 
  
    if balance < -1 { 
        if key > node.Right.Key { 
            return leftRotate(node) 
        } else if key < node.Right.Key { 
            node.Right = rightRotate(node.Right) 
            return leftRotate(node) 
        } 
    } 

	return node
}


func insertBin(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}

	if key < node.Key {
		node.Left = insertBin(node.Left, key)
	} else if key > node.Key {
		node.Right = insertBin(node.Right, key)
	} else {
		return node
	} 

	updateHeight(node)

	return node
}

func GenerateTree(count int) *AVLTree {
	avl := AVLTree{}
	for i := 0; i < count; i++ {
		key := rand.Intn(100)
		avl.Insert(key)
	}
	return &avl
}
