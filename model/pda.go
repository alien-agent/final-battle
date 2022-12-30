package model

import "golang.org/x/exp/slices"

type PDA struct {
	Epsilon, UniversalQuantifier string

	States      []string
	FinalStates []string

	InputAlphabet []string
	StackAlphabet []string

	InitialState       string
	InitialStackSymbol string

	Transitions []Transition
}

func (p PDA) IsStackIndependent(t Transition) bool {
	if t.Pop == p.UniversalQuantifier {
		return true
	}

	var PopSymbols []string
	for _, transition := range p.Transitions {
		if transition.From == t.From && transition.To == t.To {
			if transition.Pop == p.UniversalQuantifier {
				return true
			}
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
	for key := range AlphabetMap {
		if _, ok := PopSymbolsMap[key]; !ok {
			return false
		}
	}
	return true
}

func (p PDA) IsTransitionDeterministic(t Transition) bool {
	// If we can find another transition with same From, Input, and Pop (accounting for UniversalQuantifier),
	// then original transition is considered non-deterministic.
	for _, tt := range p.Transitions {
		if tt.From != t.From {
			continue
		}

		if (tt.Input == t.Input || tt.Input == p.Epsilon || t.Input == p.Epsilon) &&
			(tt.Pop == t.Pop || tt.Pop == p.UniversalQuantifier || t.Pop == p.UniversalQuantifier) &&
			!t.Equal(tt) {
			return false
		}
	}

	return true
}

func (p PDA) IsTrapState(state string) bool {
	return !slices.Contains(p.FinalStates, state) &&
		!slices.ContainsFunc(p.Transitions, func(t Transition) bool { return t.From == state && t.To != state })
}
