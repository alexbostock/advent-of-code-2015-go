package p1

import "testing"

func TestCountFloors(t *testing.T) {
	got := countFloors("(()(()(")
	expected := 3
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestFirstEntersBasementAt(t *testing.T) {
	got := firstEntersBasementAt("()())")
	expected := 5
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
