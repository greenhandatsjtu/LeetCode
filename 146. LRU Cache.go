package main

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

//use a map and double linkedList
type LRUCache struct {
	capacity int
	head     *node
	tail     *node
	mem      map[int]*node
}

func Constructor(capacity int) LRUCache {
	var head, tail *node
	head = &node{prev: nil, next: nil}
	tail = &node{prev: head, next: head}
	head.prev = tail
	head.next = tail

	mem := make(map[int]*node, capacity)

	return LRUCache{
		capacity: capacity,
		mem:      mem,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) push(key int, value int) {
	//add to linkedList
	newNode := &node{key: key, value: value, prev: this.head, next: this.head.next}
	this.head.next = newNode
	newNode.next.prev = newNode

	//add to map
	this.mem[key] = newNode
}

func (this *LRUCache) pop() {
	deleteNode := this.tail.prev

	//remove from linkedList
	deleteNode.prev.next = this.tail
	this.tail.prev = deleteNode.prev

	//remove from map
	delete(this.mem, deleteNode.key)
}

func (this *LRUCache) moveToFront(n *node) {
	next := n.next
	n.prev.next = next
	next.prev = n.prev

	n.next = this.head.next
	n.prev = this.head
	this.head.next = n
	n.next.prev = n
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.mem[key]; !ok {
		return -1
	} else {
		n := this.mem[key]
		this.moveToFront(n)
		return n.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.mem[key]; !ok {
		if len(this.mem) == this.capacity {
			this.pop()
		}
		this.push(key, value)
	} else {
		n := this.mem[key]
		n.value = value
		this.moveToFront(n)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
