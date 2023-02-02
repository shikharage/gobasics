package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var i, j = 1, 2
	j, k := 3, 4
	fmt.Println(i, j, k)
	fmt.Printf("%v%v%v\n", i, j, k)
	var l int
	fmt.Printf("%v", l)
}
