package main

import (
	"fmt"
)

func notas(valor float64) {

	notas := [6]float64{100, 50, 20, 10, 5, 2}
	moedas := [6]float64{1.0, 0.50, 0.25, 0.10, 0.05, 0.01}

	for _, notas := range notas {

		quantidade_notas := int(valor / notas)
		valor = valor - float64(quantidade_notas)*notas

		fmt.Printf("%d nota(s) de R$ %.2f\n", quantidade_notas, notas)
	}

	for _, moedas := range moedas {

		quantidade_moedas := int(valor / moedas)
		valor = valor - float64(quantidade_moedas)*moedas

		fmt.Printf("%d moeda(s) de R$ %.2f\n", quantidade_moedas, moedas)
	}
}

func main() {

	var valor float64

	fmt.Println("Digite o valor: ")
	fmt.Scan(&valor)

	notas(valor)
}
