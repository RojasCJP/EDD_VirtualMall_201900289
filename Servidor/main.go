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
	arbol := dataStructures.BTree_()

	arbol.Insert(&dataStructures.Usuario{1, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{2, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{3, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{4, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{5, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{6, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{7, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{8, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{9, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{10, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{11, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{12, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{13, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{14, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{15, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{16, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{17, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{45, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{46, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{47, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{48, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{49, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{18, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{19, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{20, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{21, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{22, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{23, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{24, "juan", "juan", "asdf", "asdf", "b", "a"})
	arbol.Insert(&dataStructures.Usuario{25, "juan", "juan", "asdf", "asdf", "b", "a"})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	reader.LevantarServer()

}
