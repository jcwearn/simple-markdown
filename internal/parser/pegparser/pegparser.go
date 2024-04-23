package pegparser

import (
	"errors"
	"strings"
)

type (
	PegParserConfig struct {
		Debug bool
	}
	PegParser struct {
		debug bool
	}
)

func NewParser(cfg PegParserConfig) PegParser {
	return PegParser{
		debug: cfg.Debug,
	}
}

func (pp PegParser) ParseInput(input string) (string, error) {
	result, err := Parse("", []byte(input))
	if err != nil {
		return "", err
	}

	if md, ok := result.(Markdown); ok {
		return strings.TrimSpace(strings.Join(md.Lines, "\n\n")), nil
	}
	return "", errors.New("failed to parse cast result as Markdown")
}
