package dictionary

import (
	"bytes"
	_ "embed"

	"github.com/storskegg/autocorrect/wordcount"
)

//go:embed words_alpha.txt
var internalDict []byte

type Dictionary interface {
	Add(word string)
	Has(needle string) bool
	Length() int
}

func NewFromFile(path string) (Dictionary, error) {
	return wordcount.NewWordCountFromDictionary(path)
}

func NewInternal() (Dictionary, error) {
	buf := bytes.NewBuffer(internalDict)

	return wordcount.NewFromReader(buf)
}
