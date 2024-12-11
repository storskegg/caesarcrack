package cipher

import (
	"strings"
	"unicode"
)

type Shift int

func Caesar(r rune, shift Shift) (shifted rune) {
	if !unicode.IsLetter(r) {
		return r
	}

	var wasUpper bool

	if unicode.IsUpper(r) {
		wasUpper = true
		r = unicode.ToLower(r)
	}

	s := Shift(r) + shift
	if s > 'z' {
		shifted = rune(s - 26)
	} else if s < 'a' {
		shifted = rune(s + 26)
	} else {
		shifted = rune(s)
	}

	if wasUpper {
		shifted = unicode.ToUpper(shifted)
	}

	return
}

func CaesarShift(in string, shift Shift) string {
	return strings.Map(func(r rune) rune {
		return Caesar(r, shift)
	}, in)
}
