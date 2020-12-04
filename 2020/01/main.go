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

	answer := calculateAnswer(numbers)
	if answer == 0 {
		util.HandleError(errors.New("couldn't find answer"))
	}

	log.Info().Int64("answer", answer).Msg("Got it!")
}

func calculateAnswer(numbers []int64) int64 {
	for _, a := range numbers {
		for _, b := range numbers {
			if a+b == 2020 {
				return a * b
			}
		}
	}

	return 0
}
