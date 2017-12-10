Find declarations and fields that have the same name as one of Go's predeclared
identifiers.

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

will print:

```
```
