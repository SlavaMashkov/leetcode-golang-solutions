package easy

import (
	"leetcode/problems/easy"
	"testing"
)

func Test_maximumWealth(t *testing.T) {
	type args struct {
		accounts [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty accounts",
			args: args{
				accounts: [][]int{},
			},
			want: 0,
		},
		{
			name: "Single account with zero wealth",
			args: args{
				accounts: [][]int{
					{0},
				},
			},
			want: 0,
		},
		{
			name: "Multiple accounts with varying wealth 1",
			args: args{
				accounts: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: 24,
		},
		{
			name: "Multiple accounts with varying wealth 2",
			args: args{
				accounts: [][]int{
					{1, 2, 3},
					{3, 2, 1},
				},
			},
			want: 6,
		},
		{
			name: "Multiple accounts with varying wealth 3",
			args: args{
				accounts: [][]int{
					{1, 5},
					{7, 3},
					{3, 5},
				},
			},
			want: 10,
		},
		{
			name: "Multiple accounts with varying wealth 4",
			args: args{
				accounts: [][]int{
					{2, 8, 7},
					{7, 1, 3},
					{1, 9, 5},
				},
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.MaximumWealth(tt.args.accounts); got != tt.want {
				t.Errorf("maximumWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
