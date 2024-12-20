package p3

import (
	"io"
)

type Puzzle3 struct {
	directions string
}

type Coords struct {
	x, y int
}

func (p *Puzzle3) LoadInput(input io.ReadSeeker) error {
	buffer, err := io.ReadAll(input)
	if err != nil {
		return err
	}
	p.directions = string(buffer)
	return nil
}

func (p *Puzzle3) SolvePart1() interface{} {
	visited := make(map[Coords]bool)
	currentPos := Coords{}
	visited[currentPos] = true
	for _, direction := range []rune(p.directions) {
		switch direction {
		case '^':
			currentPos.y--
		case 'v':
			currentPos.y++
		case '>':
			currentPos.x++
		case '<':
			currentPos.x--
		}
		visited[currentPos] = true
	}
	return len(visited)
}

func (p *Puzzle3) SolvePart2() interface{} {
	visited := make(map[Coords]bool)
	posSanta := &Coords{}
	posRobot := &Coords{}
	visited[*posSanta] = true
	for index, direction := range []rune(p.directions) {
		pos := posSanta
		if index%2 == 1 {
			pos = posRobot
		}
		switch direction {
		case '^':
			pos.y--
		case 'v':
			pos.y++
		case '>':
			pos.x++
		case '<':
			pos.x--
		}
		visited[*posSanta] = true
		visited[*posRobot] = true
	}
	return len(visited)
}
