package util

import (
	"bufio"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func ConfigureLogging() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	log.Logger = log.Output(consoleWriter)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func ReadInput(path string) ([]int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []int64
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
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
