package easy

import (
	"leetcode/problems/easy"
	"testing"
)

func TestNumberOfSteps(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "num = 14",
			args: args{num: 14},
			want: 6,
		},
		{
			name: "num = 8",
			args: args{num: 8},
			want: 4,
		},
		{
			name: "num = 6",
			args: args{num: 6},
			want: 4,
		},
		// Additional test cases
		{
			name: "num = 0",
			args: args{num: 0},
			want: 0,
		},
		{
			name: "num = 1",
			args: args{num: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.NumberOfSteps(tt.args.num); got != tt.want {
				t.Errorf("NumberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
