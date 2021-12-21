package main

import "fmt"

// CLASE 16/12

func ejercicio1() {
	palabra := "computadora"

	fmt.Printf("Cantidad de letras de la palabra \"%s\": %d\n", palabra, len(palabra))

	for _, character := range palabra {
		fmt.Printf("letra: %s\n", string(character))
	}
}

func ejercicio2(precio float32, dcto int) {
	precioFinal := precio - precio*(float32(dcto)/100)

	fmt.Printf("Precio con descuento del %d: %.2f\n", dcto, precioFinal)
}

func ejercicio3(edad int, empleado bool, antiguedad int, sueldo float32) {

	/* switch {
	case edad > 22 && empleado && antiguedad > 1 && sueldo > 100000:
		fmt.Println("Se otorga préstamo sin interés.")
	case edad > 22 && empleado && antiguedad > 1:
		fmt.Println("Se otorga préstamo con interés.")
	default:
		fmt.Println("No se otorga préstamo.")
	} */

	if edad > 22 && empleado && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Se otorga préstamo sin interés.")
		} else {
			fmt.Println("Se otorga préstamo con interés.")
		}
	} else {
		fmt.Println("No se otorga préstamo.")
	}
}

func ejercicio4(mes int) {
	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	if len(meses) < mes {
		fmt.Printf("No existe ese mes.\n")
	} else {
		fmt.Printf("%d, %s\n", mes, meses[mes-1])
	}

	//Se podría resolver con switch o de la misma forma pero con un map:
	mesesMap := map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril"}
	fmt.Printf("%d, %s\n", mes, mesesMap[mes])

	switch mes {
	case 1:
		fmt.Printf("%d, Enero\n", mes)
	case 2:
		fmt.Printf("%d, Febrero\n", mes)
	case 3:
		fmt.Printf("%d, Marzo\n", mes)
	case 4:
		fmt.Printf("%d, Abril\n", mes)
	default:
		fmt.Printf("No existe ese mes.\n")
	}
}

func ejercicio5() {
	//5-a
	alumnos := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Println("Alumnos: ", alumnos)

	//5-b
	alumnos = append(alumnos, "Gabriela")
	fmt.Println("Alumnos: ", alumnos)
}

func ejercicio6() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("Edad Benjamín: %d\n", employees["Benjamin"])

	var mayoresDe21 int
	for _, edad := range employees {
		if edad > 21 {
			mayoresDe21++
		}
	}
	fmt.Printf("Cantidad de empleados mayores de 21: %d\n", mayoresDe21)

	employees["Federico"] = 25
	fmt.Println("Empleados con Federico: ", employees)

	delete(employees, "Pedro")
	fmt.Println("Empleados sin Pedro: ", employees)
}

func main() {
	fmt.Println("--Ej. 1--")
	ejercicio1()

	fmt.Println("\n--Ej. 2--")
	ejercicio2(100.5, 30)

	fmt.Println("\n--Ej. 3--")
	ejercicio3(23, true, 2, 150000)

	fmt.Println("\n--Ej. 4--")
	ejercicio4(4)

	fmt.Println("\n--Ej. 5--")
	ejercicio5()

	fmt.Println("\n--Ej. 6--")
	ejercicio6()
}
