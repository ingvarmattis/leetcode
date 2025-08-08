package longest_common_prefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "fl",
			args: args{strs: []string{
				"flower", "flow", "flight",
			}},
			want: "fl",
		},
		{
			name: "flower",
			args: args{strs: []string{
				"flowers", "flower", "flower",
			}},
			want: "flower",
		},
		{
			name: "фыа",
			args: args{strs: []string{
				"фыаацкй", "фыаощо", "фыа",
			}},
			want: "фыа",
		},
		{
			name: "no prefix",
			args: args{strs: []string{
				"фыаацкй", "awer",
			}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
