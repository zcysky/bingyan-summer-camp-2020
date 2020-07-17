package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int){
	if t.Left!=nil{
		Walk(t.Left,ch)
	}
	ch<-t.Value
	if t.Right!=nil{
		Walk(t.Right,ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool{
	t:=10
	ch1,ch2:=make(chan int,t),make(chan int,t)
	go Walk(t1,ch1)
	go Walk(t2,ch2)
	for i:=0;i<10;i+=1{
		if <-ch1!=<-ch2{
			return false
		}
	}
	return true
}

func main() {
	t1 := tree.New(10)
	t2 := tree.New(10)
	fmt.Println(Same(t1,t2))
}
