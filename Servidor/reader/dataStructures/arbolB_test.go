package dataStructures

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	Nodo := _Nodo(nil)
	// Nodo := Nodo{leaf: true, n: 0}
	Usuario1 := Usuario{Dpi: 1}
	Usuario3 := Usuario{Dpi: 3}
	Usuario2 := Usuario{Dpi: 2}
	Usuario5 := Usuario{Dpi: 5}
	Usuario4 := Usuario{Dpi: 4}
	Nodo.Insert(&Usuario1)
	Nodo.Insert(&Usuario2)
	Nodo.Insert(&Usuario3)
	Nodo.Insert(&Usuario4)
	Nodo.Insert(&Usuario5)
	for i := 0; i < len(Nodo.User); i++ {
		if Nodo.User[i].Dpi != i+1 {
			t.Error("test failed")
		}
	}
}

func TestArbol(t *testing.T) {
	arbol := BTree_()
	arbol.Insert(&Usuario{1, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{2, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{3, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{4, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{5, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{6, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{7, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{8, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{9, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{10, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{11, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{12, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{13, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{14, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{15, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{16, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{17, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{18, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{19, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{20, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{21, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{22, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{23, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{24, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{25, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{26, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{27, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{28, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{29, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{30, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{31, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{32, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{33, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{34, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{35, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{36, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{37, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{38, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{39, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{40, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{41, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{42, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{43, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{44, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{45, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{46, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{47, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{48, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{49, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{50, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{51, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{52, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{53, "juan", "juan", "asdf", "asdf"})

	if arbol.Root.User[0].Dpi != 27 {
		t.Error("test failed")
	}
	if arbol.Root.Child[1].User[1].Dpi != 45 {
		t.Error("test failed")
	}
	if arbol.Root.Child[1].Child[0].User[1].Dpi != 33 {
		t.Error("test failed")
	}
	if arbol.Root.Child[0].Child[2].Child[0].User[0].Dpi != 19 {
		t.Error("test failed")
	}
}

func TestFindArbol(t *testing.T) {
	arbol := BTree_()
	arbol.Insert(&Usuario{1, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{2, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{3, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{4, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{5, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{6, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{7, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{8, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{9, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{10, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{11, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{12, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{13, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{14, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{15, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{16, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{17, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{18, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{19, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{20, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{21, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{22, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{23, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{24, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{25, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{26, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{27, "lidia", "orozco", "gaab", "gaab"})
	arbol.Insert(&Usuario{28, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{29, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{30, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{31, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{32, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{33, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{34, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{35, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{36, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{37, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{38, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{39, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{40, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{41, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{42, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{43, "juan", "rojas", "kibur", "kibur"})
	arbol.Insert(&Usuario{44, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{45, "pablo", "chinchilla", "donnie", "donnie"})
	arbol.Insert(&Usuario{46, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{47, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{48, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{49, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{50, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{51, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{52, "juan", "juan", "asdf", "asdf"})
	arbol.Insert(&Usuario{53, "juan", "juan", "asdf", "asdf"})

	if arbol.Find(27, arbol.Root).Nombre != "lidia" {
		t.Error("test failed")
	}
	if arbol.Find(27, arbol.Root).Correo != "orozco" {
		t.Error("test failed")
	}
	if arbol.Find(27, arbol.Root).Passwrd != "gaab" {
		t.Error("test failed")
	}
	if arbol.Find(27, arbol.Root).Cuenta != "gaab" {
		t.Error("test failed")
	}

	if arbol.Find(43, arbol.Root).Nombre != "juan" {
		t.Error("test failed")
	}
	if arbol.Find(43, arbol.Root).Correo != "rojas" {
		t.Error("test failed")
	}
	if arbol.Find(43, arbol.Root).Passwrd != "kibur" {
		t.Error("test failed")
	}
	if arbol.Find(43, arbol.Root).Cuenta != "kibur" {
		t.Error("test failed")
	}

	if arbol.Find(45, arbol.Root).Nombre != "pablo" {
		t.Error("test failed")
	}
	if arbol.Find(45, arbol.Root).Correo != "chinchilla" {
		t.Error("test failed")
	}
	if arbol.Find(45, arbol.Root).Passwrd != "donnie" {
		t.Error("test failed")
	}
	if arbol.Find(45, arbol.Root).Cuenta != "donnie" {
		t.Error("test failed")
	}
	fmt.Println("prueba de que si jalaron los test")
}

func TestPetarArbol(t *testing.T) {
	arbol := BTree_()
	for i := 1; i < 170; i++ {
		arbol.Insert(&Usuario{i, "juan", "juan", "asdf", "asdf"})
	}
	for i := 1; i < 170; i++ {
		if arbol.Find(i, arbol.Root) == nil {
			t.Error("Test failed")
		}
	}
	if arbol.Root.User[0].Dpi != 81 {
		t.Error("Test failed")
	}
}

func TestGrafo(t *testing.T) {
	var grafo Grafo
	grafo.AddNode("equis", []Conexion{{Nombre: "hola", Distancia: 5}, {Nombre: "adios", Distancia: 6}})
	grafo.AddNode("hola", []Conexion{{Nombre: "adios", Distancia: 8}})
	grafo.AddNode("adios", []Conexion{})
	if grafo.Nodos[0].Nombre != "equis" {
		t.Error("test failed")
	}
	if grafo.Nodos[1].Nombre != "hola" {
		t.Error("test failed")
	}
	if grafo.Nodos[2].Nombre != "adios" {
		t.Error("test failed")
	}
	if grafo.Conection("equis", "hola").Distancia != 5 {
		t.Error("test failed")
	}
	if grafo.Conection("hola", "equis").Distancia != 5 {
		t.Error("test failed")
	}
	if grafo.Conection("equis", "adios").Distancia != 6 {
		t.Error("test failed")
	}
	if grafo.Conection("adios", "equis").Distancia != 6 {
		t.Error("test failed")
	}
	if grafo.Conection("hola", "adios").Distancia != 8 {
		t.Error("test failed")
	}
	if grafo.Conection("adios", "hola").Distancia != 8 {
		t.Error("test failed")
	}
}

func TestDijkstra(t *testing.T) {
	grafo := Grafo{}
	grafo.AddNode("s", []Conexion{{"b", 4}, {"c", 2}})
	grafo.AddNode("b", []Conexion{{"c", 1}, {"d", 5}})
	grafo.AddNode("c", []Conexion{{"d", 8}, {"e", 10}})
	grafo.AddNode("d", []Conexion{{"e", 2}, {"t", 6}})
	grafo.AddNode("e", []Conexion{{"t", 2}})
	grafo.AddNode("t", []Conexion{})

	tupla := grafo.Dijkstra("s", "t")
	fmt.Println(tupla.Predecesor)
	fmt.Println(tupla.Definitiva)
	fmt.Println(tupla.Evaluar)
	fmt.Println(tupla.Distancia)
}
