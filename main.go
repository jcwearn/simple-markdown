package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jcwearn/simple-markdown/cmd/repl"
	"github.com/jcwearn/simple-markdown/cmd/webserver"
	"github.com/jcwearn/simple-markdown/internal/parser"
	"github.com/jcwearn/simple-markdown/internal/parser/pegparser"
	"github.com/jcwearn/simple-markdown/internal/parser/simpleparser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: simple <command>")
		os.Exit(1)
	}

	webserverCmd := flag.NewFlagSet("webserver", flag.ExitOnError)
	webserverAddr := webserverCmd.String("addr", ":4000", "HTTP network address")

	replCmd := flag.NewFlagSet("repl", flag.ExitOnError)
	replParser := replCmd.String("parser", "simple", "Parser used for the repl [simple, peg]")

	simpleParser := simpleparser.NewParser()
	pegParser := pegparser.NewParser(pegparser.PegParserConfig{Debug: false})

	switch os.Args[1] {
	case "webserver":
		webserverCmd.Parse(os.Args[2:])
		ws := webserver.NewWebServer(webserver.WebServerConfig{
			Address:      *webserverAddr,
			SimpleParser: simpleParser,
			PegParser:    pegParser,
		})
		ws.Start()
	case "repl":
		replCmd.Parse(os.Args[2:])
		var parser parser.Parser
		if *replParser == "peg" {
			parser = pegParser
		} else {
			parser = simpleParser
		}
		repl.Start(parser)
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}
