package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	originalString := []rune(s)
	newString := strings.Builder{}
	var prev, next rune
	escapingOn := false
	for i, r := range originalString {
		fmt.Printf("%q", r)

		if i > 0 {
			prev = originalString[i-1]
		} else {
			prev = 0
		}

		if i < len(s)-1 {
			next = originalString[i+1]
		} else {
			next = 0
		}

		if r == '\\' && !escapingOn {
			if next == 0 {
				return "", ErrInvalidString
			}
			escapingOn = true
			continue
		}

		if escapingOn && unicode.IsLetter(r) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(r) && !escapingOn && (prev == 0 || unicode.IsDigit(next)) {
			return "", ErrInvalidString
		}

		if unicode.IsLetter(r) || unicode.IsSpace(r) || escapingOn {
			quantity, err := strconv.Atoi(string(next))
			if err != nil {
				quantity = 1
			}
			newString.WriteString(strings.Repeat(string(r), quantity))
			escapingOn = false
		}
	}
	return newString.String(), nil
}
