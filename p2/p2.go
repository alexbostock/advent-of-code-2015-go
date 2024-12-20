package p2

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Puzzle2 struct {
	parcels []Parcel
}

type Parcel struct {
	length, width, height int
}

func (p *Puzzle2) LoadInput(input io.ReadSeeker) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "x")
		var parcel Parcel
		for index, token := range tokens {
			num, err := strconv.Atoi(token)
			if err != nil {
				return err
			}
			switch index {
			case 0:
				parcel.length = num
			case 1:
				parcel.width = num
			case 2:
				parcel.height = num
			}
		}
		p.parcels = append(p.parcels, parcel)
	}
	return nil
}

func (p *Puzzle2) SolvePart1() interface{} {
	total := 0
	for _, parcel := range p.parcels {
		total += paperRequired(parcel.length, parcel.width, parcel.height)
	}
	return total
}

func (p *Puzzle2) SolvePart2() interface{} {
	total := 0
	for _, parcel := range p.parcels {
		total += ribbonRequired(parcel.length, parcel.width, parcel.height)
	}
	return total
}

func paperRequired(length, width, height int) int {
	face1 := length * width
	face2 := length * height
	face3 := width * height
	smallestFace := face1
	if face2 < smallestFace {
		smallestFace = face2
	}
	if face3 < smallestFace {
		smallestFace = face3
	}

	return 2*(face1+face2+face3) + smallestFace
}

func ribbonRequired(length, width, height int) int {
	perimeter1 := 2 * (length + width)
	perimeter2 := 2 * (length + height)
	perimeter3 := 2 * (width + height)
	minPerimeter := perimeter1
	if perimeter2 < minPerimeter {
		minPerimeter = perimeter2
	}
	if perimeter3 < minPerimeter {
		minPerimeter = perimeter3
	}
	return minPerimeter + length*width*height
}
