package bruteforce

import (
	"log"
	"strings"

	"github.com/storskegg/caesarcrack/internal/cipher"
	"github.com/storskegg/caesarcrack/internal/confidence"
	"github.com/storskegg/caesarcrack/internal/words"
)

func parsePhrase(s string) (unsorted words.Words) {
	unsorted = strings.Split(s, " ")

	return
}

type BruteForcer interface {
	Len() int
	ParsePhrase(s string)
	InitShifted(shift cipher.Shift)
	Insert(shift cipher.Shift, idx int, shifted string)
	Join(shift cipher.Shift, sep string) string
	DoWithBest() (cipher.Shift, float64, string)
	Confidence(shift cipher.Shift) float64
	Increment(shift cipher.Shift)
	ResetConfidence(shift cipher.Shift)

	SampleText() words.Words
	CipherText() words.Words
	ShiftedConfidenceMap() *confidence.ShiftConfidence
	Shifted() map[cipher.Shift][]string
}

func New() BruteForcer {
	return &bruteForce{}
}

type bruteForce struct {
	sampleText         words.Words
	cipherText         words.Words
	shiftConfidenceMap *confidence.ShiftConfidence
	shifted            map[cipher.Shift][]string
}

func (bf *bruteForce) Len() int {
	return len(bf.cipherText)
}

func (bf *bruteForce) ParsePhrase(s string) {
	bf.cipherText = parsePhrase(s)
	n := len(bf.cipherText)
	if n > 1000 {
		n = 1000
	}
	bf.sampleText = bf.cipherText[:n]
	bf.shiftConfidenceMap = confidence.NewShiftConfidence(len(bf.sampleText))
	bf.shifted = make(map[cipher.Shift][]string)
}

func (bf *bruteForce) InitShifted(shift cipher.Shift) {
	if _, ok := bf.shifted[shift]; !ok {
		bf.shifted[shift] = make([]string, bf.Len())
	}
}

func (bf *bruteForce) Insert(shift cipher.Shift, idx int, shifted string) {
	bf.shifted[shift][idx] = shifted
}

func (bf *bruteForce) Join(shift cipher.Shift, sep string) string {
	return strings.Join(bf.shifted[shift], " ")
}

func (bf *bruteForce) DoWithBest() (cipher.Shift, float64, string) {
	bestShift, bestConfidence := bf.shiftConfidenceMap.Best()
	log.Printf("Chosen shift: %02d, confidence: %02.2f", confidence.OrigShift(bestShift), bestConfidence)

	return bestShift, bestConfidence, bf.Join(bestShift, " ")
}

func (bf *bruteForce) Confidence(shift cipher.Shift) float64 {
	return bf.shiftConfidenceMap.Confidence(shift)
}

func (bf *bruteForce) Increment(shift cipher.Shift) {
	bf.shiftConfidenceMap.Increment(shift)
}

func (bf *bruteForce) ResetConfidence(shift cipher.Shift) {
	bf.shiftConfidenceMap.Zero(shift)
}

func (bf *bruteForce) SampleText() words.Words {
	return bf.sampleText
}

func (bf *bruteForce) CipherText() words.Words {
	return bf.cipherText
}

func (bf *bruteForce) ShiftedConfidenceMap() *confidence.ShiftConfidence {
	return bf.shiftConfidenceMap
}

func (bf *bruteForce) Shifted() map[cipher.Shift][]string {
	return bf.shifted
}
