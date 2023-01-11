package solution

import (
	"final-battle/model"
	"fmt"
	"golang.org/x/exp/slices"
	"unicode/utf8"
)

type StackState struct {
	State string
	Stack []string
	Ind   int
}

func Solve(pda *model.PDA, word string) (recognized bool, wasNondeterministic bool, stackStateList []StackState) {
	stackStateList = []StackState{}
	word = word + pda.Epsilon
	// Initialize the set of visited State-Stack pairs
	visited := map[string]bool{}

	queue := []StackState{{State: pda.InitialState, Stack: []string{pda.InitialStackSymbol}, Ind: 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		key := fmt.Sprintf("%s|%v|%d", curr.State, curr.Stack, curr.Ind)
		if visited[key] || curr.Ind > utf8.RuneCountInString(word) || len(curr.Stack) == 0 {
			continue
		}
		visited[key] = true

		stackStateList = append(stackStateList, curr)

		if curr.Ind == utf8.RuneCountInString(word)-1 {
			if slices.Contains(pda.FinalStates, curr.State) {
				recognized = true
			}
		}

		cnt := 0
		for _, t := range pda.Transitions {
			if t.From != curr.State ||
				(t.Input != pda.Epsilon &&
					(curr.Ind == utf8.RuneCountInString(word) ||
						(t.Input != pda.UniversalQuantifier &&
							t.Input != string([]rune(word)[curr.Ind])))) ||
				(t.Pop != pda.UniversalQuantifier && t.Pop != pda.Epsilon && t.Pop != curr.Stack[len(curr.Stack)-1]) {
				continue
			}
			cnt += 1
			if cnt > 1 {
				wasNondeterministic = true
			}
			nextState := t.To

			nextStackLen := len(curr.Stack)
			if t.Pop != pda.Epsilon {
				nextStackLen -= 1
			}
			nextInd := curr.Ind
			if t.Input != pda.Epsilon {
				nextInd += 1
			}

			nextStack := make([]string, nextStackLen)
			copy(nextStack, curr.Stack)
			for i := len(t.Push) - 1; i >= 0; i-- {
				if t.Pop == pda.UniversalQuantifier && t.Push[i] == pda.UniversalQuantifier {
					topSymbol := curr.Stack[len(curr.Stack)-1]
					nextStack = append(nextStack, topSymbol)
					continue
				}
				nextStack = append(nextStack, t.Push[i])
			}

			next := StackState{State: nextState, Stack: nextStack, Ind: nextInd}
			queue = append(queue, next)

		}
	}

	return
}
