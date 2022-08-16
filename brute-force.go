package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type BruteForce struct {
	SampleText         Words
	CipherText         Words
	ShiftConfidenceMap *ShiftConfidence
	Shifted            map[Shift][]string
}

func (bf *BruteForce) Len() int {
	return len(bf.CipherText)
}

func (bf *BruteForce) ParsePhrase(s string) {
	bf.CipherText = parsePhrase(s)
	n := len(bf.CipherText)
	if n > 1000 {
		n = 1000
	}
	bf.SampleText = bf.CipherText[:n]
	bf.ShiftConfidenceMap = NewShiftConfidence(len(bf.SampleText))
	bf.Shifted = make(map[Shift][]string)
}

func (bf *BruteForce) InitShifted(shift Shift) {
	if _, ok := bf.Shifted[shift]; !ok {
		bf.Shifted[shift] = make([]string, bf.Len())
	}
}

func (bf *BruteForce) Insert(shift Shift, idx int, shifted string) {
	bf.Shifted[shift][idx] = shifted
}

func (bf *BruteForce) Join(shift Shift, sep string) string {
	return strings.Join(bf.Shifted[shift], " ")
}

func (bf *BruteForce) DoWithBest() (Shift, float64, string) {
	if Debug() {
		fmt.Println("Sh\tC")
		spew.Dump(bf.ShiftConfidenceMap.Map())
	}

	bestShift, bestConfidence := bf.ShiftConfidenceMap.Best()
	log.Printf("Chosen shift: %02d, confidence: %02.2f", origShift(bestShift), bestConfidence)

	return bestShift, bestConfidence, bf.Join(bestShift, " ")
}

func (bf *BruteForce) Confidence(shift Shift) float64 {
	return bf.ShiftConfidenceMap.Confidence(shift)
}

func (bf *BruteForce) Increment(shift Shift) {
	bf.ShiftConfidenceMap.Increment(shift)
}

func (bf *BruteForce) ResetConfidence(shift Shift) {
	bf.ShiftConfidenceMap.Zero(shift)
}
