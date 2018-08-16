## Go Optionals Values

A simple non-funky package with functions to represent Go's built-in value types to and from nil pointers.

    optionalString := goptionals.String("foo") // *string
    optionalString = nil
    nonOptional := goptionals.StringValue(optionalString) // ""

