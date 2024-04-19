package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jcwearn/simple-markdown/internal/simpleparser"
)

const PROMPT = ">> "

func Start(parser simpleparser.SimpleParser) {
	in := os.Stdin
	out := os.Stdout
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		fmt.Fprintf(out, fmt.Sprintln(parser.ParseInput(line)))
	}
}
