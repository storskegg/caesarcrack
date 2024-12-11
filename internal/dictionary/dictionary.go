package dictionary

import (
	"github.com/storskegg/autocorrect/wordcount"
)

type Dictionary interface {
	Add(word string)
	Has(needle string) bool
	Length() int
}

func NewFromFile(path string) (Dictionary, error) {
	return wordcount.NewWordCountFromDictionary(path)
}
