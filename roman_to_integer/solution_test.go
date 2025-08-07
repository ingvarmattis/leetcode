package roman_to_integer

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "III=3",
			args: args{"III"},
			want: 3,
		},
		{
			name: "IIII=4",
			args: args{"IIII"},
			want: 4,
		},
		{
			name: "IV=4",
			args: args{"IV"},
			want: 4,
		},
		{
			name: "VI=6",
			args: args{"VI"},
			want: 6,
		},
		{
			name: "IMCDMI=1600",
			args: args{"MCDM"},
			want: 1600,
		},
		{
			name: "I=1",
			args: args{"I"},
			want: 1,
		},
		{
			name: "empty string=0",
			args: args{""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
