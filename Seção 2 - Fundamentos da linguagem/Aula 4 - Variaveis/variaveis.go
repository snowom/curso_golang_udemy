package main

import "fmt"

func main() {

	//declaração explicita
	var variavel_1 string = "variavel 1"
	fmt.Println(variavel_1)

	//declaração implicita - inferencia de tipo
	variavel_2 := "variavel 2"
	fmt.Println(variavel_2)

	//declaração de constantes
	const constante_1 string = "constante"
	fmt.Println(constante_1)
}
