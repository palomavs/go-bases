package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	pequeño = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

type Tienda struct {
	Productos []Producto
}

type producto struct {
	Tipo, Nombre string
	Precio       float32
}

func (p producto) CalcularCostoTotal() (float32, error) {
	switch p.Tipo {
	case pequeño:
		return p.Precio, nil
	case mediano:
		return (p.Precio) * 1.03, nil
	case grande:
		return (p.Precio)*1.06 + 2500, nil
	default:
		return 0, errors.New("no existe el tipo especificado")
	}
}

type Producto interface {
	CalcularCostoTotal() (float32, error)
}

type Ecommerce interface {
	Total() float32
	Agregar(Producto)
}

func (t Tienda) Total() float32 {
	var costoTotal float32

	for _, producto := range t.Productos {
		costoProducto, err := producto.CalcularCostoTotal()
		if err != nil {
			log.Fatal(err)
			return 0
		}
		costoTotal += costoProducto
	}
	return costoTotal
}

func (t *Tienda) Agregar(producto Producto) {
	t.Productos = append(t.Productos, producto)
}

func nuevoProducto(tipo, nombre string, precio float32) Producto {
	producto := producto{tipo, nombre, precio}
	var prod Producto = producto
	return prod
}

func nuevaTienda(productos []Producto) Ecommerce {
	tienda := Tienda{productos}
	var ecommerce Ecommerce = &tienda
	return ecommerce
}

func main() {
	fmt.Println("--Ej. 3--")

	producto1 := nuevoProducto(pequeño, "vaso", 12.45)
	producto2 := nuevoProducto(grande, "mes", 134.5)

	var productos = []Producto{producto1, producto2}

	tienda := nuevaTienda(productos)
	fmt.Printf("Tienda: %+v\n\n", tienda)

	tienda.Agregar(nuevoProducto(mediano, "silla", 55.67))
	fmt.Printf("Tienda luego de agregar: %+v\n\n", tienda)
	fmt.Printf("CostoTotal de la tienda: $%.2f\n\n", tienda.Total())
}
