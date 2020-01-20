package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i > 8 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after")
}
