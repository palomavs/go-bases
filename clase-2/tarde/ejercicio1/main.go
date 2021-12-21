package main

import "fmt"

type Alumno struct {
	Nombre, Apellido string
	DNI              int
	Ingreso          Fecha
}

type Fecha struct {
	Dia, Mes, Año uint16
}

func (alumno Alumno) detalle() {
	fmt.Printf("Nombre: %s, Apellido: %s, DNI: %d, Fecha de Ingreso: %d-%d-%d\n\n", alumno.Nombre, alumno.Apellido, alumno.DNI, alumno.Ingreso.Dia, alumno.Ingreso.Mes, alumno.Ingreso.Año)
}

func main() {
	fmt.Println("--Ej. 1--")

	a := Alumno{"Clara", "Perez", 12121212, Fecha{1, 3, 2007}}
	a.detalle()
}
