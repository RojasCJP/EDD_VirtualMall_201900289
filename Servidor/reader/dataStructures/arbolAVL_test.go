package dataStructures

import (
	"fmt"
	"testing"
)

func TestAVLtree_Find(t *testing.T) {
	tree := AVLtree{}
	tree.Add(Producto{Codigo: 5, Nombre: "hola", Precio: 2.5, Descripcion: "hola", Cantidad: 35})
	tree.Add(Producto{Codigo: 10, Nombre: "que", Precio: 5, Descripcion: "que", Cantidad: 30})
	tree.Add(Producto{Codigo: 15, Nombre: "tal", Precio: 7.5, Descripcion: "tal", Cantidad: 25})
	tree.Add(Producto{Codigo: 20, Nombre: "como", Precio: 10, Descripcion: "como", Cantidad: 20})
	tree.Add(Producto{Codigo: 25, Nombre: "estas", Precio: 12.5, Descripcion: "estas", Cantidad: 15})
	tree.Add(Producto{Codigo: 30, Nombre: "espero", Precio: 15, Descripcion: "espero", Cantidad: 10})
	tree.Add(Producto{Codigo: 35, Nombre: "bien", Precio: 17.5, Descripcion: "bien", Cantidad: 5})
	nodoResultado := tree.Find(10)
	if nodoResultado.Valor.Nombre != "que" {
		t.Error("test not passed")
	}
	if nodoResultado.Valor.Precio != 5 {
		t.Error("test not passed")
	}
	if nodoResultado.Valor.Descripcion != "que" {
		t.Error("test not passed")
	}
	if nodoResultado.Valor.Cantidad != 30 {
		t.Error("test not passed")
	}
	nodoResultado.Valor.Descripcion = "hola como estas"
	nuevoNodo := tree.Find(10)
	fmt.Println(nuevoNodo)
}

func TestPreorder(t *testing.T) {
	tree := AVLtree{}
	tree.Add(Producto{Codigo: 5, Nombre: "hola", Precio: 2.5, Descripcion: "hola", Cantidad: 35})
	tree.Add(Producto{Codigo: 10, Nombre: "que", Precio: 5, Descripcion: "que", Cantidad: 30})
	tree.Add(Producto{Codigo: 15, Nombre: "tal", Precio: 7.5, Descripcion: "tal", Cantidad: 25})
	tree.Add(Producto{Codigo: 20, Nombre: "como", Precio: 10, Descripcion: "como", Cantidad: 20})
	tree.Add(Producto{Codigo: 25, Nombre: "estas", Precio: 12.5, Descripcion: "estas", Cantidad: 15})
	tree.Add(Producto{Codigo: 30, Nombre: "espero", Precio: 15, Descripcion: "espero", Cantidad: 10})
	tree.Add(Producto{Codigo: 35, Nombre: "bien", Precio: 17.5, Descripcion: "bien", Cantidad: 5})
	tree.preorder(tree.Root)
	//tree.MakeGraphviz(tree.Root)
	if preorder != "20 10 5 15 30 25 35 " {
		t.Error("test not passed")
	}
}
func TestPostorder(t *testing.T) {
	tree := AVLtree{}
	tree.Add(Producto{Codigo: 5})
	tree.Add(Producto{Codigo: 10})
	tree.Add(Producto{Codigo: 15})
	tree.Add(Producto{Codigo: 20})
	tree.Add(Producto{Codigo: 25})
	tree.Add(Producto{Codigo: 30})
	tree.Add(Producto{Codigo: 35})
	tree.postorder(tree.Root)
	if postorder != "5 15 10 25 35 30 20 " {
		t.Error("test not passed")
	}
}
func TestInorder(t *testing.T) {
	tree := AVLtree{}
	tree.Add(Producto{Codigo: 5})
	tree.Add(Producto{Codigo: 10})
	tree.Add(Producto{Codigo: 15})
	tree.Add(Producto{Codigo: 20})
	tree.Add(Producto{Codigo: 25})
	tree.Add(Producto{Codigo: 30})
	tree.Add(Producto{Codigo: 35})
	tree.inorder(tree.Root)
	if inorder != "5 10 15 20 25 30 35 " {
		t.Error("test not passed")
	}
}

func TestShow(t *testing.T) {
	matriz := Matriz{}
	matriz.init()
	matriz.add("ahola", 1, Cola{Len: 1})
	matriz.add("ahola", 2, Cola{Len: 2})
	matriz.add("ahola", 3, Cola{Len: 3})
	matriz.add("bhola", 1, Cola{Len: 4})
	matriz.add("bhola", 2, Cola{Len: 5})
	matriz.add("bhola", 3, Cola{Len: 6})
	matriz.add("chola", 1, Cola{Len: 7})
	matriz.add("chola", 2, Cola{Len: 8})
	matriz.add("chola", 3, Cola{Len: 9})
	matriz.show()
}
