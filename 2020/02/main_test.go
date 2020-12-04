package main

import "testing"

func TestIsValidPart1(t *testing.T) {
	tests := []struct {
		desc  string
		given string
		want  bool
	}{
		{"no occurrence", "1-3 a: b", false},
		{"not enough", "2-3 a: ab", false},
		{"too many", "2-3 a: aaaab", false},
		{"valid", "2-3 a: sdaab", true},
	}

	for _, test := range tests {
		t.Run(test.desc, func(tt *testing.T) {
			res := IsValidPart1(test.given)

			if res != test.want {
				if test.want {
					tt.Errorf("%s should be valid", test.given)
				} else {
					tt.Errorf("%s should be invalid", test.given)
				}
			}
		})
	}
}

func TestIsValidPart2(t *testing.T) {
	tests := []struct {
		desc  string
		given string
		want  bool
	}{
		{"valid first position", "1-3 a: abc", true},
		{"valid second position", "1-3 a: cba", true},
		{"invalid both positions", "1-3 a: aba", false},
		{"invalid no position", "1-3 a: bab", false},
	}

	for _, test := range tests {
		t.Run(test.desc, func(tt *testing.T) {
			res := IsValidPart2(test.given)

			if res != test.want {
				if test.want {
					tt.Errorf("%s should be valid", test.given)
				} else {
					tt.Errorf("%s should be invalid", test.given)
				}
			}
		})
	}
}
