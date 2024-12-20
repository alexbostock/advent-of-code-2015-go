package p4

import "testing"

func TestSearchHashWithZeros(t *testing.T) {
	got := searchHashWithZeros("abcdef", 5)
	expected := "609043"
	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
