package main

import (
	"adventofcode-2020/util"
)

func main() {
	util.RunIt("./01/input.txt", part1, part2)
}

func part1(lines []string) int64 {
	numbers := util.ConvertToInts(lines)

	for _, a := range numbers {
		for _, b := range numbers {
			if a+b == 2020 {
				return a * b
			}
		}
	}

	return 0
}

func part2(lines []string) int64 {
	numbers := util.ConvertToInts(lines)

	for _, a := range numbers {
		for _, b := range numbers {
			for _, c := range numbers {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}

	return 0
}
