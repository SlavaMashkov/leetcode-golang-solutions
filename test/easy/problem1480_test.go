package easy

import (
	"leetcode/problems/easy"
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{nums: []int{1, 2, 3, 4}},
			want: []int{1, 3, 6, 10},
		},
		{
			name: "Test Case 2",
			args: args{nums: []int{1, 1, 1, 1}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Test Case 3",
			args: args{nums: []int{3, 1, 2, 10, 1}},
			want: []int{3, 4, 6, 16, 17},
		},
		{
			name: "Test Case 4",
			args: args{nums: []int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.RunningSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunningSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
