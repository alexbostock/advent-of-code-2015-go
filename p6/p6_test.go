package p6

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	p := &Puzzle6{}
	p.LoadInput(strings.NewReader(`turn on 0,0 through 999,999
toggle 0,0 through 999,0
turn off 499,499 through 500,500
`))

	got := p.SolvePart1()
	expected := 1000000 - 1000 - 4
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
