package p5

import (
	"bufio"
	"io"
)

type Puzzle5 struct {
	strings []string
}

func (p *Puzzle5) LoadInput(input io.ReadSeeker) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		p.strings = append(p.strings, scanner.Text())
	}
	return nil
}

func (p *Puzzle5) SolvePart1() interface{} {
	count := 0
	for _, str := range p.strings {
		if isNice([]rune(str)) {
			count++
		}
	}
	return count
}

func (p *Puzzle5) SolvePart2() interface{} {
	count := 0
	for _, str := range p.strings {
		if isNice2([]rune(str)) {
			count++
		}
	}
	return count
}

func isNice(str []rune) bool {
	numVowels := 0
	hasTwoInARow := false
	containsNaughtySubstring := false

	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	naughtyPairs := map[rune]rune{'a': 'b', 'c': 'd', 'p': 'q', 'x': 'y'}

	prev := (str)[0]
	if vowels[prev] {
		numVowels++
	}
	for _, char := range (str)[1:] {
		if vowels[char] {
			numVowels++
		}
		if prev == char {
			hasTwoInARow = true
		}
		if naughtyPairs[prev] == char {
			containsNaughtySubstring = true
			break
		}
		prev = char
	}
	return numVowels >= 3 && hasTwoInARow && !containsNaughtySubstring
}

func isNice2(str []rune) bool {
	pairs := make(map[rune]map[rune]int)

	hasSamePairTwice := false
	hasRepeatWithOneInBetween := false

	prevPrev := str[0]
	prev := str[1]
	pairs[prevPrev] = make(map[rune]int)
	pairs[prevPrev][prev] = 0
	for offset, char := range str[2:] {
		index := offset + 2
		if pairs[prev] == nil {
			pairs[prev] = make(map[rune]int)
			pairs[prev][char] = index - 1
		} else {
			previousIndex, isPresent := pairs[prev][char]
			if isPresent {
				hasSamePairTwice = hasSamePairTwice || previousIndex < index-2
			} else {
				pairs[prev][char] = index - 1
			}
		}

		if prevPrev == char {
			hasRepeatWithOneInBetween = true
		}

		prevPrev = prev
		prev = char
	}
	return hasSamePairTwice && hasRepeatWithOneInBetween
}
