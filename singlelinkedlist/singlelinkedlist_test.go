// Package singlelinkedlist 单链操作
package singlelinkedlist

import (
	"fmt"
	"testing"
)

func Test_AddNodeHead(t *testing.T) {
	node := NewNodeList[any]()
	node.AddNodeHead("dfdfg")
	//node.AddNodeHead(2)
	node.Print()
}

func Test_AddNodeBefore(t *testing.T) {
	node := NewNodeList[any]()
	node.AddNodeTail("1")
	node.AddNodeTail("2")
	node.AddNodeTail("3")
	node.AddNodeTail("4")
	node.Print()

	fmt.Println(node.GetNode(3).value)

	node.AddNodeBefore(node.GetNode(3), 5)
	node.Print()
}

func Test_GetNodeBefore(t *testing.T) {
	node := NewNodeList[any]()
	node.AddNodeTail("1")
	node.AddNodeTail("2")
	node.AddNodeTail("3")
	node.AddNodeTail("4")
	node.Print()

	fmt.Println(node.GetNode(3))
}

func Test_DelNode(t *testing.T) {
	node := NewNodeList[any]()
	node.AddNodeTail("1")
	node.AddNodeTail("2")
	node.AddNodeTail("3")
	node.AddNodeTail("4")
	node.Print()

	node.DeleteNode(node.GetNode(0))
	node.Print()
}
