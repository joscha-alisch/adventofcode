package main

import (
	"adventofcode-2020/util"
	"strings"
)

func main() {
	util.RunIt("./04/input.txt", part1, part2)
}

func part1(lines []string) int64 {
	passports := ToPassports(lines)
	validators := []Validator{
		{Key: "byr", Validate: exists},
		{Key: "iyr", Validate: exists},
		{Key: "eyr", Validate: exists},
		{Key: "hgt", Validate: exists},
		{Key: "hcl", Validate: exists},
		{Key: "ecl", Validate: exists},
		{Key: "pid", Validate: exists},
	}

	return CountValid(passports, validators)
}

func part2(lines []string) int64 {
	passports := ToPassports(lines)
	validators := []Validator{
		{Key: "byr", Validate: inRange(1920, 2002)},
		{Key: "iyr", Validate: inRange(2010, 2020)},
		{Key: "eyr", Validate: inRange(2020, 2030)},
		{Key: "hgt", Validate: anyOf(inRangeWithSuffix("cm", 150, 193), inRangeWithSuffix("in", 59, 76))},
		{Key: "hcl", Validate: colorCode},
		{Key: "ecl", Validate: anyOf(eq("amb"), eq("blu"), eq("brn"), eq("gry"), eq("grn"), eq("hzl"), eq("oth"))},
		{Key: "pid", Validate: allOf(number, length(9))},
	}

	return CountValid(passports, validators)
}

type Passport map[string]string

func ToPassports(lines []string) []Passport {
	var current string
	var passports []Passport

	for i, line := range lines {
		current += " " + line

		if line == "" || i == len(lines)-1 {
			m := make(map[string]string)
			current = strings.Trim(current, " ")
			kvs := strings.Split(current, " ")
			for _, kv := range kvs {
				v := strings.Split(kv, ":")
				m[v[0]] = v[1]
			}

			passports = append(passports, m)
			current = ""
		}
	}
	return passports
}
