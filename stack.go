package DataStr

import (
	"fmt"
)

type Stack struct {
	data []interface{}
}

//印出Stack長度
func (this *Stack) Len() int {
	return len(this.data)
}

//印出element
func (this *Stack) Prt() {
	for _, v := range this.data {
		fmt.Print(v, ", ")
	}
	fmt.Println()
}

//放入element
func (this *Stack) Push(value interface{}) {
	this.data = append(this.data, value)
}

//return 最上層的element
func (this *Stack) Peek() (value interface{}) {
	if this.Empty() {
		return nil
	} else {
		return this.data[len(this.data)-1]
	}
}

//return最上層element並移除
func (this *Stack) Pop() (value interface{}) {
	if this.Empty() {
		panic("Stack is empty")
	} else {
		temp := this.data[len(this.data)-1]
		this.data = this.data[0 : len(this.data)-1]
		return temp
	}
}

//回傳此stack是否為空
func (this *Stack) Empty() bool {
	if len(this.data) == 0 {
		return true
	} else {
		return false
	}
}
