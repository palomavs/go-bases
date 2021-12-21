package main

import (
	"errors"
	"fmt"
	"os"
)

func verificarSalario(salary int) error {
	if salary < 150000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return nil
}

func main() {
	salary := 1000000
	err := verificarSalario(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Debe pagar impuesto\n")
}
