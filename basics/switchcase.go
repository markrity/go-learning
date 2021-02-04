package main

import "fmt"

func main() {
	var age = 10
	switch age {
	case 5:
		fmt.Println("Age is 5")
	case 10:
		fmt.Println("Age is 10")
	default:
		fmt.Println("Age uknown")
	}
}
