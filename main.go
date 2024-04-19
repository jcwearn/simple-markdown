package main

import (
	"fmt"
	"os"

	"github.com/jcwearn/simple-markdown/cmd/repl"
	"github.com/jcwearn/simple-markdown/cmd/webserver"
	"github.com/jcwearn/simple-markdown/internal/simpleparser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: simple <command> [arguments]")
		os.Exit(1)
	}

	parser := simpleparser.NewParser()

	switch os.Args[1] {
	case "webserver":
		ws := webserver.NewWebServer(parser)
		ws.Start()
	case "repl":
		repl.Start(parser)
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}
