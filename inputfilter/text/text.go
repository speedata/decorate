package text

import (
	"github.com/speedata/decorate/processor"
)

type inputfilter struct {
}

func init() {
	processor.RegisterInputFilter("text", inputfilter{})
}

// Highlight is called once for every input file.
func (f inputfilter) Highlight(data []byte, out chan processor.Token) {
	tok := processor.Token{}
	tok.Major = processor.MAJOR_RAW
	tok.Minor = processor.MINOR_RAW
	tok.Value = string(data)
	out <- tok
	close(out)
}
