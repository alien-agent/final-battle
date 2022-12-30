package model

import "golang.org/x/exp/slices"

type Transition struct {
	From, To, Input string
	Pop             string
	Push            []string
}

func (t Transition) Equal(other Transition) bool {
	return t.From == other.From &&
		t.To == other.To &&
		t.Input == other.Input &&
		t.Pop == other.Pop &&
		slices.Equal(t.Push, other.Push)
}
