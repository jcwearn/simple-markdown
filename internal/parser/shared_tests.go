package parser

type TestCase struct {
	Name    string
	Input   string
	Want    string
	WantErr bool
}

var SharedTestCases = []TestCase{
	{
		Name:  "empty input",
		Input: "",
		Want:  "",
	},
	{
		Name:  "new lines - 1",
		Input: "\n",
		Want:  "",
	},
	{
		Name:  "new lines - 2",
		Input: "\n\n",
		Want:  "",
	},
	{
		Name:  "new lines - 3",
		Input: "\n\n\n",
		Want:  "",
	},
	{
		Name:  "new lines - 4",
		Input: "\n\n\n\n",
		Want:  "",
	},
	{
		Name:  "new lines - interspersed in real input",
		Input: "first paragraph\n\n\n\nsecond paragraph",
		Want:  "<p>first paragraph</p>\n\n<p>second paragraph</p>",
	},
	{
		Name:  "h1",
		Input: "# Header one",
		Want:  "<h1>Header one</h1>",
	},
	{
		Name:  "h2",
		Input: "## Header two",
		Want:  "<h2>Header two</h2>",
	},
	{
		Name:  "h3",
		Input: "### Header three",
		Want:  "<h3>Header three</h3>",
	},
	{
		Name:  "h4",
		Input: "#### Header four",
		Want:  "<h4>Header four</h4>",
	},
	{
		Name:  "h5",
		Input: "##### Header five",
		Want:  "<h5>Header five</h5>",
	},
	{
		Name:  "h6",
		Input: "###### Header six",
		Want:  "<h6>Header six</h6>",
	},
	{
		Name:  "header extra #",
		Input: "####### extra #",
		Want:  "<p>####### extra #</p>",
	},
	{
		Name:  "header leading whitespace",
		Input: "    ###### leading whitespace",
		Want:  "<h6>leading whitespace</h6>",
	},
	{
		Name:  "header trailing whitespace",
		Input: "######   ",
		Want:  "",
	},
	{
		Name:  "header with no text",
		Input: "######",
		Want:  "",
	},
	{
		Name:  "header - internal spacing",
		Input: "# this is a h1         with internal spacing",
		Want:  "<h1>this is a h1 with internal spacing</h1>",
	},
	{
		Name:  "paragraph",
		Input: "here is a paragraph",
		Want:  "<p>here is a paragraph</p>",
	},
	{
		Name:  "paragraph - with link",
		Input: "here is a paragraph [with a link](example.com)",
		Want:  "<p>here is a paragraph <a href=\"example.com\">with a link</a></p>",
	},
	{
		Name:  "paragraph - multiline",
		Input: "here is a paragraph\nwith two lines",
		Want:  "<p>here is a paragraph with two lines</p>",
	},
	{
		Name:  "paragraph - multiline - with an empty line",
		Input: "here is a paragraph\n\nwith two lines",
		Want:  "<p>here is a paragraph</p>\n\n<p>with two lines</p>",
	},
	{
		Name:  "paragraph - leading whitespace",
		Input: "    paragraph with leading whitespace",
		Want:  "<p>paragraph with leading whitespace</p>",
	},
	{
		Name:  "paragraph - random brackets characters",
		Input: "this ] is [ a ] ]paragraph [ with some brackets",
		Want:  "<p>this ] is [ a ] ]paragraph [ with some brackets</p>",
	},
	{
		Name:  "paragraph - random header",
		Input: "this is a paragraph ### with a random header",
		Want:  "<p>this is a paragraph ### with a random header</p>",
	},
	{
		Name:  "paragraph - internal spacing",
		Input: "this is a paragraph         with a random header",
		Want:  "<p>this is a paragraph with a random header</p>",
	},
	{
		Name:  "h1 - with link",
		Input: "# here is a header one [with a link](example.com)",
		Want:  "<h1>here is a header one <a href=\"example.com\">with a link</a></h1>",
	},
	{
		Name:  "empty link",
		Input: "[]()",
		Want:  "",
	},
	{
		Name:  "partial link - no link",
		Input: "[has text]()",
		Want:  "<p><a href=\"\">has text</a></p>",
	},
	{
		Name:  "partial link - no text",
		Input: "[](example.com)",
		Want:  "",
	},
	{
		Name:  "malformed link",
		Input: "[some text)(example.com]",
		Want:  "<p>[some text)(example.com]</p>",
	},
	{
		Name:  "link - internal spacing",
		Input: "[  some    link  ](  example.com  )",
		Want:  "<p><a href=\"example.com\">some link</a></p>",
	},
	{
		Name:  "link - internal spacing - invalid url",
		Input: "[  some    link  ](  exam   ple.com  )",
		Want:  "<p>[ some link ]( exam ple.com )</p>",
	},
	{
		Name:  "multiple links",
		Input: "[some link](example.com) [some other link](example.com)",
		Want:  "<p><a href=\"example.com\">some link</a> <a href=\"example.com\">some other link</a></p>",
	},
	{
		Name:  "multiple links - no spacing",
		Input: "[some link](example.com)[some other link](example.com)",
		Want:  "<p><a href=\"example.com\">some link</a><a href=\"example.com\">some other link</a></p>",
	},
	{
		Name:  "multiple links - one invalid - 1",
		Input: "[](example.com)[some other link](example.com)",
		Want:  "<p><a href=\"example.com\">some other link</a></p>",
	},
	{
		Name:  "multiple links - one invalid - 2",
		Input: "[some link](example.com)[](example.com)",
		Want:  "<p><a href=\"example.com\">some link</a></p>",
	},
	{
		Name:  "mailchimp input 1",
		Input: "# Sample Document\n\nHello!\n\nThis is sample markdown for the [Mailchimp](https://www.mailchimp.com) homework assignment.",
		Want:  "<h1>Sample Document</h1>\n\n<p>Hello!</p>\n\n<p>This is sample markdown for the <a href=\"https://www.mailchimp.com\">Mailchimp</a> homework assignment.</p>",
	},
	{
		Name:  "mailchimp input 2",
		Input: "# Header one\n\nHello there\n\nHow are you?\nWhat's going on?\n\n## Another Header\n\nThis is a paragraph [with an inline link](http://google.com). Neat, eh?\n\n## This is a header [with a link](http://yahoo.com)",
		Want:  "<h1>Header one</h1>\n\n<p>Hello there</p>\n\n<p>How are you? What's going on?</p>\n\n<h2>Another Header</h2>\n\n<p>This is a paragraph <a href=\"http://google.com\">with an inline link</a>. Neat, eh?</p>\n\n<h2>This is a header <a href=\"http://yahoo.com\">with a link</a></h2>",
	},
	{
		Name:  "special characters",
		Input: "<a href=\"\">%s</a>",
		Want:  "<p><a href=\"\">%s</a></p>",
	},
}
