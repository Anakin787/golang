package main

import "fmt"

func main() {
	var dan int
	fmt.Scanf("%d", &dan)

	for i := 1; i < 10; i++ {
		fmt.Printf("%d X %d = %d\n", dan, i, dan*i)
	}
}
