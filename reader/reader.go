package reader

import "fmt"

type Tienda struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
}

type Departamento struct {
	Nombre  string
	Tiendas []Tienda
}

type Indice struct {
	Indice        string
	Departamentos []Departamento
}

type Datos struct {
	Datos []Indice
}

func Hola() {
	fmt.Println("hola que tal como estas")
}
