package parser

import (
	"fmt"
	"os"
	"strings"
)

type Parser struct {
	Arrow               string
	Delimiter           string
	TransitionDelimiter string
	UniversalQuantifier string
	Epsilon             string
}

func NewFromLanguageFile(filename string) (*Parser, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read language file: %w", err)
	}

	var result = NewDefault()
	for i, line := range strings.Split(string(contents), "\n") {
		split := strings.Split(line, " ")
		if len(split) != 2 {
			return nil, fmt.Errorf("line [%d] is malformed", i)
		}

		switch split[0] {
		case "Arrow":
			result.Arrow = split[1]
		case "Delimiter":
			result.Delimiter = split[1]
		case "TransitionDelimiter":
			result.TransitionDelimiter = split[1]
		case "Epsilon":
			result.Epsilon = split[1]
		case "UniversalQuantifier":
			result.UniversalQuantifier = split[1]
		}
	}

	return result, nil
}

func NewDefault() *Parser {
	return &Parser{
		Arrow:               "->",
		Delimiter:           ",",
		TransitionDelimiter: "/",
		UniversalQuantifier: "∀",
		Epsilon:             "ε",
	}
}
