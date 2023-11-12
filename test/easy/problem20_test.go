package easy

import (
	"leetcode/problems/easy"
	"testing"
)

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				s: "{[]}",
			},
			want: true,
		},
		{
			name: "Test Case 6",
			args: args{
				s: "([}}])",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.IsValid(tt.args.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
