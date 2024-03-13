package grab

import "context"

// Ptr takes any value of type 'T' and returns a pointer to a new copy of that value.
// It is a generic function that can handle any type.
//
// Parameter:
// - o: The value of type 'T' to be converted into a pointer.
//
// Returns:
// - *T: A pointer to a new copy of the value 'o'.
//
// Example:
// intValue := 42
// intPointer := Ptr(intValue) // intPointer now points to a new copy of intValue
//
// Note: This function is useful for creating pointers for values returned by functions
// that do not directly return a pointer, such as time.Now() in the time package.
func Ptr[T any](o T) *T {
	return &o
}

// Value takes a pointer of type 'T' and returns the underlying value.
// If the pointer is nil, it returns the zero value of type 'T'.
// It is a generic function that can handle any type.
//
// Parameter:
// - o: A pointer of type 'T'. Can be nil.
//
// Returns:
// - T: The value pointed to by 'o', or the zero value of type 'T' if 'o' is nil.
//
// Example:
// var intPointer *int
// intValue := Value(intPointer) // intValue will be 0 if intPointer is nil
//
// Note: This function is particularly useful for safely dereferencing pointers,
// especially when there's a possibility of them being nil.
func Value[T any](o *T) T {
	if o == nil {
		var out T
		return out
	}
	return *o
}

// If evaluates a boolean condition and returns one of two values based on the result.
// It is a generic function that works with any type (denoted by 'T').
//
// Parameters:
// - condition: A boolean expression that determines which value is returned.
// - ifTrue: The value returned if 'condition' is true.
// - ifFalse: The value returned if 'condition' is false.
//
// Returns:
// - T: The value of 'ifTrue' if 'condition' is true, or the value of 'ifFalse' otherwise.
//
// Example:
// result := If[bool](someCondition, trueValue, falseValue)
//
// Note: This function is particularly useful for concise inline conditional assignments and
// can replace traditional if-else statements in such scenarios.
func If[T any](condition bool, ifTrue, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// FirstNonZero returns the first non-zero element from a given list of elements.
// It is a generic function that works with any comparable type (denoted by 'T').
//
// Parameters:
//   - elements: A variadic list of elements of type 'T'. The function iterates through these elements
//     in the order they are provided and returns the first one that is not equal to the zero value of type 'T'.
//
// Returns:
//   - T: The first non-zero element in the 'elements' list. If all elements are zero or the list is empty,
//     the zero value of type 'T' is returned.
//
// Example:
// result := FirstNonZero[int](0, 3, 0, 4)
// // result will be 3, as it's the first non-zero element in the list
//
// Note: This function is useful for selecting the first significant element from a list of ordered values,
// where 'significant' means not being the zero value of the specified type. It simplifies the process of
// checking multiple values in precedence order.
func FirstNonZero[T comparable](elements ...T) T {
	var zero T
	for _, elem := range elements {
		if elem != zero {
			return elem
		}
	}
	return zero
}

// IsZero checks if the provided value is the zero value for its type.
// It is a generic function that works with any comparable type (denoted by 'T').
//
// Parameters:
//   - value: The value of type 'T' to be checked against its zero value.
//
// Returns:
//   - bool: Returns true if 'value' is the zero value for its type; otherwise, returns false.
//
// Example:
// zeroCheck := IsZero[int](0)
// // zeroCheck will be true, as 0 is the zero value for int
//
// Note: This function is useful for determining if a value is uninitialized or set to its default state.
// It is particularly handy in generic programming where the type 'T' can vary and direct comparison
// to a known zero value is not possible.
func IsZero[T comparable](value T) bool {
	var zero T
	return value == zero
}

// AllPages aggregates all items from a paginated API into a single slice.
// It is a generic function that works with any type 'T' for the items and any comparable type 'Token' for pagination tokens.
//
// Parameters:
//   - ctx: A context.Context used for cancellation and timeout control of the HTTP requests.
//   - fetchPage: A function that retrieves a single page. It takes the current pagination token (of type 'Token')
//     and returns a slice of items (of type 'T'), the next pagination token (of type 'Token'), and an error if any occurs.
//
// Returns:
//   - []T: A slice containing all aggregated items from all pages.
//   - error: An error if any occurs during the fetching of pages. If an error is returned, the slice may contain
//     the items fetched before the error occurred.
//
// Example:
// items, err := AllPages[MyItem, string](ctx, myFetchPageFunc)
// // items will contain all MyItem instances from all pages fetched using myFetchPageFunc
//
// Note: This function abstracts away the pagination logic, allowing users to easily fetch and aggregate
// items from APIs that implement pagination. The user must provide a 'fetchPage' function that knows
// how to retrieve a single page of items and the next pagination token.
func AllPages[T any, Token comparable](ctx context.Context, fetchPage func(ctx context.Context, nextToken *Token) ([]T, *Token, error)) ([]T, error) {
	var allItems []T
	var nextToken *Token

	for {
		items, newToken, err := fetchPage(ctx, nextToken)
		if err != nil {
			return nil, err
		}

		allItems = append(allItems, items...)

		if newToken == nil || IsZero(*newToken) {
			break
		}
		nextToken = newToken
	}

	return allItems, nil
}

// Map applies a transformation function to each item in a slice and returns a slice of the results.
// It is a generic function that operates on a slice of any type 'T' and applies a function that transforms each 'T' into another type 'F'.
//
// Parameters:
//   - items: A slice of items of type 'T'. These are the items to be transformed.
//   - fn: A function that takes an item of type 'T' and returns a new item of type 'F'. This function defines the transformation to be applied to each item.
//
// Returns:
//   - []F: A slice containing all the transformed items. The order of the items in this slice corresponds to the order of the items in the input slice.
//
// Example:
// originalItems := []int{1, 2, 3}
//
//	transformedItems := Map[int, string](originalItems, func(i int) string {
//	    return fmt.Sprintf("Number: %d", i)
//	})
//
// // transformedItems will be a []string with values: ["Number: 1", "Number: 2", "Number: 3"]
//
// Note: This function is useful for cases where a slice of items needs to be transformed or mapped into a new slice of a different type,
// applying a specific operation or transformation to each item. The provided 'fn' function encapsulates the logic of this transformation.
func Map[T any, F any](items []T, fn func(T) F) []F {
	var result []F
	for _, item := range items {
		result = append(result, fn(item))
	}
	return result
}

// Filter iterates over elements of a slice, returning a new slice of all elements for which the predicate `fn` returns true.
//
// Parameters:
//   - items: A slice of items of type 'T'. These are the items to be filtered.
//   - fn: A predicate function that takes an item of type 'T' and returns a bool. If 'fn' returns true, the item is included in the result.
//
// Returns:
//   - []T: A slice containing all items that satisfy the predicate 'fn'.
//
// Example:
// numbers := []int{1, 2, 3, 4, 5}
//
//	evenNumbers := Filter(numbers, func(n int) bool {
//	    return n%2 == 0
//	})
//
// // evenNumbers will be a []int: [2, 4]
func Filter[T any](items []T, fn func(T) bool) []T {
	var result []T
	for _, item := range items {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

// MapFromSlice creates a map from the given slice where the elements of the slice are the keys and the value is a generic type.
// The value for each key is set to the provided 'value'.
//
// Parameters:
//   - items: A slice of items of type 'T'. These are the keys for the map.
//   - value: The value to be associated with each key in the map.
//
// Returns:
//   - map[T]F: A map where the keys are elements from the input slice and the value for each key is set to the provided 'value'.
//
// Example:
// originalSlice := []string{"apple", "banana", "orange"}
//
// newMap := MapFromSlice(originalSlice, true)
//
// // newMap will be a map[string]bool with keys {"apple": true, "banana": true, "orange": true}
func MapFromSlice[T comparable, F any](items []T, value F) map[T]F {
	result := make(map[T]F, len(items))
	for _, item := range items {
		result[item] = value
	}
	return result
}
