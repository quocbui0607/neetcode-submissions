type Node struct {
	Key int
	Val int
	Prev *Node
	Next *Node
}

type LRUCache struct {
    capacity int
	cache map[int]*Node
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
    head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		head: head,
		tail: tail,
		capacity: capacity,
		cache : make(map[int]*Node),
	}
}

func(this *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func(this *LRUCache) addToHead(node *Node) {
	node.Next = this.head.Next
	node.Prev = this.head
	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) Get(key int) int {
   if node, ok := this.cache[key]; ok {
		this.removeNode(node)
        this.addToHead(node)
		return node.Val
	} 

	return -1

}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.Val = value
		this.removeNode(node)
        this.addToHead(node)
	} else {
		if len(this.cache) >= this.capacity {
			lastNode := this.tail.Prev
			this.removeNode(lastNode)
			delete(this.cache, lastNode.Key)
		} 
		newNode := &Node{Key: key, Val: value}
		this.addToHead(newNode)
		this.cache[key] = newNode
	}
}
