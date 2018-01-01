# predeclared [![Build Status](https://travis-ci.org/nishanths/predeclared.svg?branch=master)](https://travis-ci.org/nishanths/predeclared)

Find code that overrides one of Go's predeclared identifiers (`new`, `make`, `append`, etc.).

The list of predeclared identifiers can be found in the [spec](https://golang.org/ref/spec#Predeclared_identifiers).

## Usage

```
go get -u github.com/nishanths/predeclared

predeclared [flags] [path ...]
```

See [godoc](https://godoc.org/github.com/nishanths/predeclared) or `predeclared -h` for more.

## Examples

Given a file:

```go
package main

import "log"

// welp, the order of the parameters is different from the built-in
// copy function!
func copy(src, dst []T) {
	for i := range dst {
		if i == len(src) {
			break
		}
		string := src[i].s
		dst[i].s = string
	}
}

// welp, not the built-in print.
func print(t *T) { log.Printf("{ x=%d, y=%d }", t.x, t.y) }
```

running:

```
predeclared example.go
```

prints:

```
example.go:7:6: function "copy" has same name as predeclared identifier
example.go:12:3: variable "string" has same name as predeclared identifier
example.go:18:6: function "print" has same name as predeclared identifier
```

In the standard libary's `text` package:

```sh
$ predeclared /usr/local/go/src/text
/usr/local/go/src/text/template/exec_test.go:209:21: param "error" has same name as predeclared identifier
/usr/local/go/src/text/template/parse/node.go:496:33: param "true" has same name as predeclared identifier
/usr/local/go/src/text/template/parse/node.go:537:3: variable "rune" has same name as predeclared identifier
/usr/local/go/src/text/template/template.go:215:30: param "new" has same name as predeclared identifier
```
