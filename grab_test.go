package grab_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/common-fate/grab"
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
