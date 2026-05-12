/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    currNode := head
    for currNode != nil {
        temp := currNode.Next
        currNode.Next = prev
        prev = currNode
        currNode = temp
    }

    return prev
}
