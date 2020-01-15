package main

import "fmt"

func main() {  
	a, b := true, false
		
	c := a && b // retorna true apenas de a e b são true
	
	fmt.Println("c:", c)
	
	d := a || b  //retorna true quando a ou b é true
	
	fmt.Println("d:", d)
}

c: false
d: true