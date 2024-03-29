package main

import "fmt"

type Node struct{
	Data interface{}
	Next *Node  //  Next = &Node
}

type LList struct{
	Head *Node
	Length int
}

type Method interface {
	Insert(i int,v interface{})  //增
	Delete(i int)                //删
	GetLength() int      //getLength
	Search(v interface{}) int          //查
	isNull()bool                 //判断链表是否为空
}
//创建结点
func CreateNode(v interface{})*Node{
	return &Node{v,nil}
}
//创建空链表
func CreateList()*LList{
	return &LList{CreateNode(nil),0}
}

//i处插入节点
func (list *LList)Insert(i int,v interface{}){
	s := CreateNode(v)  //*Node
	pre := list.Head  //*Node
	for count:=0;count<=i-1;count++{
		if count == i-1{
			s.Next = pre.Next
			pre.Next = s
			list.Length++
		}
		pre = pre.Next

	}
}
//删除i处节点
func (list *LList)Delete(i int){
	pre := list.Head
	for count:=0;count<=i-1;count++{
		s := pre.Next
		if count == i-1{
			pre.Next = s.Next
			list.Length--
		}
		pre = pre.Next
	}
}
//返回链表长度
func (list *LList)GetLength() int{
	pre := list.Head
	for pre.Next!=nil{
		list.Length++
	}
	return list.Length
}
//查值v所在位置
func (list *LList)Search(v interface{}) int{
	pre := list.Head.Next
	for i:=1;i<=list.Length;i++{
		if pre.Data == v{
			return i
		}
		pre = pre.Next
	}
	return 0
}

func (list *LList) isNull() bool{
	pre := list.Head.Next
	if pre == nil{
		return true
	}
	return false
}
//打印链表
func PrintList(list *LList){
	pre := list.Head  //*Node
	fmt.Println("LList shows as follow:...")
	for i:=0;i<=list.Length;i++{
		fmt.Printf( "%v\n",pre)
		pre  = pre.Next
	}
}


func main(){
	lList := CreateList()
	//fmt.Println("List is Null:",lList.isNull())
	var M Method
	M = lList
	M.Insert(1,3)
	M.Insert(2,6)
	M.Insert(1,5)
	PrintList(lList)
	fmt.Println("List length is:",lList.Length)
	fmt.Println("元素6在位置  ",M.Search(6))
	fmt.Println("元素100在位置  ",M.Search(100))
	fmt.Println("List is Null:",lList.isNull())
	M.Delete(2)
	PrintList(lList)
	fmt.Println("List length is:",lList.Length)
}
