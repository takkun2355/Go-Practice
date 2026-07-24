package main

import (
	"fmt"
)

func number() {
	var calc int
	var n [2]int
	var op [1]string
	calc = 0

	fmt.Scan(&n[0], &op[0], &n[1])
	fmt.Println("式:", n[0], op[0], n[1])

	switch op[0] {
	case "+":
		fmt.Println(n[0] + n[1])
	case "-":
		fmt.Println(n[0] - n[1])
	case "*":
		fmt.Println(n[0] * n[1])
	case "**":
		if n[1] <= 1 {
			fmt.Println(n[0])
		} else {
			calc += (n[0] * n[0])
			for i := 0; i < n[1]-2; i++ {
				calc += (calc * n[0])
			}
		}
		fmt.Println(calc)
	case "%":
		fmt.Println(int(n[0]) % int(n[1]))
	case "/%":
		fmt.Println(int(n[0])/int(n[1]), n[0]%n[1])
	case "//":
		fmt.Println(int(n[0]) / int(n[1]))
	case "/":
		fmt.Println(float64(n[0]) / float64(n[1]))
	default:
		fmt.Println("不明なオプション")
	}
}
