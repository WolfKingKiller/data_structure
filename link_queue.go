package queue

import (
	"errors"
	"fmt"
	"strings"
)

//实现链队列

type Node struct {
	data interface{}
	next *Node
}

type LinkQueue struct {
	front *Node
	rear *Node
}

//创建空结点
func NewNode() *Node {
	return &Node{
		data: nil,
		next: nil,
	}
}

//初始化队列(构造一个空队列),队头指针和队尾指针均指向头结点
func NewLinkQueue(head *Node) *LinkQueue {
	return &LinkQueue{
		front: head,
		rear: head,
	}
}

//判断队列是否为空
func IsEmpty(linkQueue *LinkQueue,head *Node) bool {
	if linkQueue.rear==head&&linkQueue.front==head {
		return true
	}
	return false
}

//返回队列长度
func GetLength(head *Node) int {
	length:=0
	for head.next != nil {
		length++
		head=head.next
	}
	return length
}

//找到队尾结点，并返回
func LastNode(head *Node) *Node {
	lastEle:=head
	for i := 1; i <= GetLength(head); i++ {
		lastEle=lastEle.next
	}
	return lastEle
}

//插入新元素作为队尾元素
func EnterQueue(linkQueue *LinkQueue, head,node *Node) {
	lastNode := LastNode(head)
	linkQueue.rear=node
	lastNode.next=node
}

//遍历队列
func VisitQueue(head *Node) {
	fmt.Println("遍历队列的所有元素：")
	for head.next != nil {
		fmt.Print(head.next.data,"-->")
		head=head.next
	}
}

//若队列不为空，则删除队头元素,并返回队头元素，否则，返回error
func DeleteEle(linkQueue *LinkQueue,head *Node) (interface{},string,error) {
	if IsEmpty(linkQueue,head) {
		return nil,"", errors.New("非法操作，队列为空！")
	}
	ele:=head.next.data
	head.next=head.next.next
	return ele,"删除成功！",nil
}
//创建测试队列
func BuildLinkQueue() {
	//创建头结点
	head := NewNode()
	//创建空队列
	linkQueue := NewLinkQueue(head)
	var length int
	fmt.Println("请输入创建队列的长度：")
	_, _ = fmt.Scan(&length)
	var num int
	for i := 1; i <= length; i++ {
		fmt.Println("请输入插入的数据：")
		_, _ = fmt.Scan(&num)
		//创建空结点
		newNode := NewNode()
		newNode.data=num
		newNode.next=nil
		EnterQueue(linkQueue,head,newNode)
	}
	VisitQueue(head)
	fmt.Println()
	var flag string
	for i := 1; i <= length; i++ {
		fmt.Println("是否删除队头元素？y/n")
		_, _ = fmt.Scan(&flag)
		if strings.Compare("y", flag) == 0 {
			ele, _, err := DeleteEle(linkQueue, head)
			if err != nil {
				fmt.Println(err)
				break
			}else {
				fmt.Printf("成功删除队头元素：%v\n",ele)
			}
		}else {
			break
		}
	}
	fmt.Println("删除之后：")
	VisitQueue(head)
}