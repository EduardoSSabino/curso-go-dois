package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Stings: ", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0) //trás uma data
	d2 := time.Unix(0, 0)

	fmt.Println("Datas: ", d1 == d2)
	fmt.Println("Datas: ", d1.Equal(d2))

	type Pessoa struct { //cria uma classe Pessoa e define seus atributos, os métodos, eu defino em funções
		Nome   string
		Altura float32
		Peso   float64
	}

	p1 := Pessoa{"João", 1.80, 100}
	p2 := Pessoa{"Marcelo", 1.70, 80}
	fmt.Println("Pessoas: ", p1 == p2)

}
