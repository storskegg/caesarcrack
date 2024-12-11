package words

import (
	"strings"
	"unicode"
)

type Words []string

func StripNonLetters(word string) string {
	return strings.Trim(strings.Map(func(r rune) (o rune) {
		if unicode.IsLetter(r) {
			o = r
		}
		return
	}, word), "\x00")
}
