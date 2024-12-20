package p2

import "testing"

func TestPaperRequired(t *testing.T) {
	got := paperRequired(2, 3, 4)
	expected := 58
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestRibbonRequired(t *testing.T) {
	got := ribbonRequired(2, 3, 4)
	expected := 34
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
