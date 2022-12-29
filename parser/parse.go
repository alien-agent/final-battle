package parser

import (
	"final-battle/model"
	"strings"
)

var lineCleaner = strings.NewReplacer(
	" ", "",
	"\r", "",
	"\t", "",
)

// Parse reads all lines from input and produces a corresponding PDA.
func (p *Parser) Parse(input string) *model.PDA {
	var result model.PDA
	for _, v := range strings.Split(input, "\n") {
		line := lineCleaner.Replace(v)
		if line == "" {
			continue
		}

		p.ParseLine(line, &result)
	}

	result.Epsilon = p.Epsilon
	result.UniversalQuantifier = p.UniversalQuantifier
	return &result
}
