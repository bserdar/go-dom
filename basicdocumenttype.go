package dom

import (
	"bufio"
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

type BasicDocumentType struct {
	basicNode

	name     string
	publicID string
	systemID string

	defn string
}

func (dt *BasicDocumentType) GetName() string       { return dt.name }
func (dt *BasicDocumentType) GetPublicID() string   { return dt.publicID }
func (dt *BasicDocumentType) GetSystemID() string   { return dt.systemID }
func (dt *BasicDocumentType) GetDefinition() string { return dt.defn }

func scanDTDToken(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}
	// Scan until space, [, ], (, ), <, >,',' marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == '[' || r == ']' || r == '(' || r == ')' || r == '<' || r == '>' || r == ',' || r == '|' {
			if i == start {
				return i + width, data[start : i+width], nil
			}
			return i, data[start:i], nil
		}
		if unicode.IsSpace(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func parseDocumentType(content []byte) (DocumentType, bool, error) {
	scanner := bufio.NewScanner(bytes.NewReader(content))
	scanner.Split(scanDTDToken)
	state := 0
	done := false
	var ret *BasicDocumentType
	builder := strings.Builder{}
	lastTokenWord := false
	for scanner.Scan() {
		if done {
			break
		}
		token := scanner.Text()
		switch state {
		case 0:
			if token != "!DOCTYPE" {
				return nil, false, nil
			}
			ret = &BasicDocumentType{}
			state = 1

		case 1: // !DOCTYPE seen
			ret.name = token
			state = 2

		case 2: // doctype name seen
			if token == "SYSTEM" {
				state = 4
			} else if token == "PUBLIC" {
				state = 3
			} else if token == "[" {
				state = 10
			} else {
				return nil, true, ErrDOM{
					Typ: SYNTAX_ERR,
					Msg: "Document type syntax error",
				}
			}

		case 3: // PUBLIC seen
			ret.publicID = token
			state = 4

		case 4: // systemid
			ret.systemID = token
			state = 6

		case 6: // intSubset or close?
			if token == ">" {
				done = true
			} else if token == "[" {
				state = 10
			} else {
				return nil, true, ErrDOM{
					Typ: SYNTAX_ERR,
					Msg: "Document type syntax error",
				}
			}

		case 10:
			if token == "," || token == "|" || token == "(" || token == ")" || token == "<" || token == ">" {
				lastTokenWord = false
				builder.WriteString(token)
			} else if token == "]" {
				done = true
			} else {
				if lastTokenWord {
					builder.WriteRune(' ')
				}
				builder.WriteString(token)
				lastTokenWord = true
			}
		}
	}
	ret.defn = builder.String()

	if err := scanner.Err(); err != nil {
		return nil, true, ErrDOM{
			Typ: SYNTAX_ERR,
			Msg: err.Error(),
		}
	}
	return ret, true, nil
}
