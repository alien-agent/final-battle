package solution

import (
	"final-battle/model"
	"fmt"
	"strings"
)

// RenderDOT returns a string with DOT representation of PDA.
func RenderDOT(pda *model.PDA) string {
	var graph strings.Builder

	// Add the initial State
	fmt.Fprintf(&graph, "  %s [shape=doublecircle];\n", pda.InitialState)

	// Add the final states
	fmt.Fprintf(&graph, "[shape=doublecircle];\n")
	for _, state := range pda.FinalStates {
		color := "black"
		if pda.IsTrapState(state) {
			color = "red"
		}
		fmt.Fprintf(&graph, "  %s [color=%s];\n", state, color)
	}

	// Add the other states
	fmt.Fprintf(&graph, "[shape=circle];\n")
	for _, state := range pda.States {
		color := "black"
		if pda.IsTrapState(state) {
			color = "red"
		}
		if !contains(pda.FinalStates, state) && state != pda.InitialState {
			fmt.Fprintf(&graph, "  %s [color=%s];\n", state, color)
		}
	}

	// Add the transitions
	for _, transition := range pda.Transitions {
		label := fmt.Sprintf("%s, %s", transition.Input, transition.Pop)
		if len(transition.Push) > 0 {
			label += ", " + strings.Join(transition.Push, "")
		}

		is_determenistic := pda.IsTransitionDeterministic(transition)
		style := "none"
		if is_determenistic {
			style = "dashed"
		}

		is_independent := pda.IsStackIndependent(transition)
		color := "black"
		if is_independent {
			color = "green"
		}


		fmt.Fprintf(&graph, "  %s -> %s [label=\"%s\", style=%s, color=%s];\n",
			transition.From, transition.To, label, style, color)
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
