package main

import (
	"adventofcode-2020/util"
	"errors"
	"github.com/rs/zerolog/log"
)

func main() {
	util.ConfigureLogging()

	numbers, err := util.ReadInput("./01/input.txt")
	util.HandleError(err)

	answer1 := calculateAnswerPart1(numbers)
	if answer1 == 0 {
		util.HandleError(errors.New("couldn't find answer"))
	}

	log.Info().Int64("answer", answer1).Msg("Part 1")

	answer2 := calculateAnswerPart2(numbers)
	if answer2 == 0 {
		util.HandleError(errors.New("couldn't find answer"))
	}

	log.Info().Int64("answer", answer2).Msg("Part 2")
}

func calculateAnswerPart1(numbers []int64) int64 {
	for _, a := range numbers {
		for _, b := range numbers {
			if a+b == 2020 {
				return a * b
			}
		}
	}

	return 0
}

func calculateAnswerPart2(numbers []int64) int64 {
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
