package text

import (
	"github.com/speedata/decorate/processor"
)

type inputfilter struct {
}

func init() {
	processor.RegisterInputFilter("text", inputfilter{})
}

type Tokenizer struct {
	c   []*processor.Token
	pos int
}

func (t *Tokenizer) AppendToken(typ int, text string) {
	tok := &processor.Token{}
	tok.Typ = typ
	tok.Value = text
	t.c = append(t.c, tok)
}

func (t *Tokenizer) NextToken() *processor.Token {
	if t.pos >= len(t.c) {
		return nil
	}
	tok := t.c[t.pos]
	t.pos++
	return tok
}

func (f inputfilter) Highlight(data []byte) (processor.Tokenizer, error) {
	t := &Tokenizer{}
	t.AppendToken(processor.RAW, string(data))
	return t, nil
}
