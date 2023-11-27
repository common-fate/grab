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
