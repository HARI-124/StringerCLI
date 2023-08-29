package stringer

import (
	"strconv"
)

func Reverse(input string) (result string) {
	for _, a := range input {
		result = string(a) + result
	}
	return result
}

func Inspect(input string, digits bool) (count int, kind string) {
	if !digits {
		return len(input), "char"
	}
	return inspectNumbers(input), "digit"
}

func inspectNumbers(input string) (count int) {
	for _, c := range input {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			count++
		}
	}
	return count
}

func Count(input string) (result int) {
	result = len(input)
	return result
}
