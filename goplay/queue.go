package main

import (
	"fmt"
)

type Node struct {
	X int
	Y int
}

var G = make(map[Node]int)

func main() {

	n := Node{1, 2}
	G[n] = 4

	for k, v := range G {
		fmt.Println(k, v)
	}

	m := Node{1, 2}
	v, ok := G[m]
	fmt.Println(v, ok)

	//q := new(Queue)
	q := make(Queue, 0)
	q.Push(&m)
	q.Push(&Node{2, 3})
	q.Push(&n)
	//q.Pop()
	fmt.Println(q)
	for _, e := range q {
		fmt.Printf("%+v", *e)
	}
}

type Queue []*Node

func (q *Queue) Push(n *Node) {
	*q = append(*q, n)
}

func (q *Queue) Pop() (n *Node) {
	n = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) Len() int {
	return len(*q)
}

type Stack []*Node

func (q *Stack) Push(n *Node) {
	*q = append(*q, n)
}

func (q *Stack) Pop() (n *Node) {
	x := q.Len() - 1
	n = (*q)[x]
	*q = (*q)[:x]
	return
}
func (q *Stack) Len() int {
	return len(*q)
}
