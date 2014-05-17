package xml

import (
	"bufio"
	"bytes"
	"github.com/speedata/decorate/processor"
	"strings"
	"unicode"
	"unicode/utf8"
)

type inputfilter struct {
}

func init() {
	processor.RegisterInputFilter("xml", inputfilter{})
}

// tokenizer implements the processor.Tokenizer interface which is used
// in the output filter
type tokenizer struct {
	c   []*processor.Token
	pos int
}

func (t *tokenizer) appendToken(typ processor.TokenMajor, text string) {
	prev_pos := len(t.c) - 1
	if prev_pos >= 0 {
		prev_token := t.c[prev_pos]
		if prev_token.Major == typ {
			prev_token.Value = prev_token.Value + text
			return
		}
	}
	tok := &processor.Token{}
	tok.Major = typ
	tok.Value = text
	t.c = append(t.c, tok)
}

// The output filter calls NextToken until no token is left over
func (t *tokenizer) NextToken() *processor.Token {
	if t.pos >= len(t.c) {
		return nil
	}
	tok := t.c[t.pos]
	t.pos++
	return tok
}

var (
	TCOMMENT = []byte{'<', '!', '-', '-'}
)

func nameboundary(r rune) bool {
	return unicode.IsSpace(r) || r == '=' || r == '/'
}

func tokenizeXML(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if bytes.HasPrefix(data, TCOMMENT) {
		return len(TCOMMENT), TCOMMENT, nil
	}
	r, size := utf8.DecodeRune(data)
	if unicode.IsSpace(r) {
		return size, data[:size], nil
	}
	if data[0] == '<' || data[0] == '>' {
		return 1, data[:1], nil
	}
	if data[0] == '/' && data[1] == '>' {
		return 2, data[:2], nil
	}
	num := bytes.IndexFunc(data, nameboundary)
	if num > 0 {
		return num, data[:num], nil
	}
	return 1, data[:1], nil

}

func (f inputfilter) Highlight(data []byte) (processor.Tokenizer, error) {
	t := &tokenizer{}
	buf := bytes.NewBuffer(data)
	const (
		RAW = iota
		COMMENT
		STRING
		TAGSTART
		TAG
	)
	state := RAW
	scanner := bufio.NewScanner(buf)
	scanner.Split(tokenizeXML)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, `"`) {
			state = STRING
			t.appendToken(processor.MAJOR_STRING, text)
			continue
		}
		switch text {
		case "<!--":
			t.appendToken(processor.MAJOR_COMMENT, text)
			state = COMMENT
		case "-->":
			t.appendToken(processor.MAJOR_COMMENT, text)
			state = RAW
		case "<":
			t.appendToken(processor.MAJOR_NAME, text)
			state = TAGSTART
		case " ", "\n":
			switch state {
			case COMMENT:
				t.appendToken(processor.MAJOR_COMMENT, text)
			case TAGSTART:
				t.appendToken(processor.MAJOR_RAW, text)
				state = TAG
			default:
				t.appendToken(processor.MAJOR_RAW, text)
			}
		default:
			switch state {
			case COMMENT:
				t.appendToken(processor.MAJOR_COMMENT, text)
			case TAGSTART:
				t.appendToken(processor.MAJOR_NAME, text)
			case TAG:
				t.appendToken(processor.MAJOR_NAME, text)
			default:
				t.appendToken(processor.MAJOR_RAW, text)
			}
		}
	}
	return t, nil
}
