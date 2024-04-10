package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	originalString := []rune(s)
	newString := strings.Builder{}

	intervalStart := 0
	for i, r := range originalString {

		if unicode.IsDigit(r) {
			if i-1 < 0 {
				return "", ErrInvalidString
			}

			repeatingSymbol := originalString[i-1]

			if unicode.IsDigit(repeatingSymbol) {
				return "", ErrInvalidString
			}

			newString.WriteString(string(originalString[intervalStart : i-1]))
			intervalStart = i + 1

			quantity, _ := strconv.Atoi(string(r))
			newString.WriteString(strings.Repeat(string(repeatingSymbol), quantity))

		} else if i+1 >= len(originalString) {
			newString.WriteString(string(originalString[intervalStart:]))
		}

	}
	return newString.String(), nil
}
