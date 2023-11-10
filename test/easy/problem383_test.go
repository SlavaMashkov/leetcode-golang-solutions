package easy

import (
	"leetcode/problems/easy"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a <- b",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},
		{
			name: "aa <- ab",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		},
		{
			name: "aa <- aab",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.CanConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("CanConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
