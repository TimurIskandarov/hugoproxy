package binary

import (
	"fmt"
	"math/rand"
)

func SetRandomNode(avl, bin *AVLTree) {
	num := rand.Intn(100)
	avl.Insert(num)
	bin.InsertBin(num)
}

func GetAVLPage(avl, bin *AVLTree) string {
	return fmt.Sprintf(
		content,
		bin.ToMermaid(),
		avl.ToMermaid(),
	)
}

const content = `
---
menu:
    after:
        name: binary_tree
        weight: 2
title: Построение сбалансированного бинарного дерева
---

# Задача построить сбалансированное бинарное дерево
Используя AVL дерево, постройте сбалансированное бинарное дерево, на текущей странице.

Нужно написать воркер, который стартует дерево с 5 элементов, и каждые 5 секунд добавляет новый элемент в дерево.

Каждые 5 секунд на странице появляется актуальная версия, сбалансированного дерева.

При вставке нового элемента, в дерево, нужно перестраивать дерево, чтобы оно оставалось сбалансированным.

Как только дерево достигнет 100 элементов, генерируется новое дерево с 5 элементами.

Бинарное дерево

{{< mermaid >}}
graph TD
%s
{{< /mermaid >}}

Сбалансированное бинарное дерево

{{< mermaid >}}
graph TD
%s
{{< /mermaid >}}
`
