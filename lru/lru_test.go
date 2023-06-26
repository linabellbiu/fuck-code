package lru

import (
	"fmt"
	"testing"
)

func Test_set(t *testing.T) {
	l := NewLru(6)
	l.set(1)
	l.set(2)
	l.set(3)
	l.print()
	l.set(2)
	l.print()
	l.set(5)
	l.set(6)
	l.print()
	l.set(7)
	l.print()
	l.set(8)
	l.print()
	l.set(4)
	l.print()
}

func Test_Delete(t *testing.T) {
	l := NewLru(10)
	l.set(1)
	l.set(2)
	l.set(3)
	l.set(4)
	l.set(5)
	l.set(6)
	l.print()

	node := l.find(4)
	l.delete(node)
	l.print()

	node = l.find(1)
	l.delete(node)
	l.print()

	node = l.find(6)
	l.delete(node)
	l.print()

	node = l.find(5)
	l.delete(node)
	l.print()

	node = l.find(3)
	l.delete(node)
	l.print()

	node = l.find(2)
	l.delete(node)
	l.print()
}

func Test_findByIndex(t *testing.T) {
	l := NewLru(10)
	l.set(1)
	l.set(2)
	l.set(3)
	l.set(4)
	l.set(5)
	l.set(6)
	l.print()

	node := l.findByIndex(2)
	fmt.Println(node.value)
}
