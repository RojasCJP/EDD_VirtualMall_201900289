package dataStructures

import "math"

type Grafo struct {
	Nodos []NodoGrafo
}

type NodoGrafo struct {
	Nombre     string
	Conexiones []Conexion
}

type Conexion struct {
	Nombre    string
	Distancia int
}

type TuplaCamino struct {
	Definitiva bool
	Evaluar    bool
	Distancia  int
	Predecesor string
}

func (grafo *Grafo) AddNode(name string, conections []Conexion) {
	grafo.Nodos = append(grafo.Nodos, NodoGrafo{name, conections})
	for i := 0; i < len(grafo.Nodos); i++ {
		for j := 0; j < len(conections); j++ {
			if grafo.Nodos[i].Nombre == conections[j].Nombre {
				grafo.Nodos[i].Conexiones = append(grafo.Nodos[i].Conexiones, conections[j])
			}
		}
		for j := 0; j < len(grafo.Nodos[i].Conexiones); j++ {
			if grafo.Nodos[i].Conexiones[j].Nombre == name {
				if grafo.Nodos[i].Nombre != name {
					grafo.Nodos[len(grafo.Nodos)-1].Conexiones = append(grafo.Nodos[len(grafo.Nodos)-1].Conexiones, Conexion{grafo.Nodos[i].Nombre, grafo.Nodos[i].Conexiones[j].Distancia})
				}
			}
		}

	}
}

func (grafo *Grafo) Conection(from string, to string) Conexion {
	for i := 0; i < len(grafo.Nodos); i++ {
		for j := 0; j < len(grafo.Nodos[i].Conexiones); j++ {
			if grafo.Nodos[i].Nombre == from && grafo.Nodos[i].Conexiones[j].Nombre == to {
				return grafo.Nodos[i].Conexiones[j]
			}
		}
	}
	return Conexion{}
}

func (grafo *Grafo) Dijkstra(from string, to string) TuplaCamino {
	//todo este me tiene que regresar todos los nodos en una lista
	var recorrido [][]TuplaCamino
	nodoInvalido := Conexion{}
	nodoUso := from
	tuplaUso := TuplaCamino{true, false, math.MaxInt64, from}
	for i := 0; i < len(grafo.Nodos); i++ {
		recorrido = append(recorrido, []TuplaCamino{})
	}
	for i := 0; i < len(grafo.Nodos); i++ {
		for j := 0; j < len(grafo.Nodos); j++ {
			recorrido[i] = append(recorrido[i], TuplaCamino{Evaluar: true, Definitiva: false, Distancia: math.MaxInt64})
		}
	}
	for i := 0; i < len(grafo.Nodos); i++ {
		for j := 0; j < len(grafo.Nodos); j++ {
			if i-1 >= 0 {
				if recorrido[i-1][j].Evaluar == false {
					recorrido[i][j].Evaluar = false
				}
			}
			if recorrido[i][j].Evaluar == true {
				if nodoUso != grafo.Nodos[j].Nombre {
					conexion := grafo.Conection(nodoUso, grafo.Nodos[j].Nombre)
					if conexion != nodoInvalido {
						if i-1 >= 0 {
							if conexion.Distancia+tuplaUso.Distancia < recorrido[i-1][j].Distancia {
								recorrido[i][j].Predecesor = nodoUso
								recorrido[i][j].Evaluar = true
								recorrido[i][j].Distancia = conexion.Distancia + tuplaUso.Distancia
								recorrido[i][j].Definitiva = false
							} else {
								recorrido[i][j].Predecesor = recorrido[i-1][j].Predecesor
								recorrido[i][j].Evaluar = true
								recorrido[i][j].Distancia = recorrido[i-1][j].Distancia
								recorrido[i][j].Definitiva = false
							}
						} else {
							recorrido[i][j].Predecesor = tuplaUso.Predecesor
							recorrido[i][j].Evaluar = true
							recorrido[i][j].Distancia = conexion.Distancia
							recorrido[i][j].Definitiva = false
						}
					} else {
						if i-1 >= 0 {
							recorrido[i][j].Predecesor = recorrido[i-1][j].Predecesor
							recorrido[i][j].Evaluar = true
							recorrido[i][j].Distancia = recorrido[i-1][j].Distancia
							recorrido[i][j].Definitiva = false
						}
					}
				} else {
					if i-1 >= 0 {
						recorrido[i][j].Predecesor = recorrido[i-1][j].Predecesor
						recorrido[i][j].Evaluar = false
						recorrido[i][j].Distancia = recorrido[i-1][j].Distancia
						recorrido[i][j].Definitiva = true
					} else {
						recorrido[i][j].Predecesor = nodoUso
						recorrido[i][j].Evaluar = false
						recorrido[i][j].Distancia = 0
						recorrido[i][j].Definitiva = true
					}
				}
			}
		}
		tuplaUso.Distancia = math.MaxInt64
		for j := 0; j < len(grafo.Nodos); j++ {
			if recorrido[i][j].Distancia < tuplaUso.Distancia && recorrido[i][j].Definitiva == false {
				tuplaUso = recorrido[i][j]
				nodoUso = grafo.Nodos[j].Nombre
			}
		}
		if nodoUso == to {
			return tuplaUso
		}
	}
	return TuplaCamino{}
}
