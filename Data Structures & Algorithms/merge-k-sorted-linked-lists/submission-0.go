/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
     cur := dummy
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
           cur.Next = list1 
           list1 = list1.Next
        } else {
           cur.Next = list2 
           list2 = list2.Next
        }
        
        cur = cur.Next
    }

    if list1 != nil {
        cur.Next = list1
    }

    if list2 != nil {
        cur.Next = list2
    }

    return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
     if len(lists) == 0 {
        return nil
    }

    for i := 1; i < len(lists); i++ {
        lists[i] = mergeTwoLists(lists[i-1], lists[i])
    }

    return lists[len(lists)-1]
}
