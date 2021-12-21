package main

import (
	"fmt"
	"os"
)

func verificarSalario(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: el salario ingresado ($%d) no alcanza el mínimo imponible", salary)
	}
	return nil
}

func main() {
	salary := 1000
	err := verificarSalario(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Debe pagar impuesto\n")
}
