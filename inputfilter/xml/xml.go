package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/speedata/decorate/processor"
	"log"
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

func (t *tokenizer) appendToken(typ processor.Tokentype, text string) {
	prev_pos := len(t.c) - 1
	if prev_pos >= 0 {
		prev_token := t.c[prev_pos]
		if prev_token.Typ == typ {
			prev_token.Value = prev_token.Value + text
			return
		}
	}
	tok := &processor.Token{}
	tok.Typ = typ
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

func (f inputfilter) Highlight(data []byte) (processor.Tokenizer, error) {
	// we should not use xml.Decoder for that purpose, because it's not a 1:1 copy of the input
	// but for a start, it's better than nothing, or?
	t := &tokenizer{}
	r := bytes.NewReader(data)
	decoder := xml.NewDecoder(r)
	for {
		tok, err := decoder.RawToken()
		if err != nil {
			break
		}
		switch v := tok.(type) {
		case xml.StartElement:
			t.appendToken(processor.NAMETAG, fmt.Sprintf("<%s", v.Name.Local))
			for _, v := range v.Attr {
				t.appendToken(processor.RAW, " ")
				t.appendToken(processor.NAMEATTRIBUTE, v.Name.Local+"=")
				t.appendToken(processor.LITERALSTRING, fmt.Sprintf(`"%s"`, v.Value))
			}
			t.appendToken(processor.NAMETAG, fmt.Sprintf(">"))
		case xml.EndElement:
			t.appendToken(processor.NAMETAG, fmt.Sprintf(`</%s>`, v.Name.Local))
		case xml.CharData:
			t.appendToken(processor.RAW, string(v.Copy()))
		default:
			log.Printf(">>> %T", v)
		}
	}
	return t, nil
}
