package pegparser

import (
	"testing"

	p "github.com/jcwearn/simple-markdown/internal/parser"
)

func TestPegParser_ParseInput(t *testing.T) {
	for _, tt := range p.SharedTestCases {
		t.Run(tt.Name, func(t *testing.T) {
			pegParser := NewParser(PegParserConfig{Debug: false})
			got, err := pegParser.ParseInput(tt.Input)

			if (err != nil) != tt.WantErr {
				t.Errorf("PegParser.Parse() = %v", err)
			}

			if got != tt.Want {
				t.Errorf("PegParser.Parse() = %v, want %v", got, tt.Want)
			}
		})
	}
}

func BenchmarkPegParser_ParseInput(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pegParser := NewParser(PegParserConfig{Debug: false})
		pegParser.ParseInput(p.SharedTestCases[0].Input)
	}
}
