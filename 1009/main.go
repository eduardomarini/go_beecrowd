package main

import "fmt"

func comissao(salario_fixo, total_vendas float64) float64 {
	return (total_vendas * 0.15) + salario_fixo
}

func main() {

	var nome string
	var salario_fixo, total_vendas float64

	fmt.Println("Digite o nome do funcionário: ")
	fmt.Scan(&nome)

	fmt.Println("Digite o salário fixo do funcionário:")
	fmt.Scan(&salario_fixo)

	fmt.Println("Digite o total de vendas do funcionário: ")
	fmt.Scan(&total_vendas)

	novo_salario := comissao(salario_fixo, total_vendas)
	fmt.Printf("TOTAL = R$ %.2f\n", novo_salario)
}
