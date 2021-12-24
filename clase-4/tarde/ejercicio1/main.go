package main

import (
	"fmt"
	"os"
)

func recuperar() {
	errPanic := recover()
	if errPanic != nil {
		fmt.Println(errPanic)
	}
}

func leerArchivo(dir string) {
	file, err := os.Open(dir)
	defer file.Close()

	defer recuperar()

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	fmt.Println(file)
}

func main() {
	leerArchivo("customers.txt")

	fmt.Printf("\nEjecución finalizada.\n\n")
}
