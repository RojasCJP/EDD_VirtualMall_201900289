package dataStructures

import (
	"testing"
)

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
