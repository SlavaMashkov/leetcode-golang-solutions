package test

import (
	"leetcode/code-run/easy"
	"log/slog"
	"testing"
)

func Test_isTriangle(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected string
	}{
		{
			"test 1",
			[]int{
				1, 2, 3,
			},
			"NO",
		},
		{
			"test 2",
			[]int{
				3, 4, 5,
			},
			"YES",
		},
		{
			"test 3",
			[]int{
				3, 5, 4,
			},
			"YES",
		},
		{
			"test 4",
			[]int{
				4, 5, 3,
			},
			"YES",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := easy.IsTriangle(tt.arr)

			slog.Info(tt.name, slog.String("result", result), slog.String("expected", tt.expected))

			if result != tt.expected {
				t.Fail()
			}
		})
	}
}
