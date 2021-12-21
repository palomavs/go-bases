package main

import (
	"errors"
	"fmt"
	"log"
)

type Matrix struct {
	Valores      []float32
	Alto, Ancho  int
	EsCuadratica bool
	ValorMaximo  float32
}

func (m *Matrix) setValores(alto, ancho int, valores ...float32) error {
	if len(valores) != alto*ancho {
		return errors.New("se indicó una cantidad incorrecta de valores")
	}

	m.Valores = valores
	m.Alto = alto
	m.Ancho = ancho
	m.EsCuadratica = alto == ancho

	max := valores[0]
	for _, val := range valores {
		if val > max {
			max = val
		}
	}
	m.ValorMaximo = max

	return nil
}

func (m Matrix) printMatrix() {
	for i := 0; i < m.Alto; i++ {
		fmt.Printf("[")
		for j := 0; j < m.Ancho; j++ {
			fmt.Printf(" %.2f ", m.Valores[j+m.Ancho*i])
		}
		fmt.Printf("]\n")
	}
	fmt.Printf("\n Ancho: %d, Alto: %d, EsCuadratica: %t, ValorMáximo: %.2f. \n\n", m.Ancho, m.Alto, m.EsCuadratica, m.ValorMaximo)
}

func main() {
	fmt.Println("--Ej. 2--")

	var m Matrix
	err := m.setValores(3, 3, 1, 2, 3, 4.5, 2.7, 12.99, 5.5, 3.67, 10)
	if err != nil {
		log.Fatal(err)
	}
	m.printMatrix()
}
