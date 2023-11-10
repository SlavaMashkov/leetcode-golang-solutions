package easy

import (
	"leetcode/problems/easy"
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	type args struct {
		head *easy.ListNode
	}
	tests := []struct {
		name string
		args args
		want *easy.ListNode
	}{
		{
			name: "Null",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "1 node",
			args: args{
				head: &easy.ListNode{Val: 1, Next: nil},
			},
			want: &easy.ListNode{Val: 1, Next: nil},
		},
		{
			name: "2 nodes",
			args: args{
				head: &easy.ListNode{Val: 1, Next: &easy.ListNode{Val: 2, Next: nil}},
			},
			want: &easy.ListNode{Val: 2, Next: nil},
		},
		{
			name: "5 nodes",
			args: args{
				head: &easy.ListNode{Val: 1, Next: &easy.ListNode{Val: 2, Next: &easy.ListNode{Val: 3, Next: &easy.ListNode{Val: 4, Next: &easy.ListNode{Val: 5, Next: nil}}}}},
			},
			want: &easy.ListNode{Val: 3, Next: &easy.ListNode{Val: 4, Next: &easy.ListNode{Val: 5, Next: nil}}},
		},
		{
			name: "6 nodes",
			args: args{
				head: &easy.ListNode{Val: 1, Next: &easy.ListNode{Val: 2, Next: &easy.ListNode{Val: 3, Next: &easy.ListNode{Val: 4, Next: &easy.ListNode{Val: 5, Next: &easy.ListNode{Val: 6, Next: nil}}}}}},
			},
			want: &easy.ListNode{Val: 4, Next: &easy.ListNode{Val: 5, Next: &easy.ListNode{Val: 6, Next: nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.MiddleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MiddleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
