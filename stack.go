package stack

import (
	"errors"
	"fmt"
	"strings"
)

//实现栈的基本操作

//定义栈
type Stack struct {
	size int
	top int
	data []interface{}
}

//初始化栈
func NewStack(size int) *Stack {
	return &Stack{
		size: size,
		top: -1,
		data: make([]interface{},size),
	}
}

//判断栈是否已经满了
func IsFull(stack *Stack) bool {
	if stack.top==stack.size-1 {
		return true
	}
	return false
}

//判断栈是否为空
func IsEmpty(stack *Stack) bool {
	if stack.top==-1 {
		return true
	}
	return false
}

//得到栈顶元素
func GetTop(stack *Stack) (interface{},error) {
	if IsEmpty(stack)==true {
		return nil,errors.New("非法操作：栈为空！")
	}
	return stack.data[stack.top],nil
}

//入栈
func Push(data interface{},stack *Stack) (string,error) {
	if IsFull(stack)==true {
		return "",errors.New("非法操作：栈已满,插入失败！")
	}
	stack.top++
	stack.data[stack.top]=data
	return "插入成功",nil
}

//遍历栈(出栈)
func Traverse(stack *Stack) error {
	if IsEmpty(stack) == true {
		return errors.New("非法操作：栈为空！")
	}
	for i:=stack.top;i>=0;i-- {
		fmt.Print(stack.data[i]," ")
	}
	fmt.Println()
	return nil
}

//建立一个测试栈
func BuildStack() {
	var size int
	fmt.Println("请输入测试栈的大小：")
	//想要分行读取数据只能用scan和Scanln，不能用scanf，因为它会读取我们输入的换行符，导致后面读取发生错误
	_, _ = fmt.Scanln(&size)
	stack := NewStack(size)
	var data int
	var flag string
	for i := 0; ; i++ {
		fmt.Println("请插入数据：")
		_, _ = fmt.Scanln( &data)
		_, err := Push(data, stack)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("是否继续插入？y/n")
		_, _ = fmt.Scanln(&flag)
		if strings.Compare("y", flag) == 0 {
			continue
		}else{
			break
		}
	}
	err := Traverse(stack)
	if err != nil {
		fmt.Println(err)
	}
	elem, err := GetTop(stack)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(elem)
}

