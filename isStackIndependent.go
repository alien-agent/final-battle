package main

func (t *Transition) IsStackIndependent(pda *PDA) bool {

	for _, transition := range pda.Transitions {
		if transition.Pop == string(UniversalQuantifier) {
			return true
		}
	}

	numTransitions := 0
	for _, transition := range pda.Transitions {
		if transition.From == t.From && transition.To == t.To {
			numTransitions++
		}
	}
	return numTransitions == len(pda.StackAlphabet)
}
