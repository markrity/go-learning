package main

import "fmt"

func main() {
	webSites := make(map[string]int)

	webSites["site1"] = 1
	webSites["site2"] = 2

	fmt.Println(webSites["site2"])
}
