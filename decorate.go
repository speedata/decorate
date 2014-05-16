package decorate

import (
	"github.com/speedata/decorate/processor"
	"io/ioutil"

	_ "github.com/speedata/decorate/inputfilter/text"
	_ "github.com/speedata/decorate/inputfilter/xml"

	_ "github.com/speedata/decorate/outputfilter/html"
	_ "github.com/speedata/decorate/outputfilter/text"
)

// If inputfilter is the empty string, it should be guessed from the input string.
func HighlightFile(filename string, inputfilter string, outputfilter string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	ret, err := processor.Highlight(inputfilter, outputfilter, data)
	if err != nil {
		return "", err
	}
	return ret, nil
}
