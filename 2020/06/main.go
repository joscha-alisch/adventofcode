package main

import (
	"adventofcode-2020/util"
)

func main() {
	util.RunIt("./06/input.txt", part1, part2)
}

func part1(lines []string) int64 {
	result := 0
	groups := util.GroupByBlankLines(lines)

	for _, group := range groups {
		bitsets := toBitSets(group)
		result += bitsets[0].Or(bitsets[1:]...).Count()
	}

	return int64(result)
}

func part2(lines []string) int64 {
	result := 0
	groups := util.GroupByBlankLines(lines)

	for _, group := range groups {
		bitsets := toBitSets(group)
		result += bitsets[0].And(bitsets[1:]...).Count()
	}

	return int64(result)
}

func toBitSets(lines []string) []util.BitSet {
	var bitsets []util.BitSet
	for _, line := range lines {
		bitsets = append(bitsets, toBitset(line))
	}
	return bitsets
}

func toBitset(line string) util.BitSet {
	b := util.NewBitSet(26)
	for _, c := range line {
		b.Set(int(c-'a'), true)
	}
	return b
}
