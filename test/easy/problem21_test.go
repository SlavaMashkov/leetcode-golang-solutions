package easy

import (
	"leetcode/problems/easy"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		list1 *easy.ListNode
		list2 *easy.ListNode
	}
	tests := []struct {
		name string
		args args
		want *easy.ListNode
	}{
		{
			name: "Test Case 1",
			args: args{
				list1: &easy.ListNode{Val: 1, Next: &easy.ListNode{Val: 2, Next: &easy.ListNode{Val: 4}}},
				list2: &easy.ListNode{Val: 1, Next: &easy.ListNode{Val: 3, Next: &easy.ListNode{Val: 4}}},
			},
			want: &easy.ListNode{Val: 1, Next: &easy.ListNode{Val: 1, Next: &easy.ListNode{Val: 2, Next: &easy.ListNode{Val: 3, Next: &easy.ListNode{Val: 4, Next: &easy.ListNode{Val: 4}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.MergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
