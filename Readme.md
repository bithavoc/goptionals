## Go Optionals Values
[![Tests Status](https://github.com/bithavoc/goptionals/actions/workflows/test.yml/badge.svg)](https://github.com/bithavoc/goptionals/blob/master/.github/workflows/test.yml)
[![GoCover](https://gocover.io/_badge/github.com/bithavoc/goptionals)](https://gocover.io/github.com/bithavoc/goptionals)
[![Go Report Card](https://goreportcard.com/badge/github.com/bithavoc/goptionals)](https://goreportcard.com/report/github.com/bithavoc/goptionals)
[![Go Reference](https://pkg.go.dev/badge/github.com/bithavoc/goptionals.svg)](https://pkg.go.dev/github.com/bithavoc/goptionals)

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
