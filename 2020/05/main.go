package main

import (
	"adventofcode-2020/util"
	"math"
)

func main() {
	util.RunIt("./05/input.txt", part1, part2)
}

func part1(lines []string) int64 {
	var highestSeatId float64 = 0
	for _, line := range lines {
		row := binarySpace(line[:7], 'B', 'F')
		col := binarySpace(line[7:10], 'R', 'L')

		id := row*8 + col

		highestSeatId = math.Max(float64(id), highestSeatId)
	}

	return int64(highestSeatId)
}

func part2(lines []string) int64 {
	ids := NewHashset()
	maxId := 127*8 + 7

	for _, line := range lines {
		row := binarySpace(line[:7], 'B', 'F')
		col := binarySpace(line[7:10], 'R', 'L')

		id := row*8 + col

		ids.Add(id)
	}

	for i := 0; i < maxId; i++ {
		if !ids.Contains(i) && ids.Contains(i-1) && ids.Contains(i+1) {
			return int64(i)
		}
	}

	return 0
}

func binarySpace(line string, upper rune, lower rune) int {
	min := 0
	max := int(math.Pow(2, float64(len(line))))

	for _, i := range line {
		if i == upper {
			min += (max - min) / 2
		} else if i == lower {
			max -= (max - min) / 2
		}
	}

	return min
}
