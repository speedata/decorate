package html

import (
	"fmt"
	"github.com/speedata/decorate/processor"
	"html"
	"strings"
)

type outputfilter struct{}

func init() {
	processor.RegisterOutputFilter("html", outputfilter{})
}

// Gets called when the user requests HTML output
func (f outputfilter) Render(t processor.Tokenizer) string {
	classes := map[processor.Tokentype]string{
		processor.NAMETAG:       "nt",
		processor.NAMEATTRIBUTE: "na",
		processor.LITERALSTRING: "s",
	}

	var out []string
	out = append(out, fmt.Sprint(`<div class="highlight"><pre>`))

	for {
		t := t.NextToken()
		if t == nil {
			break
		}
		if t.Typ == processor.RAW {
			out = append(out, t.Value)
		} else {
			out = append(out, fmt.Sprintf(`<span class="%s">%s</span>`, classes[t.Typ], html.EscapeString(t.Value)))
		}
	}
	out = append(out, fmt.Sprint(`</pre></div>`))
	return strings.Join(out, "")
}
