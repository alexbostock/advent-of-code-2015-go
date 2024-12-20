package p6

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Puzzle6 struct {
	instructions []instruction
}

type instruction struct {
	operation operation
	from      coords
	to        coords
}

type coords struct {
	x, y int
}

type operation int

const (
	toggle operation = iota
	on
	off
)

func (p *Puzzle6) LoadInput(input io.ReadSeeker) (err error) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		var instruction instruction
		switch line[0:7] {
		case "toggle ":
			instruction.operation = toggle
			err = parseCoordsRange(line[7:], &instruction)
		case "turn on":
			instruction.operation = on
			err = parseCoordsRange(line[8:], &instruction)
		case "turn of":
			instruction.operation = off
			err = parseCoordsRange(line[9:], &instruction)
		}
		if err != nil {
			return err
		}
		p.instructions = append(p.instructions, instruction)
	}
	return scanner.Err()
}

func parseCoordsRange(text string, instruction *instruction) (err error) {
	tokens := strings.Split(text, " ")
	instruction.from, err = parseCoords(tokens[0])
	if err != nil {
		return err
	}
	instruction.to, err = parseCoords(tokens[2])
	return err
}

func parseCoords(text string) (c coords, err error) {
	tokens := strings.Split(text, ",")
	x, err := strconv.Atoi(tokens[0])
	if err != nil {
		return
	}
	y, err := strconv.Atoi(tokens[1])
	return coords{x, y}, err
}

func (p *Puzzle6) SolvePart1() interface{} {
	lights := make([][]bool, 1000)
	for y := 0; y < 1000; y++ {
		lights[y] = make([]bool, 1000)
	}

	for _, instruction := range p.instructions {
		for y := instruction.from.y; y <= instruction.to.y; y++ {
			for x := instruction.from.x; x <= instruction.to.x; x++ {
				switch instruction.operation {
				case toggle:
					lights[y][x] = !lights[y][x]
				case on:
					lights[y][x] = true
				case off:
					lights[y][x] = false
				}
			}
		}
	}

	numLit := 0
	for _, row := range lights {
		for _, light := range row {
			if light {
				numLit++
			}
		}
	}
	return numLit
}

func (p *Puzzle6) SolvePart2() interface{} {
	lights := make([][]int, 1000)
	for y := 0; y < 1000; y++ {
		lights[y] = make([]int, 1000)
	}

	for _, instruction := range p.instructions {
		for y := instruction.from.y; y <= instruction.to.y; y++ {
			for x := instruction.from.x; x <= instruction.to.x; x++ {
				switch instruction.operation {
				case toggle:
					lights[y][x] += 2
				case on:
					lights[y][x]++
				case off:
					if lights[y][x] > 0 {
						lights[y][x]--
					}
				}
			}
		}
	}

	totalBrightness := 0
	for _, row := range lights {
		for _, light := range row {
			totalBrightness += light
		}
	}
	return totalBrightness
}
