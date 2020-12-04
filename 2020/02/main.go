package main

import (
	"adventofcode-2020/util"
	"errors"
	"fmt"
	"strings"
)

func main() {
	util.RunIt("./02/input.txt", part1, part2)
}

func part1(lines []string) int64 {
	var valid int64

	for _, line := range lines {
		if IsValidPart1(line) {
			valid++
		}
	}

	return valid
}

func part2(lines []string) int64 {
	var valid int64

	for _, line := range lines {
		if IsValidPart2(line) {
			valid++
		}
	}
	return valid
}

func IsValidPart1(line string) bool {
	var from, to int
	var char, password string
	line = strings.ReplaceAll(line, "-", " ")
	line = strings.ReplaceAll(line, ":", "")
	_, err := fmt.Sscanf(line, "%d %d %s %s", &from, &to, &char, &password)
	util.HandleError(err)

	n := strings.Count(password, char)
	if n < from || n > to {
		return false
	}

	return true
}

func IsValidPart2(line string) bool {
	var from, to int
	var char, password string
	line = strings.ReplaceAll(line, "-", " ")
	line = strings.ReplaceAll(line, ":", "")
	_, err := fmt.Sscanf(line, "%d %d %s %s", &from, &to, &char, &password)
	util.HandleError(err)
	if len(password) < to {
		util.HandleError(errors.New("invalid password length"))
	}

	atFirst := string(password[from-1]) == char
	atSecond := string(password[to-1]) == char

	return atFirst != atSecond
}
