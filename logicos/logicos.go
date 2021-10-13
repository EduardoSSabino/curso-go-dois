package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) { //executa primeiros
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() { //função principal, é executado depois. "A função que resolve o problema"
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Sadável: %t", tv50, tv32, sorvete, !sorvete)
}
