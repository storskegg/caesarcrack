package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/urfave/cli/v2"
)

func Debug() bool {
	return false
}

func parsePhrase(s string) Words {
	split := strings.Split(s, " ")
	unsorted := make(Words, len(split))
	for i, w := range split {
		unsorted[i] = &Word{
			Word: w,
		}
	}

	return unsorted
}

func main() {
	app := &cli.App{
		Name:   "caesarcrack",
		Usage:  "performs dictionary brute force attack on caesar cipher",
		Action: doit,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "ciphered file",
				Required: true,
			},
			&cli.BoolFlag{
				Name:  "color",
				Usage: "colorful results",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func doit(c *cli.Context) error {
	dict := make(Dictionary)
	err := dict.Load()
	//err := dict.LoadWithDictionary("./words_alpha.txt")
	if err != nil {
		return err
	}

	fPath := c.String("file")
	fBytes, err := os.ReadFile(fPath)
	if err != nil {
		return err
	}

	bf := BruteForce{}
	bf.ParsePhrase(string(fBytes))

	var shift Shift
	for shift = 0; shift < 26; shift++ {
		//log.Printf("====[ %02d / %02d ]=====================\n", shift, origShift(shift))
		bf.InitShifted(shift)

		var t, shifted string
		found := false
		for idx, word := range bf.CipherText {
			shifted = caesarShift(word, shift)
			bf.Insert(shift, idx, shifted)
			t = strings.ToLower(stripNonLetters(shifted))
			found = dict.Has(t)

			if found {
				bf.Increment(shift)
				//log.Printf("'%s' found\n", shifted)
			} else {
				if bf.Confidence(shift) == 0 {
					bf.ResetConfidence(shift)
				}
				//log.Printf("'%s' not found\n", shifted)
			}
		}

		//log.Printf("Confidence: %2.2f\n", bf.ShiftConfidenceMap.Confidence(shift))
	}

	bestShift, bestConfidence, decipher := bf.DoWithBest()

	fmt.Printf("[%02d|%02.2f] %s\n", origShift(bestShift), bestConfidence, decipher)
	return nil
}

func stripNonLetters(word string) string {
	return strings.Trim(strings.Map(func(r rune) (o rune) {
		if unicode.IsLetter(r) {
			o = r
		}
		return
	}, word), "\x00")
}
