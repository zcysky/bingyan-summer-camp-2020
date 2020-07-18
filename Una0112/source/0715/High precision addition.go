package main

import (
	"fmt"
)

func main() {
	var s1,s2 string
	var n1,n2 int
	var a1 [10010] int
	var a2 [10010] int
	fmt.Scan(&s1)
	fmt.Scan(&s2)
	//fmt.Printf("%s\n%s\n",s1,s2)
	n1,n2=len(s1),len(s2)
	//fmt.Print(n1,n2)
	if n1>n2{
		s1,s2=s2,s1
		n1,n2=n2,n1
	}
	for i:=n1-1;i>=0;i-=1{
		a1[i]=(int)(s1[n1-i-1]-'0')
	}
	for i:=n2-1;i>=0;i-=1{
		a2[i]=(int)(s2[n2-i-1]-'0')
	}
	for i:=0;i<n2;i+=1{
		a2[i]+=a1[i]
		a2[i+1]+=a2[i]/10
		a2[i]%=10
	}
	if a2[n2]>0 {
		fmt.Printf("%d",a2[n2])
	}
	for i:=n2-1;i>=0;i-=1{
		fmt.Printf("%d",a2[i])
	}
}
