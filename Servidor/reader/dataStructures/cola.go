package dataStructures

import "strconv"

type codigo struct {
	Codigo int
}

type ValorCola struct {
	Fecha        string
	Tienda       string
	Departamento string
	Calificacion int
	Productos    []codigo
}

type Pedidos struct {
	Pedidos []ValorCola
}

type NodoCola struct {
	valor     ValorCola
	anterior  *NodoCola
	siguiente *NodoCola
}

type Cola struct {
	first *NodoCola
	last  *NodoCola
	Len   int
}

func (valorCola ValorCola) Dia() int {
	dia := valorCola.Fecha
	diaSub := string(dia[0:1])
	diaInt, _ := strconv.Atoi(diaSub)
	return diaInt
}
func (valorCola ValorCola) Mes() int {
	mes := valorCola.Fecha
	mesSub := string(mes[3:4])
	mesInt, _ := strconv.Atoi(mesSub)
	return mesInt
}
func (valorCola ValorCola) Year() int {
	year := valorCola.Fecha
	yearSub := string(year[6:9])
	yearInt, _ := strconv.Atoi(yearSub)
	return yearInt
}
