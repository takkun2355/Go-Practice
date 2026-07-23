package main

import (
	"fmt"
)

func number() {
	var n [2]int
	var op [1]string
	fmt.Scan(&n[0], &op[0], &n[1])
	fmt.Println("式:", n[0], op[0], n[1])

	switch op[0] {
	case "+":
		{
			fmt.Println(n[0] + n[1])
		}
	case "-":
		{
			fmt.Println(n[0] - n[1])
		}
	case "*":
		{
			fmt.Println(n[0] * n[1])
		}
	case "/":
		{
			fmt.Println(int(n[0]/n[1]), int(n[0]%n[1]))
		}
	default:
		{
			fmt.Println("不明なオプション")
		}
	}
}
