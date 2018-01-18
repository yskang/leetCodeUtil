package leetData

import (
	"strconv"
	"strings"
)

// ListNode has value and the pointer of next ListNode.
type ListNode struct {
	Val  int
	Next *ListNode
}

// InitListNode make a ListNode from list of ints and return a pointer of the ListNode.
func InitListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := ListNode{nums[0], nil}
	tempPointer := &head

	for i := 1; i < len(nums); i++ {
		tempPointer.Next = &ListNode{nums[i], nil}
		tempPointer = tempPointer.Next
	}

	return &head
}

func (l *ListNode) String() string {
	lists := make([]string, 0)

	current := l
	for current != nil {
		lists = append(lists, strconv.Itoa(current.Val))
		current = current.Next
	}

	return strings.Join(lists, ",")
}
