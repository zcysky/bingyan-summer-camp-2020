package main

import (
	"fmt"
)

func main() {
	var s1,s2 string
	var n1,n2 int
	var a1 [10010] int
	var a2 [10010] int
	var a3 [10010] int
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
	for i:=0;i<n1;i++{
		for j:=0;j<n2;j++{
			a3[i+j]+=a1[i]*a2[j];
		}
	}
	m:=n1+n2-1;
	for i:=0;i<m;i++{
		a3[i+1]+=a3[i]/10;
		a3[i]%=10;
	}
	for ;a3[m]==0&&m>=1; {
		m--
	}
	for i:=m;i>=0;i--{
		fmt.Printf("%d",a3[i])
	}
}