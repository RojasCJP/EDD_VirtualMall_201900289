package reader

type Tienda struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Logo         string
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
