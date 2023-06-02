// Package singlelinkedlist 单链操作
package singlelinkedlist

import "fmt"

type Node[T any] struct {
	next  *Node[T]
	value T
}

type NodeList[T any] struct {
	head   *Node[T]
	length int
}

func newNode[T any](v T) *Node[T] {
	return &Node[T]{
		value: v,
	}
}

func NewNodeList[T any]() *NodeList[T] {
	return &NodeList[T]{
		head: &Node[T]{},
	}
}

// 在某个节点之后插入节点
func (n *NodeList[T]) addNodeAfter(p *Node[T], value T) bool {
	if p == nil {
		return false
	}
	tempNode := newNode(value)
	tempNode.next = p.next
	p.next = tempNode
	n.length++
	return true
}

// AddNodeBefore 在某个节点之前插入节点
func (n *NodeList[T]) AddNodeBefore(p *Node[T], value T) bool {
	if p == nil {
		return false
	}
	cur := n.head
	for cur.next != nil {
		if cur.next == p {
			break
		}
		cur = cur.next
	}
	if cur == nil {
		return false
	}

	tempNode := newNode(value)
	tempNode.next = cur.next
	cur.next = tempNode
	n.length++
	return true
}

// DeleteNode 删除某个节点
func (n *NodeList[T]) DeleteNode(p *Node[T]) bool {
	cur := n.head
	for cur.next != nil {
		if cur == p {
			cur.next = cur.next.next
			n.length--
			return true
		}
		cur = cur.next
	}
	return false
}

// AddNodeTail 在尾部插入节点
func (n *NodeList[T]) AddNodeTail(value T) bool {
	cur := n.head
	for cur.next != nil {
		cur = cur.next
	}
	return n.addNodeAfter(cur, value)
}

// AddNodeHead 在头部插入节点
func (n *NodeList[T]) AddNodeHead(value T) bool {
	return n.addNodeAfter(n.head, value)
}

func (n *NodeList[T]) Print() {
	cur := n.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.value)
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

// GetNode 通过游标获取节点
func (n *NodeList[T]) GetNode(index int) *Node[T] {
	if index >= n.length {
		return nil
	}

	cur := n.head
	for i := 0; i < index; i++ {
		if cur == nil {
			return nil
		}
		cur = cur.next
	}
	return cur
}
