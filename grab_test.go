package grab_test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/common-fate/grab"
	"github.com/stretchr/testify/assert"
)

func ExampleString() {
	foo := "bar"

	output := grab.If(foo == "bar", "foo is bar", "foo is not bar")

	fmt.Println(output)
	// Output: foo is bar
}

func TestIf(t *testing.T) {
	type args struct {
		condition bool
		ifTrue    string
		ifFalse   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				condition: true,
				ifTrue:    "true",
				ifFalse:   "false",
			},
			want: "true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grab.If(tt.args.condition, tt.args.ifTrue, tt.args.ifFalse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstNonZero(t *testing.T) {
	type args struct {
		elements []string
	}
	testString := []struct {
		name string
		args args
		want string
	}{
		{
			name: "second element",
			args: args{[]string{"", "selected", ""}},
			want: "selected",
		},
		{
			name: "no args returns zero value",
			args: args{},
			want: "",
		},
	}
	for _, tt := range testString {
		t.Run(tt.name, func(t *testing.T) {
			if got := grab.FirstNonZero(tt.args.elements...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstNonZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllPages(t *testing.T) {

	tests := []struct {
		name    string
		pages   [][]string
		want    []string
		mockErr error
		wantErr error
	}{
		{
			name:  "no pages",
			pages: [][]string{},
			want:  []string{},
		},
		{
			name:  "one page",
			pages: [][]string{{"a", "b"}},
			want:  []string{"a", "b"},
		},
		{
			name:  "two pages",
			pages: [][]string{{"a", "b"}, {"c", "d"}},
			want:  []string{"a", "b", "c", "d"},
		},
		{
			name:    "error",
			mockErr: errors.New("mock"),
			pages:   [][]string{},
			want:    []string{},
			wantErr: errors.New("mock"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := grab.AllPages(context.Background(), func(ctx context.Context, nextToken *int) ([]string, *int, error) {
				if tt.mockErr != nil {
					return nil, nil, tt.mockErr
				}
				next := grab.Value(nextToken)
				if len(tt.pages) == 0 {
					return nil, nil, nil
				}
				return tt.pages[next], grab.If(len(tt.pages)-1 == next, nil, grab.Ptr(next+1)), nil
			})
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestIsZero(t *testing.T) {
	type args[T comparable] struct {
		value T
	}
	tests := []struct {
		name string
		args args[string]
		want bool
	}{
		{
			name: "string is zero",
			args: args[string]{value: ""},
			want: true,
		},
		{
			name: "string is not zero",
			args: args[string]{value: "hello"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grab.IsZero(tt.args.value); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		fn    func(int) string
		want  []string
	}{
		{
			name:  "empty slice",
			items: []int{},
			fn:    func(i int) string { return fmt.Sprintf("Num: %d", i) },
			want:  []string{},
		},
		{
			name:  "single item",
			items: []int{1},
			fn:    func(i int) string { return fmt.Sprintf("Num: %d", i) },
			want:  []string{"Num: 1"},
		},
		{
			name:  "multiple items",
			items: []int{1, 2, 3},
			fn:    func(i int) string { return fmt.Sprintf("Num: %d", i) },
			want:  []string{"Num: 1", "Num: 2", "Num: 3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := grab.Map(tt.items, tt.fn)
			assert.ElementsMatch(t, tt.want, got)
			if got != nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
func TestFlatMap(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		fn    func(int) []string
		want  []string
	}{
		{
			name:  "empty slice",
			items: []int{},
			fn:    func(i int) []string { return []string{fmt.Sprintf("Num: %d", i)} },
			want:  []string{},
		},
		{
			name:  "single item",
			items: []int{1},
			fn:    func(i int) []string { return []string{fmt.Sprintf("Num: %d", i)} },
			want:  []string{"Num: 1"},
		},
		{
			name:  "multiple items",
			items: []int{1, 2, 3},
			fn:    func(i int) []string { return []string{fmt.Sprintf("Num: %d", i), fmt.Sprintf("Square: %d", i*i)} },
			want:  []string{"Num: 1", "Square: 1", "Num: 2", "Square: 4", "Num: 3", "Square: 9"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := grab.FlatMap(tt.items, tt.fn)
			assert.ElementsMatch(t, tt.want, got)
			if got != nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
func TestFilter(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		fn    func(int) bool
		want  []int
	}{
		{
			name:  "filter even numbers",
			items: []int{1, 2, 3, 4, 5},
			fn:    func(i int) bool { return i%2 == 0 },
			want:  []int{2, 4},
		},
		{
			name:  "all items filtered out",
			items: []int{1, 3, 5},
			fn:    func(i int) bool { return i%2 == 0 },
			want:  []int{},
		},
		{
			name:  "no items filtered out",
			items: []int{2, 4, 6},
			fn:    func(i int) bool { return i%2 == 0 },
			want:  []int{2, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := grab.Filter(tt.items, tt.fn)
			assert.ElementsMatch(t, tt.want, got)
			if got != nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMapFromSlice(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		value bool
		want  map[string]bool
	}{
		{
			name:  "all items mapped to true",
			items: []string{"apple", "banana", "orange"},
			value: true,
			want:  map[string]bool{"apple": true, "banana": true, "orange": true},
		},
		{
			name:  "all items mapped to false",
			items: []string{"apple", "banana", "orange"},
			value: false,
			want:  map[string]bool{"apple": false, "banana": false, "orange": false},
		},
		{
			name:  "empty slice",
			items: []string{},
			value: true,
			want:  map[string]bool{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := grab.MapFromSlice(tt.items, tt.value)

			// Check if the length of the map matches the expected length
			assert.Equal(t, len(tt.want), len(got), "Length of maps mismatch")

			// Check if each key-value pair in 'want' exists in 'got'
			for key, wantValue := range tt.want {
				gotValue, exists := got[key]
				assert.True(t, exists, fmt.Sprintf("Key %s does not exist in the result map", key))
				assert.Equal(t, wantValue, gotValue, fmt.Sprintf("Value for key %s mismatch", key))
			}

			// Check if each key in 'got' exists in 'want'
			for key := range got {
				_, exists := tt.want[key]
				assert.True(t, exists, fmt.Sprintf("Key %s should not exist in the result map", key))
			}
		})
	}
}
