package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string para int
	nmr, err := strconv.Atoi("123")
	fmt.Println(nmr, err)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
