package main

import (
	"fmt"
	"os"
)

type myCustomError struct {
	msg string
}

func (e *myCustomError) Error() string {
	return e.msg
}

func verificarSalario(salary int) error {
	if salary < 150000 {
		return &myCustomError{"error: el salario ingresado no alcanza el mínimo imponible"}
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
