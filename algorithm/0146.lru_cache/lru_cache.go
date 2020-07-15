package _146_lru_cache

import "container/list"

//运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
// example
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得关键字 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得关键字 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4

//使用标准库自带的list
type LRUCache struct {
	capacity int
	data     map[int]*list.Element
	list     *list.List
}

type Node struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, list: list.New(), data: map[int]*list.Element{}}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.data[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(Node).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.data[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value = Node{key: key, val: value}
	} else {
		if this.list.Len() >= this.capacity {
			// 需要淘汰最近最少被访问的node(链表的尾节点)
			delete(this.data, this.list.Back().Value.(Node).key) // 删除map中的数据
			this.list.Remove(this.list.Back())                   // 删除链表的尾结点
		}
		// 加入新数据
		this.list.PushFront(Node{key: key, val: value})
		this.data[key] = this.list.Front()
	}
}
