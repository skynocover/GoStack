package DataStr

import (
	"fmt"
)

type node struct {
	value interface{}
	prev  *node
	next  *node
}

type Queue struct {
	head   *node
	rear   *node
	length int
}

func (this *Queue) Len() int { return this.length }

func (this *Queue) Prt() {
	temp := this.head
	if this.length != 0 {
		fmt.Print(temp.value, ", ")
		for temp.next != nil {
			temp = temp.next
			fmt.Print(temp.value, ", ")
		}
	} else {
		fmt.Print(temp)
	}
	fmt.Println()
}

func (this *Queue) Push(value interface{}) {
	newnode := node{value, this.rear, nil}
	if this.length != 0 {
		this.rear = &newnode
		this.head = &newnode
	} else {
		this.rear.next = &newnode
		this.rear = &newnode
	}
	this.length++
}

func (this *Queue) Peek() interface{} {
	if this.length == 0 {
		panic("Queue is empty")
	} else {
		return this.head.value
	}
}

func (this *Queue) Get() interface{} {
	if this.length == 0 {
		panic("Queue is empty")
	} else {
		temp := this.head.value
		this.head = this.head.next
		this.length--
		return temp
	}
}
