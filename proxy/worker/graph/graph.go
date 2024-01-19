package graph

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type Node struct {
	ID    int
	Name  string
	Form  string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
	Links []*Node
}

var forms = map[int]string{
	0: "Circle",
	1: "Rect",
	2: "Square",
	3: "Ellipse",
	4: "Round Rect",
	5: "Rhombus",
}

func GetFigureBrackets(form string) string {
	switch form {	
	case "Circle":
		return "((" + form + "))"
	case "Ellipse":
		return "(" + form + ")"		
	case "Round Rect":
		return "(" + form + ")"	
	case "Rhombus":
		return "{" + form + "}"
	default:
		return "[" + form + "]"
	}
}

func GetMermaids(nodes []*Node) string {
	var res string
	for _, node := range nodes {
		for _, otherNode := range node.Links {
			res += fmt.Sprintf(
				"%s%s --> %s%s\n",
				node.Name,
				GetFigureBrackets(node.Form),
				otherNode.Name,
				GetFigureBrackets(otherNode.Form),
			)
		}
	}
	return res
}

func NewNodes(count int) []*Node {
	var nodes []*Node
	for i := 0; i < count; i++ {
		node := &Node{
			ID:   i,
			Name: gofakeit.FirstName(),
			Form: forms[gofakeit.Number(0, 5)],
		}
		nodes = append(nodes, node)
	}

	for i := 0; i < len(nodes); i++ {
		for j := i+1; j < len(nodes); j++ {
			// добавляем все предыдущие узлы
			if len(nodes[j].Links) < 1 {
				nodes[j].Links = append(nodes[j].Links, nodes[i])
			}

			// рандомная связь
			rnd := gofakeit.Number(0, 10)
			if rnd == 2 {
				nodes[i].Links = append(nodes[i].Links, nodes[j])
			}
		}
	}

	return nodes
}

func GetGraphPage() string {
	count := gofakeit.Number(5, 30)
	nodes := NewNodes(count)
	mermGraph := GetMermaids(nodes)
	return fmt.Sprintf(
		content,
		mermGraph,
	)
}

const content = `
---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

# Построение графа

Нужно написать воркер, который будет строить граф на текущей странице, каждые 5 секунд
От 5 до 30 элементов, случайным образом. Все ноды графа должны быть связаны.

Граф

{{< mermaid >}}
graph LR
%s
{{< /mermaid >}}
`
