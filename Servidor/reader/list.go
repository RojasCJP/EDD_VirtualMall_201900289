package reader

import (
	"./dataStructures"
)

type Nodo struct {
	tienda       string
	id           int
	descripcion  string
	contacto     string
	logo         string
	calificacion int
	siguiente    *Nodo
	anterior     *Nodo
	inventario   dataStructures.AVLtree
	departamento string
}
type Lista struct {
	primero *Nodo
	ultimo  *Nodo
}

type paraInventario struct {
	Tienda       string
	Departamento string
	Calificacion int
}

func (lista *Lista) First() *Nodo {
	return lista.primero
}

func (lista *Lista) Insert(value string, id int, descripcion string, contacto string, calificacion int, logo string, inventario dataStructures.AVLtree, departamento string) {
	nodo := Nodo{value, id, descripcion, contacto, logo, calificacion, nil, nil, inventario, departamento}
	if lista.primero == nil {
		lista.primero = &nodo
		lista.ultimo = &nodo
	} else {
		nodo.siguiente = lista.primero
		lista.primero.anterior = &nodo
		lista.primero = &nodo

	}
}

func (nodo *Nodo) Siguiente() *Nodo {
	return nodo.siguiente
}

func (nodo *Nodo) Anterior() *Nodo {
	return nodo.anterior
}

func (nodo *Nodo) Value() string {
	return nodo.tienda
}

func (nodo *Nodo) IdValue() int {
	return nodo.id
}

func (lista *Lista) Show() string {
	if lista.ultimo == nil {
		return ""
	}
	var cadena string
	nodoPibot := lista.ultimo
	for nodoPibot.anterior != nil {
		//fmt.Print(nodoPibot.tienda, ", ")
		cadena += nodoPibot.tienda
		cadena += ";"
		nodoPibot = nodoPibot.anterior
	}
	cadena += nodoPibot.tienda
	//fmt.Print(nodoPibot.tienda, ", ")
	return cadena
}

func ShowArray(linealizada []Lista) string {
	var cadena string = ""
	for _, element := range linealizada {
		cadena += "["
		cadena += element.Show()
		cadena += "],"
	}
	cadena = cadena[:len(cadena)]
	return cadena
}

func (lista *Lista) Find(busqueda string) *Nodo {
	var auxiliar *Nodo
	auxiliar = lista.primero
	for auxiliar != nil {
		if auxiliar.Value() == busqueda {
			return auxiliar
		} else {
			auxiliar = auxiliar.siguiente
		}
	}
	return nil
}

func (lista *Lista) FindId(busqueda int) *Nodo {
	var auxiliar *Nodo
	auxiliar = lista.primero
	for auxiliar != nil {
		if auxiliar.IdValue() == busqueda {
			return auxiliar
		} else {
			auxiliar = auxiliar.siguiente
		}
	}
	return nil
}

func (lista *Lista) FindParaInventario(busqueda paraInventario) *Nodo {
	var auxiliar *Nodo
	auxiliar = lista.primero
	for auxiliar != nil {
		if (auxiliar.departamento == busqueda.Departamento) && (auxiliar.tienda == busqueda.Tienda) && (auxiliar.calificacion == busqueda.Calificacion) {
			return auxiliar
		} else {
			auxiliar = auxiliar.siguiente
		}
	}
	return nil
}

func (lista *Lista) ShowJson() []Tienda {
	var listaDoble []Tienda
	var auxiliar *Nodo
	auxiliar = lista.primero
	for auxiliar != nil {
		listaDoble = append(listaDoble, Tienda{auxiliar.tienda, auxiliar.descripcion, auxiliar.contacto, auxiliar.logo, auxiliar.inventario, auxiliar.departamento, auxiliar.calificacion, auxiliar.id})
		auxiliar = auxiliar.siguiente
	}
	return listaDoble
}
