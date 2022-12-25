package main

// IsStackIndependent returns whether the transition is stack independent or not
func (t *Transition) IsStackIndependent(pda *PDA) bool {

	if t.Pop == string(UniversalQuantifier) {
		return true
	}

	numTransitions := 0
	for _, transition := range pda.Transitions {
		if transition.From == t.From && transition.To == t.To {
			numTransitions++
		}
	}
	return numTransitions == len(pda.StackAlphabet)
}
