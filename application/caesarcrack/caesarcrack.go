package caesarcrack

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/storskegg/caesarcrack/internal/bruteforce"
	"github.com/storskegg/caesarcrack/internal/cipher"
	"github.com/storskegg/caesarcrack/internal/confidence"
	"github.com/storskegg/caesarcrack/internal/dictionary"
	"github.com/storskegg/caesarcrack/internal/words"
)

var flagFilePath string
var flagDictPath string

const (
	Version = "1.1.0"

	defaultDictPath = "./words_alpha.txt"
)

func init() {
	cmdRoot.Flags().StringVarP(&flagFilePath, "file", "f", "", "Path to the ciphered file.")
	cmdRoot.Flags().StringVarP(&flagDictPath, "dict", "d", defaultDictPath, "Path to the dictionary file.")

	if err := cmdRoot.MarkFlagRequired("file"); err != nil {
		panic(err)
	}
}

var cmdRoot = &cobra.Command{
	Use:     "caesarcrack",
	Short:   "Attempts to brute force decrypt a Caesar cipher.",
	Example: "caesarcrack -f ciphered.txt",
	Version: Version,
	RunE:    rootRunE,
}

func Execute() error {
	return cmdRoot.Execute()
}

func rootRunE(cmd *cobra.Command, args []string) error {
	log.Printf("Loading dictionary from %s...", flagDictPath)
	dict, err := dictionary.NewFromFile(flagDictPath)
	if err != nil {
		return err
	}
	log.Printf("Dictionary Loaded -- %d Words\n", dict.Length())

	fBytes, err := os.ReadFile(flagFilePath)
	if err != nil {
		return err
	}

	bf := bruteforce.New()
	bf.ParsePhrase(string(fBytes))

	var shift cipher.Shift
	for shift = 0; shift < 26; shift++ {
		//log.Printf("====[ %02d / %02d ]=====================\n", shift, origShift(shift))
		bf.InitShifted(shift)

		var t, shifted string
		found := false
		for idx, word := range bf.CipherText() {
			shifted = cipher.CaesarShift(word, shift)
			bf.Insert(shift, idx, shifted)
			t = strings.ToLower(words.StripNonLetters(shifted))
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

		//log.Printf("Confidence: %2.2f\n", bf.shiftConfidenceMap.Confidence(shift))
	}

	bestShift, bestConfidence, decipher := bf.DoWithBest()

	fmt.Printf("[%02d|%02.2f] %s\n", confidence.OrigShift(bestShift), bestConfidence, decipher)
	return nil
}
