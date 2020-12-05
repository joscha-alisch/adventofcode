package main

import "testing"

func TestExists(t *testing.T) {
	run([]validatorTest{
		{"does not exist", "", false, exists},
		{"does exist", "test", true, exists},
	}, t)
}

func TestInRange(t *testing.T) {
	run([]validatorTest{
		{"is in range", "2", true, inRange(0, 10)},
		{"is equal minimum", "0", true, inRange(0, 10)},
		{"is equal maximum", "10", true, inRange(0, 10)},
		{"is under range", "1", false, inRange(5, 10)},
		{"is over range", "11", false, inRange(5, 10)},
		{"is not a number", "bla", false, inRange(5, 10)},
	}, t)
}

func TestAnyOf(t *testing.T) {
	run([]validatorTest{
		{"none valid", "2", false, anyOf(invalid, invalid, invalid)},
		{"all valid", "0", true, anyOf(valid, valid)},
		{"one valid", "10", true, anyOf(invalid, valid)},
	}, t)
}

func TestInRangeWithSuffix(t *testing.T) {
	run([]validatorTest{
		{"is in range", "5cm", true, inRangeWithSuffix("cm", 0, 10)},
		{"is equal min", "0cm", true, inRangeWithSuffix("cm", 0, 10)},
		{"is equal max", "10cm", true, inRangeWithSuffix("cm", 0, 10)},
		{"under range", "0cm", false, inRangeWithSuffix("cm", 5, 10)},
		{"over range", "20cm", false, inRangeWithSuffix("cm", 0, 10)},
		{"no suffix", "10", false, inRangeWithSuffix("cm", 0, 10)},
		{"wrong suffix", "10gb", false, inRangeWithSuffix("cm", 0, 10)},
	}, t)
}

func TestColorCode(t *testing.T) {
	run([]validatorTest{
		{"valid code", "#a0B7ff", true, colorCode},
		{"no prefix", "a0B7ff", false, colorCode},
		{"too long", "#a0B7fff", false, colorCode},
		{"not hex", "#xxxxxx", false, colorCode},
	}, t)
}

func TestEq(t *testing.T) {
	run([]validatorTest{
		{"is equal", "test", true, eq("test")},
		{"not equal", "tes", false, eq("test")},
	}, t)
}

func TestAllOf(t *testing.T) {
	run([]validatorTest{
		{"all valid", "test", true, allOf(valid, valid, valid)},
		{"one invalid", "test", false, allOf(valid, invalid, valid)},
	}, t)
}

func TestNumber(t *testing.T) {
	run([]validatorTest{
		{"valid", "002010581284", true, number},
		{"includes character", "812793a123123", false, number},
	}, t)
}

func TestLength(t *testing.T) {
	run([]validatorTest{
		{"correct length", "1234", true, length(4)},
		{"wrong length", "12345", false, length(4)},
		{"empty", "", true, length(0)},
	}, t)
}

type validatorTest struct {
	desc  string
	s     string
	want  bool
	valid ValidationFunc
}

func run(tests []validatorTest, t *testing.T) {
	for _, test := range tests {
		testValidator(test.desc, test.s, test.want, test.valid, t)
	}
}
func testValidator(testname string, s string, want bool, valid ValidationFunc, t *testing.T) {
	t.Run(testname, func(tt *testing.T) {
		if valid(s) != want {
			if want {
				tt.Error("should be valid")
			} else {
				tt.Error("should be invalid")
			}
		}
	})
}

func invalid(s string) bool {
	return false
}

func valid(s string) bool {
	return true
}
