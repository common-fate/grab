# grab

Go functions to make working with pointers and ternaries more concise.

## grab.If

`grab.If` acts as a ternary operator.

```go
import "github.com/common-fate/grab"

foo := "bar"

output := grab.If(foo == "bar", "foo is bar", "foo is not bar")

// Output is "foo is bar"
```

## grab.Ptr

`grab.Ptr` returns the pointer to a value.

```go
import "github.com/common-fate/grab"

output := grab.Ptr("foo")

// Output is a pointer to a string "foo"
```

This is useful when constructing structs that take nillable fields, for example:

```go
type MyStruct struct {
    OptionalString *string
}

output := MyStruct{
    OptionalString: grab.Ptr("foo")
}
```

## grab.Value

`grab.Value` takes a pointer and returns the underlying value. If the pointer is nil, it returns the zero value of the type.

```go
import "github.com/common-fate/grab"

var intPointer *int
intValue := grab.Value(intPointer) // intValue will be 0 if intPointer is nil
```

This function is particularly useful for safely dereferencing pointers, especially when there's a possibility of them being nil.

## grab.FirstNonZero

`grab.FirstNonZero` takes a variadic list of comparable elements and returns the first non zero value comparing from start to finish of the supplied elements.

```go
import (
    "os"
    "github.com/common-fate/grab"
)

var userSuppliedPath string

path := grab.FirstNonZero(userSuppliedPath, os.Getenv("PATH"),"/default/path") // path will be the first non empty string
```

This function is particularly useful for setting configuration by precedence from a range of available sources.

Created by @JoshuaWilkes.
