package parser

import (
	"final-battle/model"
	"strings"
)

// ParseLine reads one input line and fills target PDA with consumed data.
func (p *Parser) ParseLine(line string, target *model.PDA) {
	lineSplit := strings.Split(line, p.Arrow)
	name, value := lineSplit[0], lineSplit[1]
	splitter := func(s string) []string { return strings.Split(strings.Trim(s, "{}"), ",") }

	switch name {
	case "InitialStackSymbol":
		target.InitialStackSymbol = value
	case "InitialState":
		target.InitialState = value
	case "States":
		states := splitter(value)
		target.States = states
	case "FinalStates":
		states := splitter(value)
		target.FinalStates = states
	case "InputAlphabet":
		ids := splitter(value)
		target.InputAlphabet = ids
	case "StackAlphabet":
		ids := splitter(value)
		target.StackAlphabet = ids
	case "Transition":
		valueSplit := strings.Split(value, p.TransitionDelimiter)
		pushValues := strings.Split(valueSplit[4], p.Delimiter)
		if valueSplit[4] == "" {
			pushValues = []string{}
		}
		target.Transitions = append(target.Transitions, model.Transition{
			From:  valueSplit[0],
			To:    valueSplit[1],
			Input: valueSplit[2],
			Pop:   valueSplit[3],
			Push:  pushValues,
		})
	default:
		panic("invalid input: " + line)
	}
}
