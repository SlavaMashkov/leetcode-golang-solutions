package easy

import (
	"leetcode/problems/easy"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test 1",
			args{
				numRows: 5,
			},
			[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			"test 2",
			args{
				numRows: 1,
			},
			[][]int{
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("want: %v", tt.want)

			if got := easy.Generate(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
