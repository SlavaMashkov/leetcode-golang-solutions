package easy

// 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to serve as the head of the merged list
	dummy := &ListNode{}
	curr := dummy

	// Compare the values of the nodes in both lists and merge them in ascending order
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	// Attach the remaining nodes from either list
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return dummy.Next
}
