package util

import (
	"bufio"
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func RunIt(input string, part1 func(lines []string) int64, part2 func(lines []string) int64) {
	ConfigureLogging()

	in, err := ReadInput(input)
	HandleError(err)

	answer1 := part1(in)
	if answer1 == 0 {
		HandleError(errors.New("couldn't find answer for part 1"))
	}

	log.Info().Int64("answer", answer1).Msg("Part 1")

	answer2 := part2(in)
	if answer2 == 0 {
		HandleError(errors.New("couldn't find answer for part 2"))
	}

	log.Info().Int64("answer", answer2).Msg("Part 2")
}

func ConfigureLogging() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	log.Logger = log.Output(consoleWriter)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func ReadInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func HandleError(err error) {
	if err != nil {
		log.Error().Err(err).Msg("")
		os.Exit(1)
	}
}

func ConvertToInts(arr []string) []int64 {
	var result []int64
	for _, s := range arr {
		i, err := strconv.ParseInt(s, 10, 64)
		HandleError(err)
		result = append(result, i)
	}
	return result
}

func GroupByBlankLines(lines []string) [][]string {
	var results [][]string
	var current []string

	for _, line := range lines {
		if line == "" {
			results = append(results, current)
			current = make([]string, 0)
		} else {
			current = append(current, line)
		}
	}
	results = append(results, current)

	return results
}

func CountUniqueChars(s string) int {
	count := 0
	m := NewHashset()
	for _, c := range s {
		if !m.Contains(int(c)) {
			m.Add(int(c))
			count++
		}
	}
	return count
}
