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

## grab.IsZero

grab.IsZero checks if the provided value is the zero value for its type. It works with any comparable type.

```go
import "github.com/common-fate/grab"

zeroCheck := grab.IsZero[int](0) // zeroCheck will be true, as 0 is the zero value for int

nonZeroCheck := grab.IsZero[string]("hello") // nonZeroCheck will be false, as "hello" is not the zero value for string
```

This function is useful for determining if a value is uninitialized or set to its default state, particularly in generic programming where the type can vary.

## grab.AllPages

grab.AllPages aggregates all items from a paginated API into a single slice. It works with any type for the items and any comparable type for pagination tokens.

```go
import (
    "context"
    "github.com/common-fate/grab"
)

ctx := context.Background()
items, err := grab.AllPages[MyItem, string](ctx, myFetchPageFunc) // items will contain all MyItem instances from all pages fetched using myFetchPageFunc
```

This function abstracts away the pagination logic, allowing users to easily fetch and aggregate items from APIs that implement pagination. The user must provide a 'fetchPage' function that knows how to retrieve a single page of items and the next pagination token.

Created by @JoshuaWilkes.
