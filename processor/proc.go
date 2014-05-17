package processor

import (
	"errors"
	"strings"
)

// The input filter is expected to implement function Highlight which returns a Tokenizer.
type InputFilter interface {
	Highlight([]byte, chan Token)
}

type TypeMajor int
type TypeMinor int

// This is the basic data structure in the intermediate format between
// the input filter and output filter.
type Token struct {
	Major TypeMajor
	Minor TypeMinor
	Value string
}

type HighlightFunc func([]byte, chan Token)
type RenderFunc func(chan Token, chan string)

var (
	inputfilters  map[string]HighlightFunc
	outputfilters map[string]RenderFunc
)

func init() {
	inputfilters = make(map[string]HighlightFunc)
	outputfilters = make(map[string]RenderFunc)
}

// These are the allowed token types
const (
	MAJOR_RAW TypeMajor = iota
	MAJOR_COMMENT
	MAJOR_STRING
	MAJOR_ERROR
	MAJOR_GENERIC
	MAJOR_KEYWORD
	MAJOR_NAME
	MAJOR_NUMBER
	MAJOR_VARIABLE
	MINOR_RAW TypeMinor = iota
	MINOR_NAME_TAG
	MINOR_NAME_ATTRIBUTE
)

// All lexers are required to call this function exactly once.
func RegisterInputFilter(name string, f HighlightFunc) {
	inputfilters[name] = f
}

// All output filters are required to call this function exactly once.
func RegisterOutputFilter(name string, f RenderFunc) {
	outputfilters[name] = f
}

// Return a list of available input filters.
func InputFilters() []string {
	ret := make([]string, 0, len(inputfilters))
	for v, _ := range inputfilters {
		ret = append(ret, v)
	}
	return ret
}

// Return a list of available output filters.
func OutputFilters() []string {
	ret := make([]string, 0, len(outputfilters))
	for v, _ := range outputfilters {
		ret = append(ret, v)
	}
	return ret
}

// Run the given input and output filters on the source and return a
// string of the highlighted input source and nil or, if there is an error,
// a perhaps empty string and an error.
func Highlight(inputfilter, outputfilter string, source []byte) (string, error) {
	ifilter, ok := inputfilters[inputfilter]
	if !ok {
		return "", errors.New("Input filter not declared")
	}
	ofilter, ok := outputfilters[outputfilter]
	if !ok {
		return "", errors.New("Output filter not declared")
	}

	chain := make(chan Token, 0)
	res := make(chan string, 0)
	go ifilter(source, chain)
	go ofilter(chain, res)
	var ret []string
	for {
		select {
		case str, ok := <-res:
			if ok {
				ret = append(ret, str)
			} else {
				return strings.Join(ret, ""), nil
			}
		}
	}
	return "", nil
}
