package test

import (
	"leetcode/code-run/easy"
	"testing"
)

func TestNearestNumber(t *testing.T) {
	type args struct {
		arr []int
		x   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]int{
					1, 2, 3, 4, 5,
				},
				6,
			},
			5,
		},
		{
			"test 2",
			args{
				[]int{
					5, 4, 3, 2, 1,
				},
				3,
			},
			3,
		},
		{
			"test 3",
			args{
				[]int{
					-2, -1, 0, 1, 2,
				},
				1,
			},
			1,
		},
		{
			"test 4",
			args{
				[]int{
					1, 2, 4, 6, 8,
				},
				7,
			},
			8,
		},
		{
			"test 5",
			args{
				[]int{
					1, 2, 3, 4, 5,
				},
				-2,
			},
			1,
		},
		{
			"test 6",
			args{
				[]int{
					-10, -6, 1, 2, 3,
				},
				-8,
			},
			-6,
		},
		{
			"test 7",
			args{
				[]int{
					1000, 900, 800, 700, 600,
				},
				1000,
			},
			1000,
		},
		{
			"test 8",
			args{
				[]int{
					600, 700, 800, 900, 1000,
				},
				-600,
			},
			600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.NearestNumber(tt.args.arr, tt.args.x); got != tt.want {
				t.Errorf("NearestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
