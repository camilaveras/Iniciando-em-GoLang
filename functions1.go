package main

import (  
    "fmt"
)

func calculateBill(a, b int) int {
	var totalPrice = a - b
    return totalPrice //return no final do calculo como em node.js
}

// func main() {  
//     a, b := 90, 6 //aqui eu to dando um valor para a e b, caso eu n faça isso vai ficar como undefined
//     totalPrice := calculateBill(a, b) //aqui eu to chamando a função calculatebill para que faz o calculo => a-b
//     fmt.Println("Total price is", totalPrice) //bellow i'am want for print the price 
// }

func main(){
	e, f := 10, 5
	totalPrice := calculateBill(e, f)
	fmt.Println("The new price is", totalPrice)
}