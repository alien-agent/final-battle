package main

import "fmt"

type StackState struct {
	state string
	stack []string
	ind   int
}

func (pda *PDA) solve(word string) (recognized bool, wasNondeterministic bool, stackStateList []StackState) {
	stackStateList = []StackState{}

	// Initialize the set of visited state-stack pairs
	visited := map[string]bool{}

	queue := []StackState{{state: pda.InitialState, stack: []string{pda.InitialStackSymbol}, ind: 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		key := fmt.Sprintf("%s|%v|%d", curr.state, curr.stack, curr.ind)
		if visited[key] || curr.ind > len(word) || len(curr.stack) == 0 {
			continue
		}
		visited[key] = true

		stackStateList = append(stackStateList, curr)

		if curr.ind == len(word) {
			if contains(pda.FinalStates, curr.state) {
				recognized = true
			}
			continue
		}

		cnt := 0
		for _, t := range pda.Transitions {
			if t.From != curr.state || t.Input != string(word[curr.ind]) || (t.Pop != UniversalQuantifier && t.Pop != curr.stack[len(curr.stack)-1]) {
				continue
			}
			cnt += 1
			if cnt > 1 {
				wasNondeterministic = true
			}
			nextState := t.To
			nextStack := make([]string, len(curr.stack)-1)
			copy(nextStack, curr.stack)
			for i := len(t.Push) - 1; i >= 0; i-- {
				nextStack = append(nextStack, t.Push[i])
			}

			next := StackState{state: nextState, stack: nextStack, ind: curr.ind + 1}
			queue = append(queue, next)

		}
	}

	return
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
