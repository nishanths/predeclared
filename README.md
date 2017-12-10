# predeclared [![Build Status](https://travis-ci.org/nishanths/predeclared.svg?branch=master)](https://travis-ci.org/nishanths/predeclared)

Find declarations that override one of Go's predeclared identifiers (`new`, `make`, `append`, etc.).

The list of predeclared identifiers can be found in the [spec](https://golang.org/ref/spec#Predeclared_identifiers).

## Usage

```
go get -u github.com/nishanths/predeclared

predeclared dir1 dir2 file.go
```

See [godoc](https://godoc.org/github.com/nishanths/predeclared) or `predeclared -h` for more.

## Example

Given a file:

```
package print

func make(i *int) T {
	copy := *i
	return T{copy}
}
```

running:

```
predeclared example.go
```

prints:

```
example.go:1:9: package name "print" has same name as predeclared identifier
example.go:3:6: function "make" has same name as predeclared identifier
example.go:4:2: variable "copy" has same name as predeclared identifier
```
