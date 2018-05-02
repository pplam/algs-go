package main

import "fmt"

func main() {
	array2List := func(arr []int) *ListNode {
		l := len(arr)
		if l == 0 {
			return nil
		}

		head := &ListNode{Val: arr[0]}
		if l == 1 {
			return head
		}

		tail := head
		for i := 1; i < l; i++ {
			tail.Next = &ListNode{Val: arr[i]}
			tail = tail.Next
		}

		return head
	}

	log := func(l *ListNode) {
		for p := l; p != nil; p = p.Next {
			fmt.Printf("%d -> ", p.Val)
		}
		fmt.Printf("nil\n")
	}

	l1 := array2List([]int{2, 4, 5})
	l2 := array2List([]int{5, 6, 4})
	sum := addTwoNumbers(l1, l2)

	fmt.Printf(" l1: ")
	log(l1)
	fmt.Printf(" l2: ")
	log(l2)
	fmt.Printf("sum: ")
	log(sum)
}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	sum2NodesWithCarry := func(n1, n2 *ListNode, carry int) (*ListNode, int) {
		v1, v2 := 0, 0
		if n1 != nil {
			v1 = n1.Val
		}
		if n2 != nil {
			v2 = n2.Val
		}
		sum := v1 + v2 + carry
		if sum < 10 {
			return &ListNode{Val: sum}, 0
		} else if sum < 20 {
			return &ListNode{Val: sum - 10}, 1
		} else {
			return &ListNode{Val: sum - 20}, 2
		}
	}

	head, carry := sum2NodesWithCarry(l1, l2, 0)
	tail := head

	p1, p2 := l1.Next, l2.Next
	for  p1 != nil || p2 != nil {
		tail.Next, carry = sum2NodesWithCarry(p1, p2, carry)
		tail = tail.Next
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}

	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return head
}
