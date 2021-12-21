package main

import "fmt"

type UsuarioEcommerce struct {
	Nombre, Apellido, Correo string
	Productos                []Producto
}

type Producto struct {
	Nombre        string
	Precio        float32
	CantidadTotal int
}

func nuevoProducto(nombre string, precio float32) Producto {
	return Producto{nombre, precio, 0}
}

func (u *UsuarioEcommerce) agregarProducto(producto *Producto, cantidad int) {
	producto.CantidadTotal += cantidad
	u.Productos = append(u.Productos, *producto)
}

func (u *UsuarioEcommerce) borrarProductos() {
	u.Productos = []Producto{}
}

func main() {
	fmt.Println("--Ej. 2--")

	usuario1 := UsuarioEcommerce{"UserN1", "UserA1", "1œ1.com", []Producto{}}
	usuario2 := UsuarioEcommerce{"UserN2", "UserA2", "2œ2.com", []Producto{}}

	monitor := nuevoProducto("monitor", 50000.5)
	tv := nuevoProducto("tv", 70000)
	teclado := nuevoProducto("teclado", 12426.3)

	fmt.Printf("Usuario1: %v\n", usuario1)
	fmt.Printf("Usuario2: %v\n", usuario2)
	fmt.Printf("Monitor: %v\n\n", monitor)

	usuario1.agregarProducto(&monitor, 1)
	usuario2.agregarProducto(&tv, 2)
	usuario1.agregarProducto(&teclado, 1)
	usuario2.agregarProducto(&teclado, 1)

	fmt.Printf("Usuario1 luego de agregar productos: %v\n", usuario1)
	fmt.Printf("Usuario2 luego de agregar productos: %v\n", usuario2)
	fmt.Printf("Monitor luego de agregar a usuarios: %v\n\n", monitor)

	usuario1.borrarProductos()

	fmt.Printf("Usuario1 luego de borrar productos: %v\n", usuario1)
	fmt.Printf("Monitor luego de borrar productos de usuario1: %v\n\n", monitor)
}
