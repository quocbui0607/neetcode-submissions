/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseLinkedListByK(head *ListNode, k int) *ListNode {
	var newHead *ListNode
	curr := head
	for curr != nil && k >0 {
		temp := curr.Next
		curr.Next = newHead
		newHead = curr
		curr = temp
		k--
	}

	return newHead
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    var newHead *ListNode
	var tailToExtend *ListNode
	curr := head
	for curr!=nil {
		count := 0
		newTemp := curr
		for count < k && newTemp != nil {
			count += 1
			newTemp = newTemp.Next
		}

		if count == k {
			reverseList := reverseLinkedListByK(curr, k)
			if newHead == nil {
				newHead = reverseList
			}
			
			if tailToExtend != nil {
				tailToExtend.Next = reverseList
			}
			tailToExtend = curr
			curr = newTemp
		} else {
			tailToExtend.Next = curr
			break
		}
	}

	return newHead
}
