package queue

import (
	"dot-connect/board"
)

type Node struct {
	data *board.Board
	next *Node
}

// Queue represents a queue implemented using a linked list
type Queue struct {
	head *Node
	tail *Node
	size int
}

// NewQueue creates a new empty queue
func NewQueue() *Queue {
	q := new(Queue)
	q.head = nil
	q.tail = nil
	q.size = 0
	return q
}

// Enqueue adds a new Board to the end of the queue
func (q *Queue) Enqueue(board *board.Board) {
	newNode := &Node{data: board}
	if q.tail != nil {
		q.tail.next = newNode
	}
	q.tail = newNode
	if q.head == nil {
		q.head = newNode
	}
	q.size++
}

// Dequeue removes and returns the Board at the front of the queue
func (q *Queue) Dequeue() *board.Board {
	if q.head == nil {
		return nil
	}
	board := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return board
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return q.size
}
