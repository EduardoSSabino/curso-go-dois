package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)                                 //numeros inteiros
	fmt.Println("Literal interio é", reflect.TypeOf(32000)) //imprimi o tipo do valor dentro dos parenteses

	//sem sinal so positivos... uint8 uint16 uint 32...
	var b byte = 255
	fmt.Println("O byte é ", reflect.TypeOf(b))

	//com sinal... iint 8 in 16...
	i := math.MaxInt64 //o valor maximo do da variavel int
	fmt.Println("O valor maximo do int é", i)
	fmt.Println("O tipo de i é", reflect.TypeOf(i))

	var i1 rune = 'a' //representa um mapeamento da tabela Unicode int 32
	fmt.Println("O rune é", reflect.TypeOf(i1))
	fmt.Println(i1)

	//numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(x))
	fmt.Println("O tipo do lietral 49.99 é", reflect.TypeOf(49.99)) //float 64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "Olá meu nome é Leo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1)) //len mostra o tamanho

	//string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Eduardo`
	fmt.Println("O tamanho da string é", len(s2))

	//char
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))

}
