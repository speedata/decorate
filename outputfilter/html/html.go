package html

import (
	"fmt"
	"github.com/speedata/decorate/processor"
	"html"
	"strings"
)

type filter struct{}

func init() {
	processor.RegisterOutputFilter("html", filter{})
}

func (f filter) Render(t processor.Tokenizer) string {
	var out []string
	out = append(out, fmt.Sprint(`<div class="highlight"><pre>`))

	for {
		t := t.NextToken()
		if t == nil {
			break
		}
		if t.Typ == processor.RAW {
			out = append(out, fmt.Sprint(t.Value))
		} else {
			out = append(out, fmt.Sprintf(`<span class="%s">%s</span>`, processor.Classes[t.Typ], html.EscapeString(t.Value)))
		}
	}
	out = append(out, fmt.Sprint(`</pre></div>`))
	return strings.Join(out, "")
}
