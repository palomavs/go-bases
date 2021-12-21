package main

import (
	"errors"
	"fmt"
	"os"
)

type errorNegativo struct{}

type errorLimiteHs struct{}

func (e *errorNegativo) Error() string {
	return "error, se ingresó un valor negativo"
}

func (e *errorLimiteHs) Error() string {
	return "error, el trabajador no puede haber trabajado menos de 80 hs mensuales"
}

func calcularSalario(hs float32, precioHs float32) (float32, error) {
	salarioMensual := hs * precioHs

	if hs < 80 {
		return 0, fmt.Errorf("%w: %.2f hs horas trabajadas", &errorLimiteHs{}, hs)
	}

	if salarioMensual >= 150000 {
		salarioMensual *= 0.9
	}

	return salarioMensual, nil
}

func calcularAguinaldo(mesesTrabajados int, salarios ...float32) (float32, error) {
	mejorSalario := salarios[0]

	if mesesTrabajados < 1 {
		return 0, fmt.Errorf("%w: el trabajador no puede haber trabajado menos de 1 mes (se indicó %d mes)", &errorNegativo{}, mesesTrabajados)
	}

	for _, salario := range salarios {
		if salario < 0 {
			return 0, fmt.Errorf("%w: no se puede ingresar un salario negativo ($ %.2f)", &errorNegativo{}, salario)
		} else if salario > mejorSalario {
			mejorSalario = salario
		}
	}

	aguinaldo := mejorSalario / 12 * float32(mesesTrabajados)
	return aguinaldo, nil
}

func main() {
	salario1, errs1 := calcularSalario(123, 44)
	salario2, errs2 := calcularSalario(60, 44)
	if errs1 != nil {
		fmt.Println(errors.Unwrap(errs1))
		os.Exit(1)
	} else {
		fmt.Printf("El salario del empleado 1 es $%.2f\n\n", salario1)
	}

	if errs2 != nil {
		fmt.Println(errors.Unwrap(errs2))
		fmt.Println(errs2)
		os.Exit(1)
	} else {
		fmt.Printf("El salario del empleado 2 es $%.2f\n\n", salario2)
	}

	aguinaldo1, erra1 := calcularAguinaldo(5, 123.3, 3242.3, 453.3)
	aguinaldo2, erra2 := calcularAguinaldo(0, 123.3, 3242.3, 453.3)

	if erra1 != nil {
		fmt.Println(errors.Unwrap(erra1))
		os.Exit(1)
	} else {
		fmt.Printf("El aguinaldo del empleado 1 es $%.2f\n\n", aguinaldo1)
	}
	if erra2 != nil {
		fmt.Println(errors.Unwrap(erra2))
		fmt.Println(erra2)
		os.Exit(1)
	} else {
		fmt.Printf("El aguinaldo del empleado 2 es $%.2f\n\n", aguinaldo2)
	}
}
