package palindrome_number

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "121=true",
			args: args{121},
			want: true,
		},
		{
			name: "1121=false",
			args: args{1121},
			want: false,
		},
		{
			name: "1=true",
			args: args{1},
			want: true,
		},
		{
			name: "-121=false",
			args: args{-121},
			want: false,
		},
		{
			name: "-1=false",
			args: args{-1},
			want: false,
		},
		{
			name: "0=true",
			args: args{0},
			want: true,
		},
		{
			name: "11=true",
			args: args{11},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
