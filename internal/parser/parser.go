package parser

type Parser interface {
	ParseInput(input string) (string, error)
}
