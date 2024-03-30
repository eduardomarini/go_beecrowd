package main

import "fmt"

func salario(horas_trab int, valor_hora float64) float64 {
	return float64(horas_trab) * valor_hora
}

func main() {

	var horas_trab, num_funcionario int
	var valor_hora float64

	fmt.Println("Digite o número do funcionário: ")
	fmt.Scan(&num_funcionario)

	fmt.Println("Digite o número de horas trabalhadas: ")
	fmt.Scan(&horas_trab)

	fmt.Println("Digite o valo que recebe por horas trabalhada: ")
	fmt.Scan(&valor_hora)

	fmt.Println("NUMBER = ", num_funcionario)

	sal := salario(horas_trab, valor_hora)
	fmt.Printf("SALARY = U$ %.2f\n", sal)
}
