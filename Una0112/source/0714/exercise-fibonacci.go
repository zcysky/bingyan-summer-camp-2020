package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	b1,b2 := 0,1
	return func()int{
		t := b1
		b1,b2 = b2,(b1+b2)
		return t
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
