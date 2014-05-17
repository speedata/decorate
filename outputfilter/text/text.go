package text

import (
	"github.com/speedata/decorate/processor"
)

type outputfilter struct{}

func init() {
	processor.RegisterOutputFilter("text", outputfilter{})
}

func (f outputfilter) Render(in chan processor.Token, out chan string) {
	for {
		select {
		case t, ok := <-in:
			if ok {
				out <- t.Value
			} else {
				close(out)
				return
			}
		}
	}
}
