package debug

import (
	"fmt"
	"github.com/speedata/decorate/processor"
	"strings"
)

type outputfilter struct{}

func init() {
	processor.RegisterOutputFilter("debug", outputfilter{})
}

// Gets called when the user requests HTML output
func (f outputfilter) Render(t processor.Tokenizer) string {
	tagnames := map[processor.TokenMajor]string{
		processor.RAW:           "(raw)",
		processor.TCOMMENT:      "c",
		processor.NAMETAG:       "nt",
		processor.NAMEATTRIBUTE: "na",
		processor.LITERALSTRING: "s",
	}

	var out []string

	for {
		t := t.NextToken()
		if t == nil {
			break
		}
		out = append(out, fmt.Sprintf("%-5s: %q\n", tagnames[t.Major], t.Value))
	}
	return strings.Join(out, "")
}
