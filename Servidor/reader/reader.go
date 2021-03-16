package reader

import (
	"./dataStructures"
)

type Tienda struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Logo         string
	Inventario   dataStructures.AVLtree
	Calificacion int
	Id           int
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
