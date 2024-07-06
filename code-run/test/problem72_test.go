package test

import (
	"leetcode/code-run/easy"
	"testing"
)

func Test_isListAscending(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected string
	}{
		{
			"test1",
			[]int{1, 7, 9},
			"YES",
		},
		{
			"test2",
			[]int{1, 9, 7},
			"NO",
		},
		{
			"test3",
			[]int{2, 2, 2},
			"NO",
		},
		{
			"test4",
			[]int{-1, 0, 1},
			"YES",
		},
		{
			"test5",
			[]int{1, 0, -1},
			"NO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := easy.IsListAscending(tt.arr); actual != tt.expected {
				t.Errorf("isListAscending() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}
