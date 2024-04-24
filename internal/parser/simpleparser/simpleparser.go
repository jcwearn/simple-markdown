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
			text        = strings.TrimSpace(line)
		)

		linkAllSubmatch := p.linkRegex.FindAllStringSubmatch(line, -1)
		linkSubexpNames := p.linkRegex.SubexpNames()

		for _, linkSubmatch := range linkAllSubmatch {
			if len(linkSubmatch) > 0 {
				for i, name := range linkSubexpNames {
					switch name {
					case "bracket":
						bracketText = linkSubmatch[i]

						if bracketText == "" {
							text = strings.Replace(text, linkSubmatch[0], "", 1)
						}
					case "paren":
						urlText = linkSubmatch[i]
						if urlText == "" {
							text = strings.Replace(text, linkSubmatch[0], fmt.Sprintf("<a href=\"\">%s</a>", bracketText), 1)
						} else {
							var replacement string
							urlRemovedSpaces := removeExtraSpaces(urlText)
							urlTrimmed := strings.TrimSpace(urlText)
							if urlRemovedSpaces != urlTrimmed {
								replacement = fmt.Sprintf("[%s](%s)", bracketText, urlText)
							} else {
								bracketTextTrimmed := strings.TrimSpace(bracketText)
								replacement = fmt.Sprintf("<a href=\"%s\">%s</a>", urlTrimmed, bracketTextTrimmed)
							}
							text = strings.Replace(text, linkSubmatch[0], replacement, 1)
						}
					}
				}
			}
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
			text = removeExtraSpaces(text)
			if headerLevel > 0 {
				formattedLine = fmt.Sprintf("<h%d>%s</h%d>\n\n", headerLevel, text, headerLevel)
			} else {
				if !isPoundString(text) {
					formattedLine = fmt.Sprintf("<p>%s</p>\n\n", text)
				}
			}
		}
		output += formattedLine
	}
	return strings.TrimSpace(output), nil
}

func removeExtraSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func isPoundString(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != '#' {
			return false
		}
	}
	return true
}
