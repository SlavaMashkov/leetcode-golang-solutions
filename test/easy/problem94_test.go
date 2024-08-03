package easy

import (
	"leetcode/problems/easy"
	"leetcode/problems/structures"
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	type args struct {
		root *structures.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				root: &structures.TreeNode{
					Val:  1,
					Left: nil,
					Right: &structures.TreeNode{
						Val: 2,
						Left: &structures.TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "Test Case 2",
			args: args{
				root: &structures.TreeNode{
					Val: 1,
					Left: &structures.TreeNode{
						Val: 2,
						Left: &structures.TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
						Right: &structures.TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &structures.TreeNode{
						Val: 3,
						Left: &structures.TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: &structures.TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name: "Test Case 3",
			args: args{
				root: &structures.TreeNode{
					Val: 9,
					Left: &structures.TreeNode{
						Val: 3,
						Left: &structures.TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
						Right: &structures.TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &structures.TreeNode{
						Val: 7,
						Left: &structures.TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: &structures.TreeNode{
							Val:   8,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: []int{5, 3, 2, 9, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.InorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
