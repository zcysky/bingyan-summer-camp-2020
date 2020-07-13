package main

import (
	"fmt"
	"math/big"
)

func BigNumAdd(){

	var num_str[2] string

	num1 := new(big.Int)
	num2 := new(big.Int)
	answer := new(big.Int)

	fmt.Printf("Please enter the first big num: ")
	fmt.Scanf("%s", &num_str[0])

	fmt.Printf("Please enter the second big num: ")
	fmt.Scanf("%s", &num_str[1])

	num1.SetString(num_str[0], 10)
	num2.SetString(num_str[1], 10)

	answer.Add(num1, num2)

	fmt.Println("the sum of them is :" + answer.String())
}
