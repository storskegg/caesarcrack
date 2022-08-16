package main

import (
	"math"
)

type ShiftConfidence struct {
	Max float64
	CM  map[Shift]int
}

func NewShiftConfidence(maxConfidence int) *ShiftConfidence {
	sw := ShiftConfidence{}
	sw.CM = make(map[Shift]int)
	sw.Max = float64(maxConfidence)
	return &sw
}

func (s *ShiftConfidence) Zero(shift Shift) {
	s.CM[shift] = 0
}

func (s *ShiftConfidence) Increment(shift Shift) {
	if c, ok := s.CM[shift]; ok {
		s.CM[shift] = c + 1
		return
	}
	s.CM[shift] = 1
}

func (s *ShiftConfidence) Map() map[Shift]int {
	return s.CM
}

func (s *ShiftConfidence) Best() (bestShift Shift, bestConfidence float64) {
	var best int
	for shift, confidence := range s.CM {
		if confidence > best {
			bestShift = shift
			best = confidence
		}
	}

	return bestShift, float64(best) / s.Max
}

func (s *ShiftConfidence) Confidence(shift Shift) float64 {
	cs, ok := s.CM[shift]
	if !ok {
		return 0
	}

	return float64(cs) / s.Max
}

func origShift(shift Shift) (orig int) {
	orig = int(math.Abs(float64(shift - 26)))
	if orig == 26 {
		orig = 0
	}
	return
}
