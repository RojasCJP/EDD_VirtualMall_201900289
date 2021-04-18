package dataStructures

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os/exec"
	"strconv"
	"strings"
)

var cuerpoGraphGrafo string

type Grafo struct {
	Nodos                []NodoGrafo
	PosicionInicialRobot string
	Entrega              string
}

type NodoGrafo struct {
	Nombre  string
	Enlaces []Conexion
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
				grafo.Nodos[i].Enlaces = append(grafo.Nodos[i].Enlaces, conections[j])
			}
		}
		for j := 0; j < len(grafo.Nodos[i].Enlaces); j++ {
			if grafo.Nodos[i].Enlaces[j].Nombre == name {
				if grafo.Nodos[i].Nombre != name {
					grafo.Nodos[len(grafo.Nodos)-1].Enlaces = append(grafo.Nodos[len(grafo.Nodos)-1].Enlaces, Conexion{grafo.Nodos[i].Nombre, grafo.Nodos[i].Enlaces[j].Distancia})
				}
			}
		}

	}
}

func (grafo *Grafo) Conection(from string, to string) Conexion {
	for i := 0; i < len(grafo.Nodos); i++ {
		for j := 0; j < len(grafo.Nodos[i].Enlaces); j++ {
			if grafo.Nodos[i].Nombre == from && grafo.Nodos[i].Enlaces[j].Nombre == to {
				return grafo.Nodos[i].Enlaces[j]
			}
		}
	}
	return Conexion{}
}

func (grafo *Grafo) Dijkstra(from string, to string) TuplaCamino {
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

type CaminosPrecio struct {
	Nombre     string
	Conexiones []Conexion
}

type Caminos struct {
	Nombre     string
	Conexiones string
}

func (grafo *Grafo) EncontrarCamino(from string, to string) ([]Caminos, []CaminosPrecio) {
	var anterior string = to
	var camino []Caminos
	var grafoCompleto []CaminosPrecio
	var tupla TuplaCamino
	for i := 0; i < len(grafo.Nodos); i++ {
		grafoCompleto = append(grafoCompleto, CaminosPrecio{Nombre: grafo.Nodos[i].Nombre, Conexiones: grafo.Nodos[i].Enlaces})
	}
	for anterior != from {
		tupla = grafo.Dijkstra(from, anterior)
		camino = append(camino, Caminos{Nombre: anterior, Conexiones: tupla.Predecesor})
		anterior = tupla.Predecesor
	}
	return camino, grafoCompleto
}

func (grafo *Grafo) EncontrarCaminoVariosNodos(nodos []string) ([]Caminos, []CaminosPrecio) {
	var caminos []Caminos
	var grafoCompleto []CaminosPrecio
	for i := 0; i < len(nodos)-1; i++ {
		caminos1, grafoCompleto1 := grafo.EncontrarCamino(nodos[i], nodos[i+1])
		grafoCompleto = grafoCompleto1
		for j := 0; j < len(caminos1); j++ {
			caminos = append(caminos, caminos1[j])
		}
	}
	return caminos, grafoCompleto
}

func (grafo *Grafo) MakeGraphGrafo(conexiones []string) {
	camino, completo := grafo.EncontrarCaminoVariosNodos(conexiones)
	var conexionesHechas []string
	cuerpoGraphGrafo += "graph grafo{\n"
	cuerpoGraphGrafo += conexiones[0] + " [color=\"red\"];\n"
	cuerpoGraphGrafo += conexiones[len(conexiones)-1] + " [color=\"green\"];\n"
	for i := 0; i < len(completo); i++ {
		conexionesHechas = append(conexionesHechas, completo[i].Nombre)
		for j := 0; j < len(completo[i].Conexiones); j++ {
			if inConexiones(completo[i].Conexiones[j].Nombre, conexionesHechas) {
				cuerpoGraphGrafo += completo[i].Nombre + "--" + completo[i].Conexiones[j].Nombre + " [label=\"" + strconv.Itoa(completo[i].Conexiones[j].Distancia) + "\"];\n"
			}
		}
	}
	for i := 0; i < len(camino); i++ {
		cuerpoGraphGrafo += camino[i].Nombre + "--" + camino[i].Conexiones + ";\n"
	}
	cuerpoGraphGrafo += "}"
	fmt.Println(cuerpoGraphGrafo)
}

func (grafo Grafo) MakeFileGrafo(camino []string) {
	grafo.MakeGraphGrafo(camino)
	err := ioutil.WriteFile("../Cliente/src/assets/graphviz/grafo.dot", []byte(cuerpoGraphGrafo), 0644)
	if err != nil {
		log.Fatal(err)
	}
	s := "dot.exe -Tpng ../Cliente/src/assets/graphviz/grafo.dot -o ../Cliente/src/assets/arboles/grafo.png"
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	err1 := cmd.Start()
	if err1 != nil {
		log.Printf("Command finishes with error: %v", err1)
	}
	err1 = cmd.Wait()
	if err1 != nil {
		log.Printf("Command finishes with error: %v", err1)
	}
	cuerpoGraphGrafo = ""
}

func inConexiones(peticion string, conexiones []string) bool {
	for i := 0; i < len(conexiones); i++ {
		if peticion == conexiones[i] {
			return true
		}
	}
	return false
}
