package model

type Transition struct {
	From, To, Input string
	Pop             string
	Push            []string
}
