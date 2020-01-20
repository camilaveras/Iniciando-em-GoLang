package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        if i%2 == 0 { //se isso for feito é o resultado for 0 ele é par, então embaixo o continue faz ele pular para o proximo loop, imprimindo apenas os impares
            continue 
        }
        fmt.Printf("%d ", i)
    }
}

// com continue 1 3 5 7 9 //somente aqueles que são impares
// sem continue 1 2 3 4 5 6 7 8 9 10 // todos os numeros do i
