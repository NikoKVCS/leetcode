// +build ignore

package main

func main() {
	a := Constructor(1)
	a.Put(2, 2)
	a.Get(2)
	a.Put(3, 3)
	a.Get(2)
	a.Get(3)
}

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	capacity int
	length   int
	hash     map[int]*Node
	head     *Node
	rear     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{0, 0, nil, nil}
	rear := &Node{0, 0, head, nil}
	head.next = rear
	return LRUCache{capacity, 0, make(map[int]*Node, 0), head, rear}
}

func (this *LRUCache) MoveToHead(node *Node) {
	A := node.prev
	B := node.next
	A.next = B
	B.prev = A
	this.InsertToHead(node)
}

func (this *LRUCache) InsertToHead(node *Node) {
	this.head.next.prev = node
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.hash[key]; ok {
		this.MoveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) DeleteRear() {
	this.rear.prev.prev.next = this.rear
	this.rear.prev = this.rear.prev.prev
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.hash[key]; ok {
		this.MoveToHead(node)
		node.val = value
	} else if this.length >= this.capacity {
		delete(this.hash, this.rear.prev.key)
		this.DeleteRear()
		node := &Node{key, value, nil, nil}
		this.InsertToHead(node)
		this.hash[key] = node
	} else {
		node := &Node{key, value, nil, nil}
		this.InsertToHead(node)
		this.hash[key] = node
		this.length++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
