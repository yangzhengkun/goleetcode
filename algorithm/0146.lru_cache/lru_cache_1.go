package _146_lru_cache

//自己实现双向链表

type LRUCache struct {
	size                 int
	capacity             int
	cache                map[int]*DLinkedNode
	dummyHead, dummyTail *DLinkedNode
}

type DLinkedNode struct {
	key   int
	value int
	prev  *DLinkedNode
	next  *DLinkedNode
}

func newNode(key, value int) *DLinkedNode {
	return &DLinkedNode{key: key, value: value}
}
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:      0,
		capacity:  capacity,
		cache:     make(map[int]*DLinkedNode),
		dummyHead: newNode(0, 0),
		dummyTail: newNode(0, 0),
	}
	l.dummyHead.next = l.dummyTail
	l.dummyTail.prev = l.dummyHead
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		// 缓存不命中
		return -1
	}
	// 缓存命中
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key, value int) {
	if _, ok := this.cache[key]; !ok {
		// 缓存不命中,则加入
		node := newNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		//缓存命中,将命中节点移到链表头部
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.dummyHead
	node.next = this.dummyHead.next
	this.dummyHead.next.prev = node
	this.dummyHead.next = node
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.dummyTail.prev
	this.removeNode(node)
	return node
}
