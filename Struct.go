package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{}
	fmt.Println(v)
	v.X = 3
	fmt.Println(v)
}
