package main

import "fmt"
import "math/big"

func main() {
	a, b := "", ""
	fmt.Scanf("%s%s", &a, &b)
	num1, num2 := new(big.Int), new(big.Int)
	num1.SetString(a, 10)
	num2.SetString(b, 10)
	num1 = num1.Add(num1, num2)

	fmt.Println(num1.String())
}
