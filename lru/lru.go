package lru

import "fmt"

type node struct {
	value any
	prev  *node
	next  *node
}

type linkedList struct {
	head   *node
	length int
	cap    int
}

func newNode(value any) *node {
	return &node{
		value: value,
	}
}

func NewLru(cap int) *linkedList {
	return &linkedList{
		length: 0,
		cap:    cap,
	}
}

func (l *linkedList) set(value any) bool {
	n := newNode(value)

	// 存在历史数据,移动到头部
	if node := l.find(value); node != nil {
		return l.move(node)
	}

	// 判断容量满了,删除结尾的节点
	if l.length == l.cap {
		l.deleteLast()
	}

	if l.head == nil {
		l.head = n
	} else {
		temp := l.head
		l.head = n
		temp.prev = n
		l.head.next = temp
	}
	l.length++
	return true
}

func (l *linkedList) move(node *node) bool {
	l.delete(node)
	cur := l.head
	cur.prev = node
	node.next = cur
	l.head = node
	l.head.prev = nil
	l.length++
	return true
}

func (l *linkedList) find(value any) *node {
	if l.head == nil {
		return nil
	}
	cur := l.head
	for cur != nil {
		if cur.value == value {
			break
		}
		cur = cur.next
	}
	return cur
}

func (l *linkedList) findByIndex(i int) *node {
	if l.head == nil {
		return nil
	}
	if l.length-1 < i {
		return nil
	}
	cur := l.head
	for p := 0; p != i; p++ {
		cur = cur.next
	}
	return cur
}

func (l *linkedList) deleteLast() bool {
	node := l.findByIndex(l.length - 1)
	return l.delete(node)
}

func (l *linkedList) delete(n *node) bool {
	if l.head == nil {
		return false
	}
	cur := l.head
	for cur != nil {
		if cur == n {
			if cur.prev == nil && cur.next == nil {
				l.head = nil
			} else if cur.prev == nil {
				l.head = cur.next
				l.head.prev = nil
			} else {
				temp := cur.prev
				temp.next = cur.next
				if cur.next != nil {
					cur.next.prev = temp
				}
			}
			l.length--
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *linkedList) print() {
	if l.head == nil {
		return
	}
	temp := l.head
	for temp != nil {
		fmt.Printf("%v", temp.value)
		temp = temp.next
		if temp != nil {
			fmt.Print("->")
		}
	}
	fmt.Printf(":%d", l.length)
	fmt.Println()
	temp = l.head
	for temp != nil {
		if temp.prev != nil {
			fmt.Printf("%v", temp.prev.value)
			fmt.Print("<-")
		}
		if temp.next == nil {
			fmt.Printf("%v", temp.value)
		}
		temp = temp.next
	}
	fmt.Printf(":%d", l.length)
	fmt.Println()
}
