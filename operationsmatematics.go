package main

import (
	
	"fmt"
	"math"
)
func main(){
	a, b := 142.8, 350.9
	c := math.Min(a,b)
	fmt.Println("the value minimum is", c)
}
