package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	mp := make(map[int]bool, len(nums))
	for _, i := range nums {
		mp[i] = true
	}

	new_head := &ListNode{Next: head}

	curr := new_head
	for curr.Next != nil {
		if mp[curr.Next.Val] {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return new_head.Next
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{nums[i], nil}
		curr = curr.Next
	}
	return head
}

func PrintList(head *ListNode) {
	temp := head
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}

func main() {
	list := []int{1, 2, 3, 4, 5}
	head := buildList(list)
	nums := []int{1, 2, 3}

	modifiedHead := modifiedList(nums, head)
	PrintList(modifiedHead)

}
