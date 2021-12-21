package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ejercicio1() {
	texto := "ID,PRECIO,CANTIDAD"
	i := 0
	for i < 10 {
		i++
		texto = fmt.Sprintf("%s\n%d,%.2f,%d", texto, 123, 350.50, 22)
	}

	err := os.WriteFile("./productos.csv", []byte(texto), 0644)

	if err != nil {
		log.Fatal(err)
		fmt.Printf("Se produjo un error al crear el archivo: %s\n", err)
		return
	}

	fmt.Printf("%s\n", texto)
}

func ejercicio2() {
	file, errFile := os.Open("productos.csv")
	if errFile != nil {
		log.Fatal(errFile)
	}
	defer file.Close()

	data, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, fila := range data {
		fmt.Printf("%s\t\t\t%s   %s\n", fila[0], fila[1], fila[2])
	}
}

func main() {
	fmt.Println("--Ej. 1--")
	ejercicio1()

	fmt.Println("\n--Ej. 2--")
	ejercicio2()
}
