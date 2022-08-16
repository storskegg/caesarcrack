package main

import (
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

type Shift int

func caesar(r rune, shift Shift) (shifted rune) {
	if !unicode.IsLetter(r) {
		return r
	}

	var wasUpper bool

	if unicode.IsUpper(r) {
		wasUpper = true
		r = unicode.ToLower(r)
		if Debug() {
			spew.Dump(struct {
				R string
				U bool
			}{
				R: string(r),
				U: wasUpper,
			})
		}
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

func caesarShift(in string, shift Shift) string {
	return strings.Map(func(r rune) rune {
		return caesar(r, shift)
	}, in)
}
