package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float32
	Cantidad int
}

type Servicio struct {
	Nombre          string
	Precio, minutos float32
}

type Mantenimiento struct {
	Nombre string
	Precio float32
}

func sumarProductos(productos []Producto, p chan float32) {
	var total float32
	fmt.Println("Arrancó sumar PROD")

	for _, producto := range productos {
		total +=
			producto.Precio * float32(producto.Cantidad)
	}
	fmt.Println("Terminó sumar PROD")
	p <- total
}

func sumarServicios(servicios []Servicio, s chan float32) {
	var total float32
	fmt.Println("Arrancó sumar SS")

	for _, servicio := range servicios {
		if servicio.minutos < 30 {
			total += servicio.Precio
		} else {
			total += servicio.minutos / 30 * servicio.Precio
		}
	}
	fmt.Println("Terminó sumar SS")
	s <- total
}

func sumarMantenimiento(mantenimientos []Mantenimiento, m chan float32) {
	var total float32
	fmt.Println("Arrancó sumar MANT")

	for _, mantenimiento := range mantenimientos {
		total += mantenimiento.Precio
	}
	fmt.Println("Terminó sumar MANT")
	m <- total
}

func main() {
	var p, s, m = make(chan float32), make(chan float32), make(chan float32)

	go sumarProductos([]Producto{{"Mesa", 100.23, 12}, {"Silla", 50.65, 44}}, p)
	go sumarServicios([]Servicio{{"Arreglo de aire", 120.45, 63}}, s)
	go sumarMantenimiento([]Mantenimiento{{"Aire", 300.54}, {"Computadora", 200.65}}, m)

	sumaTotal := <-p + <-s + <-m
	fmt.Printf("\nMonto final: %.2f\n\n", sumaTotal)
}
