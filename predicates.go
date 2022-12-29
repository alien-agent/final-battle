package main

import "golang.org/x/exp/slices"

func (p *PDA) IsStackIndependent(t Transition) bool {
	if t.Pop == string(UniversalQuantifier) {
		return true
	}

	var PopSymbols []string
	for _, transition := range p.Transitions {
		if transition.From == t.From && transition.To == t.To {
			PopSymbols = append(PopSymbols, transition.Pop)
		}
	}

	AlphabetMap := make(map[string]struct{})
	for _, element := range p.StackAlphabet {
		AlphabetMap[element] = struct{}{}
	}

	PopSymbolsMap := make(map[string]struct{})
	for _, element := range PopSymbols {
		PopSymbolsMap[element] = struct{}{}
	}

	// Check if the two sets are equal
	if len(AlphabetMap) != len(PopSymbolsMap) {
		return false
	}
	equal := true
	for key := range AlphabetMap {
		if _, ok := PopSymbolsMap[key]; !ok {
			equal = false
			break
		}
	}
	if equal {
		return true
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
