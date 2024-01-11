package solutions

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example1", args: args{s: "42"}, want: 42},
		{name: "example2", args: args{s: "   -42"}, want: -42},
		{name: "example3", args: args{s: "4193 with words"}, want: 4193},
		{name: "example4", args: args{s: "20000000000000000000"}, want: 2147483647},
		{name: "example5", args: args{s: "-20000000000000000000"}, want: -2147483648},
		{name: "example6", args: args{s: "Word of 987"}, want: 0},
		{name: "example7", args: args{s: "    -88827   5655  U"}, want: -88827},
		{name: "example8", args: args{s: "-5-"}, want: -5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
