package utils

import "testing"

func TestStringInList(t *testing.T) {
	tests := []struct {
		s        string
		list     []string
		expected bool
	}{
		{s: "a", list: []string{"a", "b", "c"}, expected: true},
		{s: "s", list: []string{"a", "b", "c"}, expected: false},
		{s: "s", list: []string{}, expected: false},
		{s: "", list: []string{}, expected: false},
	}

	for _, test := range tests {
		r := StringInList(test.s, test.list)

		if r != test.expected {
			t.Errorf("result '%v' different from '%v'", r, test.expected)
		}
	}
}
