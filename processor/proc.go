package processor

import (
	"errors"
)

type InputFilter interface {
	Highlight([]byte) (Tokenizer, error)
}

type OutputFilter interface {
	Render(Tokenizer) string
}

type Tokenizer interface {
	AppendToken(int, string)
	NextToken() *Token
}

type Token struct {
	Typ   int
	Value string
}

var (
	inputfilters  map[string]InputFilter
	outputfilters map[string]OutputFilter
	Classes       map[int]string
)

func init() {
	inputfilters = make(map[string]InputFilter)
	outputfilters = make(map[string]OutputFilter)
	Classes = map[int]string{
		NAMETAG:       "nt",
		NAMEATTRIBUTE: "na",
		LITERALSTRING: "s",
	}
}

const (
	RAW = iota
	NAMETAG
	NAMEATTRIBUTE
	LITERALSTRING
)

func RegisterInputFilter(name string, filter InputFilter) {
	inputfilters[name] = filter
}
func RegisterOutputFilter(name string, filter OutputFilter) {
	outputfilters[name] = filter
}

func InputFilters() []string {
	ret := make([]string, 0, len(inputfilters))
	for v, _ := range inputfilters {
		ret = append(ret, v)
	}
	return ret
}

func OutputFilters() []string {
	ret := make([]string, 0, len(outputfilters))
	for v, _ := range outputfilters {
		ret = append(ret, v)
	}
	return ret
}

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
