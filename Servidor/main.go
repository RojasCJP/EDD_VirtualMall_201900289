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
	grafo.AddNode("equis", []dataStructures.Conexion{{Nombre: "hola", Distancia: 6}})
	grafo.AddNode("hola", []dataStructures.Conexion{{Nombre: "adios", Distancia: 8}})
	grafo.AddNode("adios", []dataStructures.Conexion{})
	grafo.Dijkstra("equis", "adios")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	reader.LevantarServer()
}
