// Package goptionals provides simple non-funky functions to represent Go's built-in value types to and from nil pointers.
//
//    optionalString := goptionals.String("foo") // *string
//    optionalString = nil
//    nonOptional := goptionals.StringValue(optionalString) // ""
//
// Supports string, bool, float64, int, int64 and time.Time
package goptionals
