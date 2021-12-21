package main

import "fmt"

// CLASE 16/12

func ejercicio1() {
	var (
		nombre    = "Paloma"
		direccion = "Pje Rep 2024"
	)

	fmt.Printf("Nombre: %s - Dirección: %s\n", nombre, direccion)
}

func ejercicio2() {
	var temp, humedad, presion float32 = 35.0, 10.0, 3.6
	fmt.Printf("Temperatura: %.1f - Humedad: %.1f - Presión: %.1f\n", temp, humedad, presion)
}

func ejercicio3() {
	var lnombre string
	var lapellido string //cambio nombre a lapellido pq no se usaba
	var edad int

	lapellido = "Apellido"        //= un string no un 6
	var licenciaDeConducir = true //el nombre de la var estaba mal y faltaba el tipo?
	var estaturaDeLaPersona int   //el nobmre estaba mal
	cantidadDeHijos := 2

	// Agrego valores a las que no estaban inicializadas
	lnombre = "Alumno"
	edad = 16
	estaturaDeLaPersona = 2

	// Agrego prints así *usa* las variables y no da error:
	fmt.Println("Nombre: ", lnombre)
	fmt.Println("Apellido: ", lapellido)
	fmt.Println("Edad: ", edad)
	fmt.Println("Tiene licencia: ", licenciaDeConducir)
	fmt.Println("Estatura: ", estaturaDeLaPersona)
	fmt.Println("Cantidad de hijos: ", cantidadDeHijos)
}

func ejercicio4() {
	var apellido string = "Gomez"
	var edad int = 35             //int es sin ""
	boolean := false              //sin ; al final y false sin ""
	var sueldo float32 = 45857.90 //debería ser float, no string
	var nombre string = "Julián"

	fmt.Println("apellido: ", apellido)
	fmt.Println("edad: ", edad)
	fmt.Println("boolean: ", boolean)
	fmt.Println("sueldo: ", sueldo)
	fmt.Println("nombre: ", nombre)
}

func main() {
	fmt.Println("--Ej. 1--")
	ejercicio1()

	fmt.Println("\n--Ej. 2--")
	ejercicio2()

	fmt.Println("\n--Ej. 3--")
	ejercicio3()

	fmt.Println("\n--Ej. 4--")
	ejercicio4()
}
