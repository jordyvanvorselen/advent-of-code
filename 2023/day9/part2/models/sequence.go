package models

import (
	"github.com/samber/lo"
	"slices"
)

type Sequence []int

func (s Sequence) Prediction() int {
	children := s.children()
	slices.Reverse(children)

	var prediction int

	for i, _ := range children {
		if len(children) == i+1 {
			continue
		}

		childAbove := children[i+1]

		a := prediction
		b := childAbove[0]

		prediction = b - a

		childAbove = append(childAbove, prediction)
	}

	return prediction
}

func (s Sequence) children() []Sequence {
	allSequences := []Sequence{s}

	for {
		newSeq := allSequences[len(allSequences)-1].childSequence()
		allSequences = append(allSequences, newSeq)

		if lo.Every([]int{0}, newSeq) {
			break
		}
	}

	return allSequences
}

func (s Sequence) childSequence() Sequence {
	var seq Sequence
	for i := 0; i < len(s)-1; i++ {
		a := s[i]
		b := s[i+1]

		seq = append(seq, b-a)
	}
	return seq
}
