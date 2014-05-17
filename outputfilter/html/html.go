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
	classes_major := map[processor.TypeMajor]string{
		processor.MAJOR_COMMENT:  "c",
		processor.MAJOR_STRING:   "s",
		processor.MAJOR_ERROR:    "err",
		processor.MAJOR_GENERIC:  "gen",
		processor.MAJOR_KEYWORD:  "kw",
		processor.MAJOR_NAME:     "name",
		processor.MAJOR_NUMBER:   "num",
		processor.MAJOR_VARIABLE: "var",
	}
	classes_minor := map[processor.TypeMinor]string{
		processor.MINOR_NAME_ATTRIBUTE: "natt",
		processor.MINOR_NAME_TAG:       "ntag",
	}

	var out []string
	out = append(out, fmt.Sprint(`<div class="highlight"><pre>`))
	var cls string

	for {
		t := t.NextToken()
		if t == nil {
			break
		}
		if t.Major == processor.MAJOR_RAW {
			out = append(out, t.Value)
		} else {
			if t.Minor == processor.MINOR_RAW {
				cls = classes_major[t.Major]
			} else {
				cls = strings.Join([]string{classes_major[t.Major], classes_minor[t.Minor]}, " ")
			}
			out = append(out, fmt.Sprintf(`<span class="%s">%s</span>`, cls, html.EscapeString(t.Value)))
		}
	}
	out = append(out, fmt.Sprint(`</pre></div>`))
	return strings.Join(out, "")
}
