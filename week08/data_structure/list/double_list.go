package main

import "fmt"

type ListNode struct {
	Value int
	Prev *ListNode
	Next *ListNode
}

type DoubleList struct {
	Head *ListNode
	Tail *ListNode
	Length int
}

func (list *DoubleList) Append(x int){
	node := &ListNode{Value: x}
	tail := list.Tail
	if tail == nil{
		list.Head = node
		list.Tail = node
	} else{
		tail.Next = node
		node.Prev = tail
		list.Tail = node
	}
	list.Length+=1
}

//根据 双向链表的索引 idx取出对应的值
func(list *DoubleList) Get(idx int) *ListNode {
	if list.Length <= idx{
		return nil
	}
	curr := list.Head
	for i:=0;i<idx;i++{
		curr=curr.Next
	}
	return curr
}

func (list *DoubleList) InsertAfter(x int,prevNode *ListNode) {
	node := &ListNode{Value: x} //node的Prev、next有待赋值
	if prevNode.Next ==nil { //prevNode本来是尾元素
		prevNode.Next = node
		node.Prev = prevNode
	} else{
		nextNode := prevNode.Next // 获取prevNode的下一个元素，node就要插入到他两之间
		nextNode.Prev = node // 插入一个node，原有的Prev、Next 哪些会受到影响，要考虑清楚
		node.Next = nextNode
		prevNode.Next = node
		node.Prev = prevNode
	}
}

func (list *DoubleList) Traverse(){
	curr := list.Head
	for curr !=nil{
		fmt.Printf("%d",curr.Value)
		curr=curr.Next
	}
	fmt.Println()
}

func testDoubleList(){
	list:=new(DoubleList)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Traverse()
	node:= list.Get(3)
	list.InsertAfter(9,node)
	list.Traverse()
}

func main(){
	//testList()
	testDoubleList()
}