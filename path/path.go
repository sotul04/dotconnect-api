package path

import (
	"fmt"
)

type Path struct {
	before *Path
	Row    int
	Col    int
	Length int
}

func (p *Path) Before() *Path {
	if p == nil {
		return nil
	}
	if v := p.before; v != nil {
		return v
	}
	return nil
}

func New(row, col int) *Path {
	p := new(Path)
	p.before = nil
	p.Length = 1
	p.Row = row
	p.Col = col
	return p
}

func (o *Path) ExtendPath(row, col int) *Path {
	p := new(Path)
	p.before = o
	p.Length = o.Length + 1
	p.Row = row
	p.Col = col
	return p
}

func (p *Path) IsSamePoint(row, col int) bool {
	return p.Row == row && p.Col == col
}

func (p *Path) Visited(row, col int) *Path {
	if p == nil {
		return nil
	}
	if p.Row == row && p.Col == col {
		return p
	}
	return p.before.Visited(row, col)
}

func (p *Path) ShowPath() {
	if p == nil {
		return
	}
	if p.before != nil {
		p.before.ShowPath()
	}
	fmt.Printf("Position: %d, row: %d, col: %d\n", p.Length, p.Row, p.Col)
}

func (p *Path) ToString() string {
	if p == nil {
		return ""
	}
	return p.before.ToString() + fmt.Sprintf("(%d,%d)", p.Row, p.Col)
}

func (p *Path) Surround(rowSize, colSize int) [][]int {
	var paths [][]int
	moves := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, move := range moves {
		newRow := p.Row + move[0]
		newCol := p.Col + move[1]
		if 0 <= newRow && 0 <= newCol && newRow < rowSize && newCol < colSize {
			paths = append(paths, []int{newRow, newCol})
		}
	}
	return paths
}

func (p *Path) ToPoints() [][]int {
	var points [][]int
	for p != nil {
		points = append([][]int{{p.Row, p.Col}}, points...)
		p = p.Before()
	}
	return points
}
