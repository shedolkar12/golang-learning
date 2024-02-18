package main

import (
	"fmt"
	"time"
)

func f(n int) {
	fmt.Println(time.TimeOnly)
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(0)
	}
	var input string = "rajesh"
	// fmt.Scanln(&input)
	time.Sleep(5)
	fmt.Println(input)
}
