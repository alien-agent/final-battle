package main

import (
	"fmt"
	"strings"
)

// ToDOT generates a DOT format representation of the PDA
func (pda *PDA) ToDOT() string {
	var graph strings.Builder

	// Add the initial state
	fmt.Fprintf(&graph, "  %s [shape=doublecircle];\n", pda.InitialState)

	// Add the final states
	for _, state := range pda.FinalStates {
		fmt.Fprintf(&graph, "  %s [shape=doublecircle];\n", state)
	}

	// Add the other states
	for _, state := range pda.States {
		if !contains(pda.FinalStates, state) && state != pda.InitialState {
			fmt.Fprintf(&graph, "  %s [shape=oval];\n", state)
		}
	}

	// Add the transitions
	for _, transition := range pda.Transitions {
		label := fmt.Sprintf("%s, %s", transition.Input, transition.Pop)
		if len(transition.Push) > 0 {
			label += ", " + strings.Join(transition.Push, "")
		}
		fmt.Fprintf(&graph, "  %s -> %s [label=\"%s\"];\n", transition.From, transition.To, label)
	}

	return fmt.Sprintf("digraph {\n%s}", graph.String())
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}