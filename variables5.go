package main

import "fmt"

func main(){
	name := "camila" // forma um de declarar variavel
	sobrenome := "veras" // forma um de declarar variavel
	name, sobrenome := "camilaaa", "verassss" // forma dois de declarar variavel, comentar alguma das formas para ver ela em ação
	fmt.Println("variable name is", name, "variable sobrenome is", sobrenome)
}