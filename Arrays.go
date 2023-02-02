package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Hello1"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [3]int{1, 2, 4}
	fmt.Println(primes[2])
	p := &a
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)

}
