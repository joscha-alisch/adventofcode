package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type ValidationFunc func(string) bool
type Validator struct {
	Key      string
	Validate ValidationFunc
}

func CountValid(passports []Passport, validators []Validator) (valid int64) {
	for _, passport := range passports {
		if IsValid(passport, validators) {
			valid++
		}
	}
	return
}

func IsValid(passport Passport, validators []Validator) bool {
	for _, validator := range validators {
		if !validator.Validate(passport[validator.Key]) {
			return false
		}
	}

	return true
}

func exists(v string) bool {
	return v != ""
}

func inRange(min int, max int) ValidationFunc {
	return func(s string) bool {
		var n int
		_, err := fmt.Sscanf(s, "%d", &n)
		return err == nil && n >= min && n <= max
	}
}

func anyOf(v ...ValidationFunc) ValidationFunc {
	return func(s string) bool {
		for _, validationFunc := range v {
			if validationFunc(s) {
				return true
			}
		}
		return false
	}
}

func inRangeWithSuffix(suffix string, min int, max int) ValidationFunc {
	rangeValid := inRange(min, max)
	return func(s string) bool {
		if !strings.HasSuffix(s, suffix) {
			return false
		}

		n := strings.TrimSuffix(s, suffix)
		return rangeValid(n)
	}
}

func colorCode(s string) bool {
	if len(s) != 7 || !strings.HasPrefix(s, "#") {
		return false
	}

	_, err := hex.DecodeString(strings.TrimPrefix(s, "#"))
	return err == nil
}

func eq(s string) ValidationFunc {
	return func(v string) bool {
		return v == s
	}
}

func allOf(v ...ValidationFunc) ValidationFunc {
	return func(s string) bool {
		for _, validationFunc := range v {
			if !validationFunc(s) {
				return false
			}
		}
		return true
	}
}

func number(s string) bool {
	for _, i := range s {
		if i < '0' || i > '9' {
			return false
		}
	}
	return true
}

func length(n int) ValidationFunc {
	return func(s string) bool {
		return len(s) == n
	}
}
