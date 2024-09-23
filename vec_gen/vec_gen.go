package vecgen

import "math/rand"

type Queue struct {
	Step uint
	Node *Queue
}

func NewQueue() Queue {
	return Queue{Step: 0, Node: nil}
}

func (q *Queue) Append() {
	head := q
	for head.Node != nil {
		head = head.Node
	}
	(*head).Node = &Queue{Node: nil, Step: (*head).Step + 1}
}

func RandVec(conteo int, max_val int) []int {
	vec := []int{}

	for i := 0; i < conteo; i++ {
		vec = append(vec, rand.Intn(max_val))
	}
	return vec
}
