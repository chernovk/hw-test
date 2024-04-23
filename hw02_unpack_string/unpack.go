package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"github.com/rivo/uniseg"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	newString := strings.Builder{}
	bufferString := ""
	escapingOn := false
	previousGrapheme := ""
	previousGraphemeLen := 0
	multiplicationJustHappened := false

	gr := uniseg.NewGraphemes(s)
	for gr.Next() {
		grapheme := gr.Str()

		if grapheme == "\\" && !escapingOn {
			escapingOn = true
			continue
		}

		if escapingOn {
			if grapheme != "\\" && !unicode.IsDigit([]rune(grapheme)[0]) {
				return "", ErrInvalidString
			}
			bufferString += grapheme
			previousGrapheme = grapheme
			previousGraphemeLen = len(grapheme)
			escapingOn = false
			continue
		}

		if unicode.IsDigit([]rune(grapheme)[0]) {
			if previousGrapheme == "" || multiplicationJustHappened {
				return "", ErrInvalidString
			}
			if grapheme == "0" {
				bufferString = bufferString[:len(bufferString)-previousGraphemeLen]
			} else {
				quantity, _ := strconv.Atoi(grapheme)
				bufferString += strings.Repeat(previousGrapheme, quantity-1)
			}
			newString.WriteString(bufferString)
			multiplicationJustHappened = true
			bufferString = ""
			continue
		}

		bufferString += grapheme
		previousGrapheme = grapheme
		previousGraphemeLen = len(grapheme)
		multiplicationJustHappened = false
	}
	newString.WriteString(bufferString)
	if escapingOn {
		return "", ErrInvalidString
	}
	return newString.String(), nil
}
