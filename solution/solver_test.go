package solution

import (
	"final-battle/model"
	"testing"
)

const testEpsilon = "ε"

func TestBracketSequencePDA(t *testing.T) {
	pda := &model.PDA{
		Epsilon: testEpsilon,

		States:      []string{"q0", "q1"},
		FinalStates: []string{"q0"},

		InputAlphabet: []string{"(", ")"},
		StackAlphabet: []string{"Z0", "B"},

		InitialState:       "q0",
		InitialStackSymbol: "Z0",

		Transitions: []model.Transition{
			{"q0", "q1", "(", "Z0", []string{"B", "Z0"}},
			{"q1", "q1", "(", "B", []string{"B", "B"}},
			{"q1", "q1", ")", "B", []string{}},
			{"q1", "q0", testEpsilon, "Z0", []string{"Z0"}},
		},
	}

	tests := []struct {
		word     string
		expected bool
	}{
		{"()", true},
		{"((()))", true},
		{"(()())", true},
		{"()()()", true},
		{"()(()())((()))(()())()(()())((()))", true},
		{"()(()())((()))(()())()())())((()))", false},
		{"())(", false},
		{"(()", false},
		{"())", false},
	}

	for _, test := range tests {
		recognized, _, _ := Solve(pda, test.word)
		if recognized != test.expected {
			t.Errorf("unexpected result for word %q: got %v, want %v", test.word, recognized, test.expected)
		}
	}
}
