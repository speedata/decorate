package text

import (
	"github.com/speedata/decorate/processor"
	"strings"
)

type filter struct{}

func init() {
	processor.RegisterOutputFilter("text", filter{})
}

func (f filter) Render(t processor.Tokenizer) string {
	var out []string

	for {
		t := t.NextToken()
		if t == nil {
			break
		}
		out = append(out, t.Value)
	}
	return strings.Join(out, "")
}
