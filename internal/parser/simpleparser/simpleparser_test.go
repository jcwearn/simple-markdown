package simpleparser

import (
	"testing"

	p "github.com/jcwearn/simple-markdown/internal/parser"
)

func TestSimpleParser_ParseInput(t *testing.T) {
	for _, tt := range p.SharedTestCases {
		t.Run(tt.Name, func(t *testing.T) {
			simpleParser := NewParser()
			got, err := simpleParser.ParseInput(tt.Input)
			if (err != nil) != tt.WantErr {
				t.Errorf("PegParser.Parse() = %v", err)
			}
			if got != tt.Want {
				t.Errorf("SimpleParser.ParseInput() = %v, want %v", got, tt.Want)
			}
		})
	}
}
