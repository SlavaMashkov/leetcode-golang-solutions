package test

import (
	"leetcode/code-run/easy"
	"testing"
)

func TestIsInputPossible(t *testing.T) {
	tests := []struct {
		name string
		set  map[int]bool
		arr  []int
		want int
	}{
		{
			"test 1",
			map[int]bool{
				1: true,
				2: true,
				3: true,
			},
			[]int{
				1, 1, 2, 3,
			},
			0,
		},
		{
			"test 2",
			map[int]bool{
				1: true,
				2: true,
				3: true,
			},
			[]int{
				1, 0, 0, 1,
			},
			1,
		},
		{
			"test 3",
			map[int]bool{
				5: true,
				7: true,
				3: true,
			},
			[]int{
				1, 2, 3,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.IsInputPossible(tt.set, tt.arr); got != tt.want {
				t.Errorf("IsInputPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
