package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func (p *PDA) IsStackIndependent(t Transition) bool {
	if t.Pop == string(UniversalQuantifier) {
		return true
	}

	var array2 []string
	for _, transition := range p.Transitions {
		if transition.From == t.From && transition.To == t.To {
			array2 = append(array2, transition.Pop)
		}
	}


	// Create a map to store the elements of array1
	set1 := make(map[string]struct{})
	for _, element := range p.StackAlphabet {
		set1[element] = struct{}{}
	}

	// Create a map to store the elements of array2
	set2 := make(map[string]struct{})
	for _, element := range array2 {
		set2[element] = struct{}{}
	}

	// Check if the two sets are equal
	if len(set1) == len(set2) {
		equal := true
		for key := range set1 {
			if _, ok := set2[key]; !ok {
				equal = false
				break
			}
		}
		if equal {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (p *PDA) IsTransitionDeterministic(t Transition) bool {
	return !(t.Input == Epsilon) &&
		!slices.ContainsFunc(p.Transitions, func(tt Transition) bool { return tt.From == t.From && tt.Input == t.Input })
}

func (p *PDA) IsTrapState(state string) bool {
	return !slices.Contains(p.FinalStates, state) &&
		!slices.ContainsFunc(p.Transitions, func(t Transition) bool { return t.From == state })
}
