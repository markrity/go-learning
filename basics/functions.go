package main

import "fmt"

func main() {
	var a = 30
	var b = 20
	a, b = sum(a, b)
	fmt.Println(a, b)
}

func sum(num1 int, num2 int) (int, int) {
	var res int
	res = num1 + num2

	return res, num1 * num2
}
