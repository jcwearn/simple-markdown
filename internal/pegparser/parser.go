package pegparser

import (
	"errors"
	"strings"
)

func ParseInput(input string) (string, error) {
	result, err := Parse("", []byte(input))
	if err != nil {
		return "", err
	}

	if md, ok := result.(Markdown); ok {
		return strings.TrimSpace(strings.Join(md.Lines, "\n\n")), nil
	}
	return "", errors.New("failed to parse cast result as Markdown")
}
