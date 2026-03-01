package unpacker

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string format")

func Unpack(s string) (string, error) {
	r := []rune(s)

	if len(r) == 0 {
		return "", nil
	}
	
	
	if unicode.IsDigit(r[0]) {
		return "", ErrInvalidString
	}
	
	var myStr strings.Builder
	var prevRune rune

	for _, myRune := range r {

		if unicode.IsDigit(myRune) {
			num, _ := strconv.Atoi(string(myRune))

			for i := 0; i < num - 1; i++ {
				myStr.WriteRune(prevRune)
			}

			continue // Переход к следующей руне
		}

		myStr.WriteRune(myRune)
		prevRune = myRune
	}

	return myStr.String(), nil
}