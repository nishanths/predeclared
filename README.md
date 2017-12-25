# predeclared [![Build Status](https://travis-ci.org/nishanths/predeclared.svg?branch=master)](https://travis-ci.org/nishanths/predeclared)

Find code that overrides one of Go's predeclared identifiers (`new`, `make`, `append`, etc.).

The list of predeclared identifiers can be found in the [spec](https://golang.org/ref/spec#Predeclared_identifiers).

## Usage

```
go get -u github.com/nishanths/predeclared

predeclared [flags] [path ...]
```

See [godoc](https://godoc.org/github.com/nishanths/predeclared) or `predeclared -h` for more.

## Example

Given a file:

```
package main

import "log"

// welp, the order of the parameters is different from the builtin
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

// welp, not the builtin print.
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
