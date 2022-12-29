package main

import (
	"bufio"
	"final-battle/model"
	"final-battle/parser"
	"final-battle/solution"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"strings"
)

const (
	ModeSTDIN = "stdin"
	ModeFile  = "file"
)

type Options struct {
	Mode             string `short:"m" long:"mode" choice:"stdin" choice:"file" required:"true"`
	LanguageFilename string `long:"language-filename" description:"path to syntax-defining file"`
	PDAFilename      string `long:"pda-filename" description:"path to file with PDA definition" required:"true"`
	WordsFilename    string `long:"words-filename" description:"path to file with words to solve"`
}

func main() {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Println("Failed to parse arguments:", err)
		return
	}

	prs := parser.NewDefault()
	if opts.LanguageFilename != "" {
		prs, err = parser.NewFromLanguageFile(opts.LanguageFilename)
		if err != nil {
			fmt.Println("Unable to initialize language:", err)
			os.Exit(1)
		}
	}

	pdaContent, err := os.ReadFile(opts.PDAFilename)
	if err != nil {
		fmt.Println("Failed to read PDA: ", err)
		os.Exit(1)
	}
	pda := prs.Parse(string(pdaContent))

	if opts.Mode == ModeSTDIN {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			solveAndReport(pda, scanner.Text())
		}
	} else {
		contents, err := os.ReadFile(opts.WordsFilename)
		if err != nil {
			fmt.Println("Failed to read words from file:", err)
			os.Exit(1)
		}

		for _, word := range strings.Split(string(contents), "\n") {
			solveAndReport(pda, word)
		}
	}
}

func solveAndReport(pda *model.PDA, word string) {
	recognized, nonDet, stackTrace := solution.Solve(pda, word)
	fmt.Printf("Input: \"%s\"; Recognized? : %v; Non-Deterministic? : %v; Stack history:\n", word, recognized, nonDet)
	for _, v := range stackTrace {
		fmt.Printf("%v - %v\n", v.State, v.Stack)
	}
}
