package main

type StackState struct {
	state string
	stack []string
}

func (pda *PDA) solve(word string) (recognized bool, wasNondeterministic bool, stackStateList []StackState) {
	stack := []string{pda.InitialStackSymbol}

	currentState := pda.InitialState

	stackStateList = []StackState{{currentState, stack}}

	for _, input := range word {
		var nextTransition Transition
		foundTransition := false
		for _, t := range pda.Transitions {
			if t.From == currentState && t.Input == string(input) && t.Pop == stack[len(stack)-1] {
				nextTransition = t
				foundTransition = true
				break
			}
		}

		if !foundTransition {
			return false, false, stackStateList
		}

		stack = stack[:len(stack)-1]

		for i := len(nextTransition.Push) - 1; i >= 0; i-- {
			stack = append(stack, nextTransition.Push[i])
		}

		currentState = nextTransition.To

		stackStateList = append(stackStateList, StackState{state: currentState, stack: stack})
	}

	return contains(pda.FinalStates, currentState), false, stackStateList
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
