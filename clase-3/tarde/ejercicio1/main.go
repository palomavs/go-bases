package main

import "fmt"

type Usuario struct {
	Nombre, Apellido, Correo, Contraseña string
	Edad                                 int
}

func (u *Usuario) cambiarNombre(nombre, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}

func (u *Usuario) cambiarEdad(edad int) {
	u.Edad = edad
}

func (u *Usuario) cambiarCorreo(correo string) {
	u.Correo = correo
}

func (u *Usuario) cambiarContraseña(contraseña string) {
	u.Contraseña = contraseña
}

func main() {
	fmt.Println("--Ej. 1--")

	usuarioPrueba := Usuario{"Paloma", "Selci", "p@p.com", "123456", 22}

	fmt.Printf("Usuario: %+v\n", usuarioPrueba)

	usuarioPrueba.cambiarNombre("Palo", "Vazquez")
	usuarioPrueba.cambiarEdad(23)
	usuarioPrueba.cambiarCorreo("vs@vs.com")
	usuarioPrueba.cambiarContraseña("abcdefg")

	fmt.Printf("Usuario luego de los cambios: %+v\n\n", usuarioPrueba)
}
