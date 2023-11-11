package easy

import (
	"leetcode/problems/easy"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "Case 2",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			name: "Case 3",
			args: args{
				strs: []string{"apple", "app", "application"},
			},
			want: "app",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.LongestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
