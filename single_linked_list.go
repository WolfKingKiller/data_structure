package data_structure

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)


//单链表

type Node struct {
	data int
	next *Node
}

//创建头结点，这里头结点不存放数据
func BuildHead() *Node {
	return &Node{
		data: -999,
		next: nil,
	}
}

//创建空节点,这里给data赋值为-999，表示空
func BuildEmptyNode() *Node {
	return &Node{
		data: -999,
		next: nil,
	}
}

//判断链表是否为空
func IsEmpty(head *Node) bool {
	if head.next==nil {
		return true
	}
	return false
}

//判断下一个结点是否为空
func IsEmptyNext(node *Node) bool {
	if node.next.data==-999 {
		return true
	}
	return false
}
//返回链表长度
func GetLinkedListLength(head *Node) int {
	length:=0
	for head.next != nil {
		head=head.next
		length++
	}
	return length
}
//将链表的数据升序排列
func SortIncrease(head *Node)  {

}

//有序插入结点
func InsertToIncrease(node,head *Node)  {
	n:=head
	if node.data <= head.next.data {
		node.next=head.next
		head.next=node
		//head=n
	}else if node.data >= GetLastNode(head).data {
		GetLastNode(head).next=node
	}else{
		for i := 1; i <= GetLinkedListLength(n); i++ {
			if node.data >= head.next.data && node.data <= head.next.next.data {
				node.next=head.next.next
				head.next.next=node
				//head=n
				break
			}else{
				head=head.next
			}
		}
	}
}

//查找链表的尾结点
func GetLastNode(node *Node) *Node {
	for node.next!=nil {
		node=node.next
	}
	return node
}

//在链表尾插入结点
func InsertToListLast(node,head *Node) error {
	lastNode := GetLastNode(head)
	lastNode.next=node
	return nil
}

//链表为空时，插入结点
func InsertToFirstNode(head, node *Node) *Node {
	head.next=node
	return head
}

//遍历链表
func GetAll(node *Node) int {
	return node.next.data
}

//根据对应位置删除该元素
func DeleteElem(index int,head *Node) error {
	if index>GetLinkedListLength(head) {
		return errors.New("索引超出链表长度,删除失败！")
	}
	for i := 1; i < index; i++ {
		head=head.next
	}
	head.next=head.next.next
	return nil
}

//创建一个链表
func BuildLinkedList() {
	head := BuildHead()
	rand.Seed(time.Now().Unix())
	for i := 1; i <= 10; i++ {
		n:=head
		data := rand.Intn(100)
		node := BuildEmptyNode()
		node.data=data
		if IsEmpty(head) == true {
			head = InsertToFirstNode(head, node)
		}else{
			//_ = InsertToListLast(node,head)
			InsertToIncrease(node,head)
		}
		for i := 1; i <= GetLinkedListLength(head); i++ {
			fmt.Print(GetAll(n),"——>")
			n=n.next
		}
		fmt.Print("\n")
	}
	n:=head
	err := DeleteElem(2, head)
	if err != nil {
		fmt.Println("\n",err)
		fmt.Print("\n")
	}
	fmt.Println("最后的链表元素为：")
	for i := 1; i <= GetLinkedListLength(head); i++ {
		fmt.Print(GetAll(n),"——>")
		n=n.next
	}
}
