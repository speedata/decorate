package processor

import (
	"errors"
)

// The input filter is expected to implement function Highlight which returns a Tokenizer.
type InputFilter interface {
	Highlight([]byte) (Tokenizer, error)
}

type OutputFilter interface {
	Render(Tokenizer) string
}

// The output filter calls NextToken until it returns nil.
type Tokenizer interface {
	NextToken() *Token
}

type Tokentype int

// This is the basic data structure in the intermediate format between
// the input filter and output filter.
type Token struct {
	Typ   Tokentype
	Value string
}

var (
	inputfilters  map[string]InputFilter
	outputfilters map[string]OutputFilter
)

func init() {
	inputfilters = make(map[string]InputFilter)
	outputfilters = make(map[string]OutputFilter)
}

// These are the allowed token types
const (
	RAW Tokentype = iota
	NAMETAG
	NAMEATTRIBUTE
	LITERALSTRING
)

// All lexers are required to call this function exactly once.
func RegisterInputFilter(name string, filter InputFilter) {
	inputfilters[name] = filter
}

// All output filters are required to call this function exactly once.
func RegisterOutputFilter(name string, filter OutputFilter) {
	outputfilters[name] = filter
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
	ifilter := inputfilters[inputfilter]
	if ifilter == nil {
		return "", errors.New("Input filter not declared")
	}

	ofilter := outputfilters[outputfilter]
	if ofilter == nil {
		return "", errors.New("Output filter not declared")
	}

	tokenizer, err := ifilter.Highlight(source)
	if err != nil {
		return "", err
	}
	ret := ofilter.Render(tokenizer)
	return ret, nil
}
