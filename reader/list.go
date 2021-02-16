package reader

import "fmt"

type Nodo struct {
	tienda    string
	siguiente *Nodo
	anterior  *Nodo
}
type Lista struct {
	primero *Nodo
	ultimo  *Nodo
}

func (lista *Lista) First() *Nodo {
	return lista.primero
}

func (lista *Lista) Insert(value string) {
	nodo := Nodo{value, nil, nil}
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

func (lista *Lista) Show() string {
	if lista.primero == nil {
		return ""
	}
	var cadena string
	nodoPibot := lista.primero
	for nodoPibot.siguiente != nil {
		fmt.Print(nodoPibot.tienda, ", ")
		cadena += nodoPibot.tienda
		cadena += ","
		nodoPibot = nodoPibot.siguiente
	}
	cadena += nodoPibot.tienda
	fmt.Print(nodoPibot.tienda, ", ")
	return cadena
}

func ShowArray(linealizada []Lista) string {
	var cadena string = "["
	for _, element := range linealizada {
		cadena += "["
		cadena += element.Show()
		cadena += "],"
	}
	cadena= cadena[:len(cadena)]
	cadena+="]"
	return cadena
}
