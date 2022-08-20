package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"io"
	"log"
	"os"
)

var (
	//go:embed words_alpha.txt
	wordsList []byte
)

type Dictionary map[string]struct{}

func (d Dictionary) Has(needle string) bool {
	_, ok := d[needle]
	return ok
}

func (d Dictionary) Load() error {
	log.Println("Loading internal dictionary...")
	buf := bytes.NewBuffer(wordsList)

	return d.scanToDict(buf)
}

func (d Dictionary) LoadWithDictionary(path string) error {
	log.Printf("Loading dictionary from %s...", path)
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return d.scanToDict(f)
}

func (d Dictionary) scanToDict(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	n := 0
	for scanner.Scan() {
		n++
		d[scanner.Text()] = struct{}{}
	}
	log.Printf("Dictionary Loaded -- %d Words\n", n)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}
