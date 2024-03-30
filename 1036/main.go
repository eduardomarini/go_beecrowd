package main

import (
	"fmt"
	"math"
)

func bhaskara(a, b, c float64) (float64, float64, error) {

	delta := math.Pow(b, 2) - 4*a*c

	if delta < 0 || a == 0 {
		return 0, 0, fmt.Errorf("impossivel calcular")
	}

	x1 := (-b + math.Sqrt(delta)) / (2 * a)
	x2 := (-b - math.Sqrt(delta)) / (2 * a)

	return x1, x2, nil
}

func main() {

	var a, b, c float64

	fmt.Println("Digite o valor de A: ")
	fmt.Scan(&a)
	fmt.Println("Digite o valor de B: ")
	fmt.Scan(&b)
	fmt.Println("Digite o valor de C: ")
	fmt.Scan(&c)

	x1, x2, err := bhaskara(a, b, c)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("R1 = %.5f\n", x1)
		fmt.Printf("R2 = %.5f\n", x2)
	}
}
