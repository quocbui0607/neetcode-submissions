/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
		return nil
	}

	hashMap := make(map[*Node]*Node)
	curr := head
	for curr != nil {
		hashMap[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		hashMap[curr].Next = hashMap[curr.Next]
		hashMap[curr].Random = hashMap[curr.Random]
		curr = curr.Next
	}

	return hashMap[head]
}
