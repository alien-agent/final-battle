package main

import (
	"fmt"
	"regexp"
	"strings"
)

var cleanerRegex = regexp.MustCompile(`[\s{}]+`)

func parse(input string) *PDA {
	var (
		lines  = strings.Split(input, "\n")
		result PDA
	)

	for _, line := range lines {
		cleanLine := cleanerRegex.ReplaceAllString(line, "")
		if cleanLine == "" {
			continue
		}
		parseLine(cleanLine, &result)
	}

	fmt.Println(result)
	return &result
}

func parseLine(line string, pda *PDA) {
	lineSplit := strings.Split(line, "->")
	name, value := lineSplit[0], lineSplit[1]

	switch name {
	case "InitialStackSymbol":
		pda.InitialStackSymbol = value
	case "InitialState":
		pda.InitialState = value
	case "States":
		states := strings.Split(value, ",")
		pda.States = states
	case "FinalStates":
		states := strings.Split(value, ",")
		pda.FinalStates = states
	case "InputAlphabet":
		ids := strings.Split(value, ",")
		pda.InputAlphabet = ids
	case "StackAlphabet":
		ids := strings.Split(value, ",")
		pda.StackAlphabet = ids
	case "Transition":
		valueSplit := strings.Split(value, "/")
		pda.Transitions = append(pda.Transitions, Transition{
			From:  valueSplit[0],
			To:    valueSplit[1],
			Input: valueSplit[2],
			Pop:   valueSplit[3],
			Push:  strings.Split(valueSplit[4], ","),
		})
	default:
		panic("invalid input: " + line)
	}
}
