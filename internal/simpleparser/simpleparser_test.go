package simpleparser

import (
	"testing"
)

func TestSimpleParser_ParseInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty input",
			input: "",
			want:  "",
		},
		{
			name:  "new lines - 1",
			input: "\n",
			want:  "",
		},
		{
			name:  "new lines - 2",
			input: "\n\n",
			want:  "",
		},
		{
			name:  "new lines - 3",
			input: "\n\n\n",
			want:  "",
		},
		{
			name:  "new lines - 4",
			input: "\n\n\n\n",
			want:  "",
		},
		{
			name:  "new lines - interspersed in real input",
			input: "first paragraph\n\n\n\nsecond paragraph",
			want:  "<p>first paragraph</p>\n\n<p>second paragraph</p>",
		},
		{
			name:  "h1",
			input: "# Header one",
			want:  "<h1>Header one</h1>",
		},
		{
			name:  "h2",
			input: "## Header two",
			want:  "<h2>Header two</h2>",
		},
		{
			name:  "h3",
			input: "### Header three",
			want:  "<h3>Header three</h3>",
		},
		{
			name:  "h4",
			input: "#### Header four",
			want:  "<h4>Header four</h4>",
		},
		{
			name:  "h5",
			input: "##### Header five",
			want:  "<h5>Header five</h5>",
		},
		{
			name:  "h6",
			input: "###### Header six",
			want:  "<h6>Header six</h6>",
		},
		{
			name:  "extra #",
			input: "####### extra #",
			want:  "<p>####### extra #</p>",
		},
		{
			name:  "paragraph",
			input: "here is a paragraph",
			want:  "<p>here is a paragraph</p>",
		},
		{
			name:  "paragraph - with link",
			input: "here is a paragraph [with a link](example.com)",
			want:  "<p>here is a paragraph <a href=\"example.com\">with a link</a></p>",
		},
		{
			name:  "paragraph - multiline",
			input: "here is a paragraph\nwith two lines",
			want:  "<p>here is a paragraph\nwith two lines</p>",
		},
		{
			name:  "paragraph - multiline - with an empty line",
			input: "here is a paragraph\n\nwith two lines",
			want:  "<p>here is a paragraph</p>\n\n<p>with two lines</p>",
		},
		{
			name:  "h1 - with link",
			input: "# here is a header one [with a link](example.com)",
			want:  "<h1>here is a header one <a href=\"example.com\">with a link</a></h1>",
		},
		{
			name:  "empty link",
			input: "[]()",
			want:  "<p>[]()</p>",
		},
		{
			name:  "partial link - no link",
			input: "[has text]()",
			want:  "<p>[has text]()</p>",
		},
		{
			name:  "partial link - no text",
			input: "[](example.com)",
			want:  "<p>[](example.com)</p>",
		},
		{
			name:  "malformed link",
			input: "[some text)(example.com]",
			want:  "<p>[some text)(example.com]</p>",
		},
		{
			name:  "mailchimp input 1",
			input: "# Sample Document\n\nHello!\n\nThis is sample markdown for the [Mailchimp](https://www.mailchimp.com) homework assignment.",
			want:  "<h1>Sample Document</h1>\n\n<p>Hello!</p>\n\n<p>This is sample markdown for the <a href=\"https://www.mailchimp.com\">Mailchimp</a> homework assignment.</p>",
		},
		{
			name:  "mailchimp input 2",
			input: "# Header one\n\nHello there\n\nHow are you?\nWhat's going on?\n\n## Another Header\n\nThis is a paragraph [with an inline link](http://google.com). Neat, eh?\n\n## This is a header [with a link](http://yahoo.com)",
			want:  "<h1>Header one</h1>\n\n<p>Hello there</p>\n\n<p>How are you?\nWhat's going on?</p>\n\n<h2>Another Header</h2>\n\n<p>This is a paragraph <a href=\"http://google.com\">with an inline link</a>. Neat, eh?</p>\n\n<h2>This is a header <a href=\"http://yahoo.com\">with a link</a></h2>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			simpleParser := NewParser()

			got := simpleParser.ParseInput(tt.input)
			if got != tt.want {
				t.Errorf("SimpleParser.ParseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
