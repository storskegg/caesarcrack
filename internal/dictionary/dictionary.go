package dictionary

import (
	"bufio"
	"log"
	"os"
)

type Dictionary map[string]struct{}

func (d Dictionary) Has(needle string) bool {
	_, ok := d[needle]
	return ok
}

func (d Dictionary) LoadWithDictionary(path string) error {
	log.Printf("Loading dictionary from %s...", path)
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	n := 0
	for scanner.Scan() {
		if word := scanner.Text(); word == "" {
			continue
		} else {
			n++
			d[word] = struct{}{}
		}
	}
	log.Printf("Dictionary Loaded -- %d Words\n", n)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
