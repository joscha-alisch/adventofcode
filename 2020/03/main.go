package main

import (
	"adventofcode-2020/util"
)

func main() {
	util.RunIt("./03/input.txt", part1, part2)
}

func part1(lines []string) int64 {
	return checkTrees(lines, 3, 1)
}

func part2(lines []string) int64 {
	return checkTrees(lines, 1, 1) *
		checkTrees(lines, 3, 1) *
		checkTrees(lines, 5, 1) *
		checkTrees(lines, 7, 1) *
		checkTrees(lines, 1, 2)
}

func checkTrees(hill []string, right int, down int) int64 {
	pos := 0
	n := len(hill[0])
	var trees int64

	for i := 0; i < len(hill); i += down {
		row := hill[i]
		if string(row[pos]) == "#" {
			trees++
		}

		pos = pos + right
		if pos >= n {
			pos = pos - n
		}
	}

	return trees
}
