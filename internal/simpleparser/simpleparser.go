package simpleparser

import (
	"fmt"
	"regexp"
	"strings"
)

type (
	SimpleParser struct {
		hRegex *regexp.Regexp
		aRegex *regexp.Regexp
	}
)

const headerRegex = `(?P<header>^#{1,6})(?P<text>\s(.*)$)`
const linkRegex = `\[([^\]]+)\]\(([^)]+)\)`

func NewParser() SimpleParser {
	hRegex := regexp.MustCompile(headerRegex)
	aRegex := regexp.MustCompile(linkRegex)
	return SimpleParser{
		hRegex: hRegex,
		aRegex: aRegex,
	}
}

func (p SimpleParser) ParseInput(input string) string {
	if strings.TrimSpace(input) == "" {
		return ""
	}

	var output string
	lines := strings.Split(input, "\n\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		var (
			text        = p.getTextWithLink(line)
			headerLevel int
		)

		headingSubmatch := p.hRegex.FindStringSubmatch(text)
		subexpNames := p.hRegex.SubexpNames()

		if len(headingSubmatch) > 0 {
			for i, name := range subexpNames {
				switch name {
				case "header":
					headerLevel = len(headingSubmatch[i])
				case "text":
					text = strings.TrimSpace(headingSubmatch[i])
				}
			}
		}

		var formattedLine string
		if headerLevel > 0 {
			formattedLine = fmt.Sprintf("<h%d>%s</h%d>\n\n", headerLevel, text, headerLevel)
		} else {
			formattedLine = fmt.Sprintf("<p>%s</p>\n\n", text)
		}
		output += formattedLine
	}
	return strings.TrimSpace(output)
}

func (p SimpleParser) getTextWithLink(s string) string {
	return p.aRegex.ReplaceAllString(s, "<a href=\"$2\">$1</a>")
}
