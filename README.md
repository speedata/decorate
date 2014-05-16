decorate
========

Go syntax highlighting framework


Installation
------------

`go get github.com/speedata/decorate`


How to use decorate?
-------------------

See the application `gohigh` for an example: https://github.com/speedata/gohigh

```go
decorate.HighlightFile(inputfile, inputfilter, outputfilter)
```

where `inputfile` is the filename (+ path) to the input file, `inputfilter` the name of the input filter (currently only `text` and `xml` are supported) and `outputfilter` is currently either `html` or `text`.



Other:
-----

Status: pre alpha<br>
Supported/maintained: yes<br>
Contribution welcome: yes (pull requests, issues)<br>
Main page: https://github.com/speedata/decorate<br>
License: MIT<br>
Contact: gundlach@speedata.de<br>

