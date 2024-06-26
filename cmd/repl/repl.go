package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jcwearn/simple-markdown/internal/parser"
)

const PROMPT = ">> "

func Start(parser parser.Parser) {
	in := os.Stdin
	out := os.Stdout
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		parsed, err := parser.ParseInput(line)
		if err != nil {
			fmt.Fprint(out, fmt.Sprintln(err))
		} else {
			fmt.Fprint(out, fmt.Sprintln(parsed))
		}
	}
}
