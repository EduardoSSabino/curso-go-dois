package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10) //gera um numero aleatorio ate 10
}

func main() {
	if i := numeroAleatorio(); i > 5 { //também suportado no switch
		fmt.Println("Ganhou!")
		fmt.Println(i)
	} else {
		fmt.Println("Perdeu!")
	}
}
