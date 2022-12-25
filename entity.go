package main

const (
	Epsilon             = 'ε'
	UniversalQuantifier = '∀'
)

type PDA struct {
	States      []string
	FinalStates []string

	InputAlphabet []string
	StackAlphabet []string

	InitialState       string
	InitialStackSymbol string

	Transitions []Transition
}

type Transition struct {
	From, To, Input string
	Pop             string
	Push            []string
}
