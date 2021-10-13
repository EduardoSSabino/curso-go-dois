package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	//Poderiamos ate dizer que um slice é um parte de um array
	a2 := [5]int{1, 2, 3, 4, 5} //array
	s2 := a2[1:3]               //Slice! Esse meu meu slice é um trecho do meu array a2, meu trecho tem inicio da primeira posição até a segunda.
	fmt.Println(a2, s2)
}
