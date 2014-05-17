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
	tagnames := map[processor.TypeMajor]string{
		processor.MAJOR_RAW:      "raw",
		processor.MAJOR_COMMENT:  "comment",
		processor.MAJOR_STRING:   "string",
		processor.MAJOR_ERROR:    "error",
		processor.MAJOR_GENERIC:  "generic",
		processor.MAJOR_KEYWORD:  "keyword",
		processor.MAJOR_NAME:     "name",
		processor.MAJOR_NUMBER:   "number",
		processor.MAJOR_VARIABLE: "variable",
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
