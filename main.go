package main

import (
	"adventofcode2015/p1"
	"adventofcode2015/p2"
	"adventofcode2015/p3"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Puzzle interface {
	LoadInput(input io.ReadSeeker) error
	SolvePart1() interface{}
	SolvePart2() interface{}
}

func main() {
	puzzles := []Puzzle{
		&p1.Puzzle1{},
		&p2.Puzzle2{},
		&p3.Puzzle3{},
	}

	var puzzleNum int
	if len(os.Args) > 1 {
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		puzzleNum = num
	}

	if puzzleNum > 0 {
		solvePuzzle(puzzleNum, puzzles[puzzleNum-1])
	} else {
		for puzzleNum, puzzle := range puzzles {
			solvePuzzle(puzzleNum+1, puzzle)
		}
	}
}

func solvePuzzle(puzzleNum int, puzzle Puzzle) {
	fmt.Printf("Puzzle %d\n", puzzleNum)
	input, err := os.Open(fmt.Sprintf("./input/%d.txt", puzzleNum))
	if err != nil {
		panic(err)
	}
	err = puzzle.LoadInput(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(puzzle.SolvePart1())
	fmt.Println(puzzle.SolvePart2())
	fmt.Println()
}
