package p5

import "testing"

func TestIsNice(t *testing.T) {
	examples := []struct {
		str      string
		expected bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, example := range examples {
		got := isNice([]rune(example.str))
		if example.expected != got {
			t.Errorf("%v: expected %v, got %v", example.str, example.expected, got)
		}
	}
}

func TestIsNice2(t *testing.T) {
	examples := []struct {
		str      string
		expected bool
	}{

		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
		{"aaa", false},
		{"aaaa", true},
		{"abab", true},
	}

	for _, example := range examples {
		got := isNice2([]rune(example.str))
		if example.expected != got {
			t.Errorf("%v: expected %v, got %v", example.str, example.expected, got)
		}
	}
}
