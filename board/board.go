package board

import (
	"dot-connect/path"
)

type Size struct {
	Row int
	Col int
}

func NewSize(row, col int) *Size {
	s := new(Size)
	s.Row = row
	s.Col = col
	return s
}

func (s *Size) IsIndexValid(row, col int) bool {
	return 0 <= row && row < s.Row && 0 <= col && col < s.Col
}

func CopyArray(original [][]int) [][]int {
	copied := make([][]int, len(original))
	for i := range original {
		copied[i] = make([]int, len(original[i]))
		copy(copied[i], original[i])
	}

	return copied
}

type Board struct {
	Board [][]int
	Size  *Size
	Step  *path.Path
	Count int //block Count
}

func NeWBoard(board [][]int, size *Size, step *path.Path) *Board {
	b := new(Board)
	b.Board = board
	b.Size = size
	b.Step = step
	count := 0
	for _, rows := range board {
		for _, cell := range rows {
			if cell == 1 {
				count++
			}
		}
	}
	b.Count = count
	return b
}

func (b *Board) ExtendBoard(step *path.Path) *Board {
	nb := new(Board)
	nb.Board = b.Board
	nb.Count = b.Count
	nb.Size = b.Size
	nb.Step = step
	return nb
}

func (b *Board) IsValidMove(row, col int) bool {
	checkBoard := CopyArray(b.Board)
	step := b.Step
	for step != nil {
		checkBoard[step.Row][step.Col] = 2
		step = step.Before()
	}
	checkBoard[row][col] = 3
	moves := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	countSingleWay := 0
	for i := range checkBoard {
		for j := range checkBoard[i] {
			cell := checkBoard[i][j]
			if cell == 1 || cell == 2 || cell == 3 {
				continue
			}
			countWay := 0
			for _, move := range moves {
				nrow := i + move[0]
				ncol := j + move[1]
				if b.Size.IsIndexValid(nrow, ncol) && (checkBoard[nrow][ncol] == 0 || checkBoard[nrow][ncol] == 3) {
					countWay++
				}
			}
			if countWay == 0 {
				// fmt.Println("Zeroo")
				return false
			} else if countWay == 1 {
				// fmt.Println(i, j, b.Step)
				countSingleWay++
				if countSingleWay == 2 {
					// fmt.Println("More than 2")
					return false
				}
			}
		}
	}
	return true
}

func (b *Board) IsNextBlocked(row, col int) bool {
	return b.Board[row][col] == 1
}

// preq: IsValidMove (row, col)
func (b *Board) Visit(row, col int) *Board {
	visited := b.Step.Visited(row, col)
	if visited == nil {
		visited = b.Step.ExtendPath(row, col)
	}
	return b.ExtendBoard(visited)
}

func (b *Board) IsFinished() bool {
	return b.Step.Length >= b.Size.Row*b.Size.Col-b.Count
}
