package p1

import (
	"io"
)

type Puzzle1 struct {
}

func (p *Puzzle1) LoadInput(input io.ReadSeeker) error {
	return nil
}

func (p *Puzzle1) SolvePart1() interface{} {
	return 0
}

func (p *Puzzle1) SolvePart2() interface{} {
	return 0
}
