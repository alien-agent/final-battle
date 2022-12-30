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
	// Add the initial state
	fmt.Fprintf(&graph, "  %s [shape=doublecircle];\n", pda.InitialState)

	// Add the final states
	//for _, state := range pda.FinalStates {
	//	color := "black"
	//	if pda.IsTrapState(state) {
	//		color = "red"
	//	}
	//	fmt.Fprintf(&graph, "  %s [shape=doublecircle, color=%s];\n", state, color)
	//	fmt.Fprintf(&graph, "  %s [color=%s];\n", state, color)
	//}

	// Add the other states
	for _, state := range pda.States {
		color := "black"
		if pda.IsTrapState(state) {
			color = "red"
		}

		shape := "circle"
		if slices.Contains(pda.FinalStates, state) {
			shape = "doublecircle"
		}
		if !contains(pda.FinalStates, state) && state != pda.InitialState {
			fmt.Fprintf(&graph, "  %s [shape=%s, color=%s];\n\n", state, shape ,color)
		}
	}

	// Add the transitions
	for _, transition := range pda.Transitions {
		label := fmt.Sprintf("%s/%s/", transition.Input, transition.Pop)
		if len(transition.Push) > 0 {
			label += strings.Join(transition.Push, ", ")
		}

		IsDeterministic := pda.IsTransitionDeterministic(transition)
		style := "none"
		if IsDeterministic {
			style = "dashed"
		}

		IsIndependent := pda.IsStackIndependent(transition)
		color := "black"
		if IsIndependent {
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
