/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }
    
    slow := head
    fast := head
    for fast!=nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    var prev *ListNode
    cur := slow.Next
    slow.Next = nil
    for cur!=nil {
        temp := cur.Next
        cur.Next = prev

        prev = cur
        cur = temp
    }

    first, second := head, prev
    for second != nil {
        // Lưu lại các nút tiếp theo
        tmp1, tmp2 := first.Next, second.Next
        
        // Nối xen kẽ
        first.Next = second
        second.Next = tmp1
        
        // Di chuyển con trỏ đi tiếp
        first = tmp1
        second = tmp2
    }
}