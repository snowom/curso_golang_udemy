package main

import (
	"errors"
	"fmt"
)

func main() {

	// Numeros Inteiros - Negativos e Positivos
	var var1 int8 = 0  // cria uma variável int que suporta até 8 bits
	var var2 int16 = 0 // cria uma variável int que suporta até 16 bits
	var var3 int32 = 0 // cria uma variável int que suporta até 32 bits
	var var4 int64 = 0 // cria uma variável int que suporta até 64 bits
	var var5 int = 0   // cria uma variável int que suporta a quantidade de bits máximo que o processador da máquina suporta
	fmt.Println(var1, var2, var3, var4, var5)

	// Numeros Inteiros Positivos - UINT (unsigned int)
	var var6 uint8 = 0  // cria uma variável int positiva que suporta até 8 bits
	var var7 uint16 = 0 // cria uma variável int positiva que suporta até 16 bits
	var var8 uint32 = 0 // cria uma variável int positiva que suporta até 32 bits
	var var9 uint64 = 0 // cria uma variável int positiva que suporta até 64 bits
	var var10 uint = 0  // cria uma variável int positiva positiva que suporta a quantidade de bits máximo que o processador da máquina suporta
	fmt.Println(var6, var7, var8, var9, var10)

	// alias de int32 - Usados também para lidar com tabela ASCII
	var var11 rune = -20

	// alias de uint8
	var var12 byte = 8
	fmt.Println(var11, var12)

	// Numeros Reais
	var var13 float32 = 0.0 // cria uma variavel float que suporta até 32 bits
	var var14 float64 = 0.0 // cria uma variavel float que suporta até 64 bits
	var15 := 0.89           // cria uma variavel float que suporta a quantidade de bits máximo que o processador da máquina suporta
	fmt.Println(var13, var14, var15)

	// String
	var var16 string = "teste" // cria uma variável do tipo string
	fmt.Println(var16)

	// "Char" - Não existe no Go
	var17 := 'T' // cria uma variável do tipo int32 que terá como valor a posição do caractere na tabela ASCII
	fmt.Println(var17)

	// Booleano
	var var18 bool = true  // cria uma variável booleana com valor verdadeiro
	var var19 bool = false // cria uma variável booleana com valor falso - valor 0
	fmt.Println(var18, var19)

	//Error
	var var20 error = errors.New("Novo Erro")
	fmt.Println(var20)
}
