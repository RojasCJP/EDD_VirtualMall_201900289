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
	Valor     ValorCola
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
	diaSub := string(dia[0:2])
	diaInt, _ := strconv.Atoi(diaSub)
	return diaInt
}
func (valorCola ValorCola) Mes() int {
	mes := valorCola.Fecha
	mesSub := string(mes[3:5])
	mesInt, _ := strconv.Atoi(mesSub)
	return mesInt
}
func (valorCola ValorCola) Year() int {
	year := valorCola.Fecha
	yearSub := string(year[6:10])
	yearInt, _ := strconv.Atoi(yearSub)
	return yearInt
}

func (cola *Cola) Add(nodoCola NodoCola) {
	if cola.first == nil {
		cola.first = &nodoCola
		cola.last = &nodoCola
	} else {
		nodoCola.siguiente = cola.first
		cola.first.anterior = &nodoCola
		cola.first = &nodoCola

	}
	cola.Len++
}

func (cola *Cola) Pop() {
	//todo tengo que hacer este metodo
}

func (cola *Cola) AllProducts() []int {
	var productos []int
	auxiliar := cola.first
	final := false
	for !final {
		for i := 0; i < len(auxiliar.Valor.Productos); i++ {
			productos = append(productos, auxiliar.Valor.Productos[i].Codigo)
		}
		if auxiliar == cola.last {
			final = true
		} else {
			auxiliar = auxiliar.siguiente
		}
	}
	return productos
}
