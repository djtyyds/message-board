package service

import (
	"fmt"
)

type Node struct { //创建一个左斜树
	Left *Node
	Data interface{}
}

func NewNode(left *Node) *Node {
	return &Node{left, nil}
}

func (node *Node) Print() { //输出
	fmt.Print(node.Data, " ")
}
func (node *Node) SetData(comment interface{}) { //给节点赋值
	if node == nil {
		fmt.Println("setting value to nil.node ignored.")
		return
	}
	node.Data = comment
}

func (node *Node) PreOrder() { //前序遍历,因为只有左斜树，只遍历左斜树
	if node == nil {
		return
	}
	node.Print()
	node.Left.PreOrder()
}
