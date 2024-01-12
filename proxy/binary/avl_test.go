package binary

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_height(t *testing.T) {
	node := new(Node)
	got := height(node)
	if got != 0 {
		t.Errorf("Ожидалось, что высота nil узла будет 0, получено %v", got)
	}

	node = &Node{Key: 10, Height: 5}
	got = height(node)
	if got != 5 {
		t.Errorf("Ожидалось, что height(node)=3, получено %v", got)
	}
}

func Test_max(t *testing.T) {
	got := max(10, 5)
	if got != 10 {
		t.Errorf("Ожидалось, что max(10,5)=10, получено %d", got)
	}	
	got = max(5, 10)
	if got != 10 {
		t.Errorf("Ожидалось, что max(5,10)=10, получено %d", got)
	}	
	got = max(5, 5)
	if got != 5 {
		t.Errorf("Ожидалось, что max(5,5)=5, получено %d", got)
	}
}

func Test_updateHeight(t *testing.T) {
	node := &Node{Key: 40}
	node.Left = &Node{Key: 30}
	node.Right = &Node{Key: 50}

	updateHeight(node.Left)  // высота левого поддерева - 1
	updateHeight(node.Right) // высота правого поддерева - 1
	updateHeight(node)       // высота всего дерева - 2

	want := 2
	if node.Height != want {
		t.Errorf("Ожидалось, что высота узла будет %d, получено %d", want, node.Height)
	}
}

func Test_getBalance(t *testing.T) {
	node := &Node{Key: 40}
	node.Left = &Node{Key: 30}
	node.Right = &Node{Key: 50}

	updateHeight(node.Left)
	updateHeight(node.Right)

	got := getBalance(node)

	want := 1 - 1
	if got != want {
		t.Errorf("Ожидалось, что баланс узла будет %v, получено %v", want, got)
	}
}

func Test_leftRotate(t *testing.T) {
	x := &Node{Key: 30}
	x.Right = &Node{Key: 50}
	x.Right.Left = &Node{Key: 40}

	got := leftRotate(x)

	if reflect.DeepEqual(got, x.Right) {
		t.Errorf("Ожидалось, что функция вернет правое поддерево x, получено %v", got)
	}

	if got.Left != x {
		fmt.Println(got)
		t.Errorf("Ожидалось, что узел x станет левым поддеревом узла y")
	}
}

func Test_rightRotate(t *testing.T) {
	y := &Node{Key: 50}
	y.Left = &Node{Key: 30}
	y.Left.Right = &Node{Key: 40}

	got := rightRotate(y)

	if got.Right != y {
		t.Errorf("Ожидалось, что функция вернет левое поддерево y, получено %v", got)
	}

	if got.Right != y {
		t.Errorf("Ожидалось, что узел y станет правым поддеревом узла x")
	}
}
