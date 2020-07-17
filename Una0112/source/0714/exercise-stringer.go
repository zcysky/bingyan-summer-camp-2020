package main

import "fmt"

type Person struct {
	Name string
	IPAddr [4]int
}

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (p Person) String() string {
	return fmt.Sprintf("%v.%v.%v.%v",p.IPAddr[0],p.IPAddr[1],p.IPAddr[2],p.IPAddr[3])
}

func main() {
	a := Person{"loopback", [4]int{127, 0, 0, 1}}
	z := Person{"googleDNS", [4]int{8, 8, 8, 8}}
	fmt.Println(a, z)
}