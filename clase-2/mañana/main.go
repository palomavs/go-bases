package main

import (
	"errors"
	"fmt"
)

const (
	minimo    = "minimo"
	maximo    = "maximo"
	promedio  = "promedio"
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func calcularImpuesto(salario float32) float32 {
	if salario >= 150000 {
		return salario * 0.27
	} else if salario >= 50000 {
		return salario * 0.17
	} else {
		return 0
	}
}

func calcularPromedio(notas ...int32) (float32, error) {
	var sumatoria float32

	for _, nota := range notas {
		sumatoria += float32(nota)
		if nota < 0 {
			return 0, errors.New("La nota no puede ser negativa.")
		}
	}

	return sumatoria / float32(len(notas)), nil
}

func calcularSalario(mins float32, categoria string) (float32, error) {
	hours := mins / 60

	switch categoria {
	case "C":
		return 1000 * hours, nil
	case "B":
		return (1500 * hours) * 1.2, nil
	case "A":
		return (3000 * hours) * 1.5, nil
	default:
		return 0, errors.New("La categoría no existe.")
	}
}

func calcularEstadisticas(operacion string) (func(calificaciones ...int) int, error) {

	switch operacion {
	case minimo:
		return func(calificaciones ...int) int {
			var min int

			for i, c := range calificaciones {
				if i == 0 {
					min = c
				} else {
					if c < min {
						min = c
					}
				}
			}
			return min
		}, nil
	case maximo:
		return func(calificaciones ...int) int {
			var max int

			for i, c := range calificaciones {
				if i == 0 {
					max = c
				} else {
					if c > max {
						max = c
					}
				}
			}
			return max
		}, nil
	case promedio:
		return func(calificaciones ...int) int {
			var sumatoria int

			for _, c := range calificaciones {
				sumatoria += c
			}

			return sumatoria / len(calificaciones)
		}, nil
	default:
		return nil, errors.New("No se encuentra dicha operación.")
	}
}

func animal(animal string) (func(cantidad int) float32, error) {
	switch animal {
	case perro:
		return func(cantidadPerros int) float32 {
			return float32(cantidadPerros) * 10
		}, nil
	case gato:
		return func(cantidadGatos int) float32 {
			return float32(cantidadGatos) * 5
		}, nil
	case hamster:
		return func(cantidadHam int) float32 {
			return float32(cantidadHam) * 0.25
		}, nil
	case tarantula:
		return func(cantidadTar int) float32 {
			return float32(cantidadTar) * 0.15
		}, nil
	default:
		return nil, errors.New("No existe ese animal.")
	}
}

func main() {
	fmt.Println("--Ej. 1--")
	impuesto := calcularImpuesto(170000)
	fmt.Printf("Impuesto: $ %.2f\n", impuesto)

	fmt.Println("\n--Ej. 2--")
	prom, error := calcularPromedio(8, 7, 5, -10)
	if error == nil {
		fmt.Printf("Promedio: %.2f\n", prom)
	} else {
		fmt.Printf("Promedio: no se pudo calcular el promedio - %s\n", error)
	}

	fmt.Println("\n--Ej. 3--")
	salario, error := calcularSalario(120, "C")
	if error == nil {
		fmt.Printf("Salario: %.2f\n", salario)
	} else {
		fmt.Printf("Salario: no se pudo calcular el salario - %s\n", error)
	}

	fmt.Println("\n--Ej. 4--")
	minFunc, errMin := calcularEstadisticas(minimo)
	maxFunc, errMax := calcularEstadisticas(maximo)
	promFunc, errProm := calcularEstadisticas(promedio)
	_, errIncorr := calcularEstadisticas("incorrecta")

	if errMin == nil {
		valorMinimo := minFunc(10, 5, 6, 3)
		fmt.Printf("ValorMinimo: %d\n", valorMinimo)
	}
	if errMax == nil {
		valorMaximo := maxFunc(9, 4, 7, 5)
		fmt.Printf("ValorMaximo: %d\n", valorMaximo)
	}
	if errProm == nil {
		valorPromedio := promFunc(10, 7, 5, 6)
		fmt.Printf("ValorPromedio: %d\n", valorPromedio)
	}
	if errIncorr != nil {
		fmt.Printf("Error incorrecta: %s\n", errIncorr)
	}

	fmt.Println("\n--Ej. 5--")
	calcularAlimentoPerro, errPerro := animal(perro)
	calcularAlimentoGato, errGato := animal(gato)
	calcularAlimentoHamster, errHam := animal(hamster)
	calcularAlimentoTarantula, errTar := animal(tarantula)
	_, errAnimalIncorr := animal("animal-incorrecto")

	var cantTotal float32
	if errPerro == nil {
		cantTotal = calcularAlimentoPerro(5)
		fmt.Printf("CantTotal - perros: %.2f\n", cantTotal)
	}
	if errGato == nil {
		cantTotal += calcularAlimentoGato(9)
		fmt.Printf("CantTotal + gatos: %.2f\n", cantTotal)
	}
	if errHam == nil {
		cantTotal += calcularAlimentoHamster(4)
		fmt.Printf("CantTotal + ham: %.2f\n", cantTotal)
	}
	if errTar == nil {
		cantTotal += calcularAlimentoTarantula(2)
		fmt.Printf("CantTotal + tar: %.2f\n", cantTotal)
	}
	if errAnimalIncorr != nil {
		fmt.Printf("Error incorrecta: %s\n", errAnimalIncorr)
	}
	fmt.Printf("La cantidad total de alimento a comprar es: %.2f kg\n", cantTotal)
}
