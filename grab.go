package grab

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
