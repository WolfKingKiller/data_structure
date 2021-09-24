package double_linked_list

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//双链表

type Node struct {
	data int
	next *Node
	prior *Node
}

//创建头结点，头结点的data不存储任何信息,data里面的-999不存在任何实际意义
func BuildHead() *Node {
	return &Node{
		data: -999,
		next: nil,
		prior: nil,
	}
}

//创建空结点，并传入结点数据
func BuildNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
		prior: nil,
	}
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

//传入头结点,判断链表是否为空
func IsEmpty(head *Node) bool {
	if head.next==nil {
		return true
	}
	return false
}

//有序插入结点
func InsertToIncrease(node,head *Node)  {
	n:=head
	if node.data<=head.next.data {
		node.next=head.next
		node.prior=head
		head.next=node
	}else if GetLastNode(head).data<=node.data {
		n:=GetLastNode(head)
		n.next=node
		node.prior=n
	}else {
		for i := 1; i <= GetLinkedListLength(n); i++ {
			if node.data >= head.next.data && node.data <= head.next.next.data {
				node.next=head.next.next
				node.prior=head.next
				head.next.next.prior=node
				head.next.next=node
				break
			}else {
				head=head.next
			}
		}
	}
}

//链表为空时，插入结点
func InsertToFirstNode(head, node *Node) *Node {
	head.next=node
	node.prior=head
	return head
}

//查找链表的尾结点
func GetLastNode(node *Node) *Node {
	for node.next!=nil {
		node=node.next
	}
	return node
}

//正向遍历链表
func GetAllForward(head *Node) int {
	return head.next.data
}

//反向遍历链表
func GetAllOpposite(node *Node) int {
	return node.data
}

//删除对应位置的元素
func DeleteElem(index int, head *Node) error {
	if index> GetLinkedListLength(head) {
		return errors.New("索引超出链表长度,删除失败！")
	}
	for i := 1; i < index; i++ {
		head=head.next
	}
	head.next.next.prior=head
	head.next=head.next.next
	return nil
}

//利用随机数创建一个有序双链表
func BuildDoubleLinkedList() {
	head := BuildHead()
	rand.Seed(time.Now().Unix())
	for i := 1; i <= 10; i++ {
		n:=head
		data := rand.Intn(100)
		fmt.Print(data,"\t")
		node := BuildNode(data)
		if IsEmpty(head) == true {
			head = InsertToFirstNode(head, node)
		}else {
			InsertToIncrease(node,head)
		}
		for i := 1; i <= GetLinkedListLength(head); i++ {
			fmt.Print(GetAllForward(n),"——>")
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
	fmt.Println("正向遍历链表的元素为：")
	for i := 1; i <= GetLinkedListLength(head); i++ {
		fmt.Print(GetAllForward(n),"——>")
		n=n.next
	}
	fmt.Print("\n")
	n2:=GetLastNode(head)
	fmt.Println("反向遍历链表的元素为：")
	for i := 1; i <= GetLinkedListLength(head); i++ {
		fmt.Print(GetAllOpposite(n2),"——>")
		n2=n2.prior
	}
}

