package main

// IsStackIndependent returns true if the PDA is stack independent, false otherwise
func (pda *PDA) IsStackIndependent() bool {
	for _, transition := range pda.Transitions {
		if transition.Pop != string(UniversalQuantifier) {
			return false
		}
	}

	// For each state, check if the number of transitions from that state
	// for a given input symbol is equal to the number of stack symbols
	for _, state := range pda.States {
		inputSymbols := make(map[string]int)
		for _, transition := range pda.Transitions {
			if transition.From == state {
				inputSymbols[transition.Input]++
			}
		}
		for _, transition := range pda.Transitions {
			if transition.From == state && len(transition.Push) != inputSymbols[transition.Input] {
				return false
			}
		}
	}
	return true
}