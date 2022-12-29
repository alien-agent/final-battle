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

	numTransitions := 0
	for _, transition := range p.Transitions {
		if transition.From == t.From && transition.To == t.To {
			numTransitions++
		}
	}
	return numTransitions == len(p.StackAlphabet)
}

func (p PDA) IsTransitionDeterministic(t Transition) bool {
	return !(t.Input == p.Epsilon) &&
		!slices.ContainsFunc(p.Transitions, func(tt Transition) bool { return tt.From == t.From && tt.Input == t.Input })
}

func (p PDA) IsTrapState(state string) bool {
	return !slices.Contains(p.FinalStates, state) &&
		!slices.ContainsFunc(p.Transitions, func(t Transition) bool { return t.From == state })
}
