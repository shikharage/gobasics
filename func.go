package main

import "fmt"

func main() {
	fmt.Println(check(10, 15))
}

func check(i, j int) int {
	return i + j
}
