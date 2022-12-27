package main

import (
	"strings"
)

var lineCleaner = strings.NewReplacer(
	" ", "",
	"\r", "",
	"\t", "",
)

type Parser struct {
	Arrow               string
	Delimiter           string
	TransitionDelimiter string
	UniversalQuantifier string
	Epsilon             string
}

func (p *Parser) Parse(input string) *PDA {
	var result PDA
	for _, v := range strings.Split(input, "\n") {
		line := lineCleaner.Replace(v)
		if line == "" {
			continue
		}

		p.parseLine(line, &result)
	}

	return &result
}

func (p *Parser) parseLine(line string, target *PDA) {
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
		target.Transitions = append(target.Transitions, Transition{
			From:  valueSplit[0],
			To:    valueSplit[1],
			Input: valueSplit[2],
			Pop:   valueSplit[3],
			Push:  strings.Split(valueSplit[4], p.Delimiter),
		})
	default:
		panic("invalid input: " + line)
	}
}
