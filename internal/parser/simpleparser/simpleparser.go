package simpleparser

import (
	"fmt"
	"regexp"
	"strings"
)

type SimpleParser struct {
	headerRegex *regexp.Regexp
	linkRegex   *regexp.Regexp
}

const headerRegex = `(?P<header>^#{1,6})(?P<text>\s(.*)$)`
const linkRegex = `\[(?P<bracket>[^\]]*)\]\((?P<paren>[^)]*)\)`

func NewParser() SimpleParser {
	hRegex := regexp.MustCompile(headerRegex)
	aRegex := regexp.MustCompile(linkRegex)
	return SimpleParser{
		headerRegex: hRegex,
		linkRegex:   aRegex,
	}
}

func (p SimpleParser) ParseInput(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", nil
	}

	var output string
	lines := strings.Split(input, "\n\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		var (
			headerLevel int
			bracketText string
			urlText     string
			text        = line
		)

		linkSubmatch := p.linkRegex.FindStringSubmatch(line)
		linkSubexpNames := p.linkRegex.SubexpNames()
		if len(linkSubmatch) > 0 {
			for i, name := range linkSubexpNames {
				switch name {
				case "bracket":
					bracketText = strings.TrimSpace(linkSubmatch[i])
				case "paren":
					urlText = strings.TrimSpace(linkSubmatch[i])
				}
			}

			if bracketText == "" {
				text = p.linkRegex.ReplaceAllString(text, "")
			}

			if urlText == "" {
				text = p.linkRegex.ReplaceAllString(text, "<a href=\"\">$1</a>")
			}

			text = p.linkRegex.ReplaceAllString(text, "<a href=\"$2\">$1</a>")
		}

		headingSubmatch := p.headerRegex.FindStringSubmatch(text)
		headerSubexpNames := p.headerRegex.SubexpNames()
		if len(headingSubmatch) > 0 {
			for i, name := range headerSubexpNames {
				switch name {
				case "header":
					headerLevel = len(headingSubmatch[i])
				case "text":
					text = strings.TrimSpace(headingSubmatch[i])
				}
			}
		}

		var formattedLine string
		if text != "" {
			text = strings.TrimSpace(text)
			if headerLevel > 0 {
				formattedLine = fmt.Sprintf("<h%d>%s</h%d>\n\n", headerLevel, text, headerLevel)
			} else {
				formattedLine = fmt.Sprintf("<p>%s</p>\n\n", text)
			}
		}
		output += formattedLine
	}
	return strings.TrimSpace(output), nil
}
