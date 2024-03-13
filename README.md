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

`grab.IsZero` checks if the provided value is the zero value for its type. It works with any comparable type.

```go
import "github.com/common-fate/grab"

zeroCheck := grab.IsZero[int](0) // zeroCheck will be true, as 0 is the zero value for int

nonZeroCheck := grab.IsZero[string]("hello") // nonZeroCheck will be false, as "hello" is not the zero value for string
```

This function is useful for determining if a value is uninitialized or set to its default state, particularly in generic programming where the type can vary.

## grab.AllPages

`grab.AllPages` aggregates all items from a paginated API into a single slice. It works with any type for the items and any comparable type for pagination tokens.

```go
import (
    "context"
    "github.com/common-fate/grab"
)

ctx := context.Background()
items, err := grab.AllPages[MyItem, string](ctx, myFetchPageFunc) // items will contain all MyItem instances from all pages fetched using myFetchPageFunc
```

This function abstracts away the pagination logic, allowing users to easily fetch and aggregate items from APIs that implement pagination. The user must provide a 'fetchPage' function that knows how to retrieve a single page of items and the next pagination token.

## grab.Map

`grab.Map` applies a transformation function to each item in a slice, returning a slice of the results. It is a generic function that can operate on a slice of any type T and applies a function that transforms each T into another type F.

```go
import (
    "fmt"
    "github.com/common-fate/grab"
)

// Example usage of Map
items := []int{1, 2, 3}
transformedItems := grab.Map(items, func(i int) string {
    return fmt.Sprintf("Number: %d", i)
})
// transformedItems will be a []string: ["Number: 1", "Number: 2", "Number: 3"]
```

This function is particularly useful for cases where a slice of items needs to be transformed or mapped into a new slice of a different type, applying a specific operation or transformation to each item. The provided fn function encapsulates the logic of this transformation. It's a versatile tool for processing collections of data in a concise and readable manner.

## grab.Filter

`grab.Filter` filters elements of a slice based on a provided predicate function. It operates on a slice of any type `T` and returns a new slice containing only elements that satisfy the predicate.

```go
import (
    "fmt"
    "github.com/common-fate/grab"
)

// Example usage of Filter
numbers := []int{1, 2, 3, 4, 5}
evenNumbers := grab.Filter(numbers, func(n int) bool {
    return n%2 == 0
})
// evenNumbers will be a []int: [2, 4]
```

This function is useful when you need to extract elements from a collection based on specific criteria. The predicate function fn determines whether each item in the input slice should be included in the result. It's a practical tool for processing and manipulating slices.

## grab.MapFromSlice

`grab.MapFromSlice` creates a map from the given slice, where the elements of the slice become the keys and a provided value is associated with each key. It operates on a slice of any type T and returns a map with keys of type T and values of any type F.

```go
Copy code
import (
"fmt"
"github.com/common-fate/grab"
)

// Example usage of MapFromSlice
fruits := []string{"apple", "banana", "orange"}
fruitMap := grab.MapFromSlice(fruits, true)
// fruitMap will be a map[string]bool: {"apple": true, "banana": true, "orange": true}
```

This function is particularly useful when you need to quickly create a mapping of elements from a slice, associating each element with a common value. It simplifies the process of generating maps from slices and can be handy in various scenarios where key-value associations are required.

Created by @JoshuaWilkes.
