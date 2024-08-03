package easy

import (
	"leetcode/problems/easy"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	type args struct {
		root *easy.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				getTree1(),
			},
			3,
		},
		{
			"test 2",
			args{
				root: getTree2(),
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.MaxDepth(tt.args.root); got != tt.want {
				t.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: create function that will create binary tree based on slice
func getTree1() *easy.TreeNode {
	n15 := &easy.TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}

	n7 := &easy.TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	n20 := &easy.TreeNode{
		Val:   20,
		Left:  n15,
		Right: n7,
	}

	n9 := &easy.TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}

	n3 := &easy.TreeNode{
		Val:   3,
		Left:  n9,
		Right: n20,
	}

	return n3
}

func getTree2() *easy.TreeNode {
	n2 := &easy.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	n1 := &easy.TreeNode{
		Val:   1,
		Left:  nil,
		Right: n2,
	}

	return n1
}
