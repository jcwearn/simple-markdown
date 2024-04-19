# simpleparser
This is a very basic implemtation of a markdown parser.  It only supports a subset of markdown defined [here](https://gist.github.com/mc-interviews/305a6d7d8c4ba31d4e4323e574135bf9#formatting-specifics).

## Usage
To run simpleparser as a webserver use the following command:
```bash
go run main.go webserver
```

## Sample Request
```
curl --location 'localhost:4000/v1/parse' \
--header 'Content-Type: text/plain' \
--data '# Header one

Hello there

How are you?
What'\''s going on?

## Another Header

This is a paragraph [with an inline link](http://google.com). Neat, eh?

## This is a header [with a link](http://yahoo.com)'
```