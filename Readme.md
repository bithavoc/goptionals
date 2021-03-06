## Go Optionals Values
[![CircleCI](https://circleci.com/gh/bithavoc/goptionals.svg?style=svg&circle-token=6f78d762987052548e37de3da948df0aeb49548b)](https://circleci.com/gh/bithavoc/goptionals)
[![GoCover](https://gocover.io/_badge/github.com/bithavoc/goptionals)](https://gocover.io/github.com/bithavoc/goptionals)
[![Go Report Card](https://goreportcard.com/badge/github.com/bithavoc/goptionals)](https://goreportcard.com/report/github.com/bithavoc/goptionals)
[![GoDoc](https://godoc.org/github.com/bithavoc/goptionals?status.svg)](https://godoc.org/github.com/bithavoc/goptionals)

> A simple non-funky package with functions to represent Go's built-in value types to and from nil pointers.

Example:

    optionalString := goptionals.String("foo") // *string
    optionalString = nil
    nonOptional := goptionals.StringValue(optionalString) // ""


## Test

```
go test
```

with coverage:

```
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```
