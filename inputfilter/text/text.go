package text

import (
	"github.com/speedata/decorate/processor"
)

type inputfilter struct {
}

func init() {
	processor.RegisterInputFilter("text", inputfilter{})
}

// The Tokenizer struct contains raw material (using )
type Tokenizer struct {
	c   []*processor.Token
	pos int
}

// NextToken is called by the output filter until it returns nil.
func (t *Tokenizer) NextToken() *processor.Token {
	if t.pos >= len(t.c) {
		return nil
	}
	tok := t.c[t.pos]
	t.pos++
	return tok
}

// Highlight is called once for every input file.
func (f inputfilter) Highlight(data []byte) (processor.Tokenizer, error) {
	t := &Tokenizer{}

	tok := &processor.Token{
		Major: processor.MAJOR_RAW,
		Minor: 0,
		Value: string(data),
	}
	t.c = append(t.c, tok)

	return t, nil
}
