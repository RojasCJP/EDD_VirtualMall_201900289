package main

import (
	"./reader"
	"./reader/dataStructures"
	"net/http"
)

func main() {
	//matriz := dataStructures.Matriz{}
	//matriz.Init()
	//matriz.Add("chola", 2, dataStructures.Cola{Len: 8})
	//matriz.Add("bhola", 3, dataStructures.Cola{Len: 6})
	//matriz.Add("ahola", 2, dataStructures.Cola{Len: 2})
	//matriz.Add("ahola", 3, dataStructures.Cola{Len: 3})
	//matriz.Add("bhola", 2, dataStructures.Cola{Len: 5})
	//matriz.Add("bhola", 1, dataStructures.Cola{Len: 4})
	//matriz.Add("ahola", 1, dataStructures.Cola{Len: 1})
	//matriz.Add("chola", 1, dataStructures.Cola{Len: 7})
	//matriz.Add("chola", 3, dataStructures.Cola{Len: 9})
	//matriz.Show()
	grafo := dataStructures.Grafo{}
	grafo.AddNode("a", []dataStructures.Conexion{{Nombre: "b", Distancia: 2}, {Nombre: "c", Distancia: 3}})
	grafo.AddNode("b", []dataStructures.Conexion{{Nombre: "d", Distancia: 5}, {Nombre: "e", Distancia: 2}})
	grafo.AddNode("c", []dataStructures.Conexion{{Nombre: "e", Distancia: 5}})
	grafo.AddNode("d", []dataStructures.Conexion{{Nombre: "e", Distancia: 1}, {Nombre: "z", Distancia: 2}})
	grafo.AddNode("e", []dataStructures.Conexion{{Nombre: "z", Distancia: 4}})
	grafo.AddNode("z", []dataStructures.Conexion{})

	grafo.Dijkstra("a", "c")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	reader.LevantarServer()
}
