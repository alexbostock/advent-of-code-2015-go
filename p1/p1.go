package p1

import (
	"io"
)

type Puzzle1 struct {
	instructions string
}

func (p *Puzzle1) LoadInput(input io.ReadSeeker) error {
	buffer, err := io.ReadAll(input)
	if err != nil {
		return err
	}
	p.instructions = string(buffer)
	return nil
}

func (p *Puzzle1) SolvePart1() interface{} {
	return countFloors(p.instructions)
}

func (p *Puzzle1) SolvePart2() interface{} {
	return firstEntersBasementAt(p.instructions)
}

func countFloors(instructions string) (count int) {
	for _, instruction := range []rune(instructions) {
		if instruction == '(' {
			count++
		} else {
			count--
		}
	}
	return
}

func firstEntersBasementAt(instructions string) int {
	currentFloor := 0
	for index, instruction := range []rune(instructions) {
		if instruction == '(' {
			currentFloor++
		} else {
			currentFloor--
			if currentFloor == -1 {
				return index + 1
			}
		}
	}
	panic("No solution")
}
