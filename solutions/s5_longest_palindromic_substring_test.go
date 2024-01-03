package solutions

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{s: "babad"},
			want: "aba",
		},
		{
			name: "example2",
			args: args{s: "cbbd"},
			want: "bb",
		},
		{
			name: "example3",
			args: args{s: "bb"},
			want: "bb",
		},
		{
			name: "example4",
			args: args{s: "aacabdkacaa"},
			want: "aca",
		},
		{
			name: "example5",
			args: args{s: "ssssss"},
			want: "ssssss",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
