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

		text := parseLinks(line, p.linkRegex)
		text, headerLevel := parseHeader(text, p.headerRegex)

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

func parseLinks(line string, linkRegex *regexp.Regexp) string {
	var (
		bracketText string
		urlText     string
		text        = strings.TrimSpace(line)
	)

	linkAllSubmatch := linkRegex.FindAllStringSubmatch(line, -1)
	linkSubexpNames := linkRegex.SubexpNames()

	for _, linkSubmatch := range linkAllSubmatch {
		if len(linkSubmatch) > 0 {
			for i, name := range linkSubexpNames {
				switch name {
				case "bracket":
					bracketText = linkSubmatch[i]

					if bracketText == "" {
						// remove submatch if there is no text in []
						text = strings.Replace(text, linkSubmatch[0], "", 1)
					}
				case "paren":
					var replacement string
					urlText = linkSubmatch[i]
					if urlText == "" {
						// replace submatch with a link that doesn't have an href
						replacement = fmt.Sprintf("<a href=\"\">%s</a>", bracketText)
					} else {
						urlRemovedSpaces := removeExtraSpaces(urlText)
						urlTrimmed := strings.TrimSpace(urlText)
						if urlRemovedSpaces != urlTrimmed {
							// if there is whitespace in the link (meaning it's invalid)
							replacement = fmt.Sprintf("[%s](%s)", bracketText, urlText)
						} else {
							// typical replacement
							bracketTextTrimmed := strings.TrimSpace(bracketText)
							replacement = fmt.Sprintf("<a href=\"%s\">%s</a>", urlTrimmed, bracketTextTrimmed)
						}
					}
					text = strings.Replace(text, linkSubmatch[0], replacement, 1)
				}
			}
		}
	}
	return text
}

func parseHeader(line string, headerRegex *regexp.Regexp) (string, int) {
	var headerLevel int
	headingSubmatch := headerRegex.FindStringSubmatch(line)
	headerSubexpNames := headerRegex.SubexpNames()
	if len(headingSubmatch) > 0 {
		for i, name := range headerSubexpNames {
			switch name {
			case "header":
				headerLevel = len(headingSubmatch[i])
			case "text":
				line = strings.TrimSpace(headingSubmatch[i])
			}
		}
	}
	return line, headerLevel
}
