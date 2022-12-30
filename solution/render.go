package solution

import (
	"final-battle/model"
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

// RenderDOT returns a string with DOT representation of PDA.
func RenderDOT(pda *model.PDA) string {
	var graph strings.Builder

	fmt.Fprintf(&graph, "  graph [pad=\"1\", nodesep=\"1\", ranksep=\"3\"];\n")

	// Add states
	for _, state := range pda.States {
		color := "black"
		if pda.IsTrapState(state) {
			color = "red"
		}

		shape := "circle"
		if slices.Contains(pda.FinalStates, state) {
			shape = "doublecircle"
		}

		style := "none"
		if state == pda.InitialState {
			style = "dashed"
		}
		fmt.Fprintf(&graph, "  %s [shape=%s, color=%s, style=%s];\n", state, shape, color, style)
	}

	// Add transitions
	for _, transition := range pda.Transitions {
		label := fmt.Sprintf("%s/%s/", transition.Input, transition.Pop)
		if len(transition.Push) > 0 {
			label += strings.Join(transition.Push, ",")
		}

		style := "none"
		if pda.IsTransitionDeterministic(transition) {
			style = "dashed"
		}

		color := "black"
		if pda.IsStackIndependent(transition) {
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
