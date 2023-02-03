package main

import "fmt"

func main() {
	// slice literals is like an array literal without the length
	primes := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
	}
	fmt.Println(s)

	s1 := []int{2, 3, 5, 7, 11, 13}

	s1 = s1[:4]
	fmt.Println(s1) // 3, 5, 7

	s1 = s1[:2]
	fmt.Println(len(s1)) //2
}
