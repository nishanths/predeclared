# predeclared [![Build Status](https://travis-ci.org/nishanths/predeclared.svg?branch=master)](https://travis-ci.org/nishanths/predeclared)

Find declarations that override one of Go's predeclared identifiers (`new`, `make`, `append`, etc.).

The list of predeclared identifiers can be found in the [spec](https://golang.org/ref/spec#Predeclared_identifiers).

## Usage

```
$ go get -u github.com/nishanths/predeclared
$ predeclared file1.go file2.go dir1
```

See [godoc](https://godoc.org/github.com/nishanths/predeclared) or `predeclared -h` for more.

## Example

Given a file:

```
package print

func append(s int) {
	copy := s
	x.Push(copy)
}

type F interface {
	new() T
}

func (p Pool) new() {}
```

running:

```
predeclared -q file.go
```

prints:

```
testdata/example.go:1:9: package name "print" has same name as predeclared identifier
testdata/example.go:3:6: function "append" has same name as predeclared identifier
testdata/example.go:4:2: variable "copy" has same name as predeclared identifier
testdata/example.go:9:2: method "new" has same name as predeclared identifier
testdata/example.go:12:15: method "new" has same name as predeclared identifier
```
