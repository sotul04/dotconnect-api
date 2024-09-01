package solver

import (
	"dot-connect/board"
	"dot-connect/path"
	"dot-connect/queue"
	"fmt"
)

type Solver struct {
	Board       *board.Board
	Solution    *path.Path
	CounterNode int
	Found       bool
}

func New(board *board.Board) *Solver {
	s := new(Solver)
	s.Board = board
	s.Solution = nil
	s.CounterNode = 0
	s.Found = false
	return s
}

func (s *Solver) Solve() {
	queue := queue.NewQueue()
	visited := make(map[string]struct{})

	// heap.Push(&queue, prioqueue.NewItem(s.Board))
	queue.Enqueue(s.Board)
	visited[s.Board.Step.ToString()] = struct{}{}
	fmt.Println("Start the search")
	for !queue.IsEmpty() && !s.Found {
		currentBoard := queue.Dequeue()
		s.CounterNode++
		if currentBoard.IsFinished() {
			s.Found = true
			s.Solution = currentBoard.Step
			continue
		}
		nextSteps := currentBoard.Step.Surround(currentBoard.Size.Row, currentBoard.Size.Col)
		for _, item := range nextSteps {
			if !currentBoard.IsNextBlocked(item[0], item[1]) && currentBoard.IsValidMove(item[0], item[1]) {
				newStep := currentBoard.Step.ExtendPath(item[0], item[1])
				step := newStep.ToString()
				if _, exitst := visited[step]; !exitst {
					newBoard := currentBoard.Visit(item[0], item[1])
					queue.Enqueue(newBoard)
					visited[step] = struct{}{}
				}
			}
		}
	}
	fmt.Println("Done")
}
