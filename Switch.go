package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	fmt.Println(os)

	switch os {
	case "ubu":
		fmt.Println("ubu")
	case "windows":
		fmt.Println("win")
	default:
		fmt.Println("def")
	}
}
