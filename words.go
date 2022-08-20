package main

import "github.com/fatih/color"

type Words []*Word

func (w Words) Slice() []string {
	s := make([]string, len(w))
	for i, word := range w {
		s[i] = word.String()
	}
	return s
}

type Word struct {
	Word  string
	found bool
}

func (w *Word) String() string {
	if w.found {
		return w.Word
	}
	return color.RedString(w.Word)
}
