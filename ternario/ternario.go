package main

import "fmt"

//Nao tem operador ternario
func ObterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

//tem como fazer o if de forma mais simplificada e tão eficiente quanto

/*func ObterResultado2(nota1 float64) string{
	return nota1 >= 6? "Aprovado" : "Reprovado"
}
*/
func main() {
	fmt.Println(ObterResultado(9))
}
