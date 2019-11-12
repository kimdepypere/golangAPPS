package tapps

import (
	"testing"
)

func Test_hello(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty value",
			args: args{message: ""},
			want: "Nothing to declare",
		}, {
			name: "value filled in ('testvalue')",
			args: args{message: "testvalue"},
			want: "the message was: testvalue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hello(tt.args.message); got != tt.want {
				t.Errorf("hello("+tt.args.message+") = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_goodbye(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Goodbye test",
			want: "goodbye",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodbye(); got != tt.want {
				t.Errorf("goodbye() = %v, want %v", got, tt.want)
			}
		})
	}
}
