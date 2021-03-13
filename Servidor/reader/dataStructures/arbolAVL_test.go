package dataStructures

import (
	"testing"
)

func TestPreorder(t *testing.T) {
	tree := AVLtree{}
	tree.add(Producto{Codigo: 5})
	tree.add(Producto{Codigo: 10})
	tree.add(Producto{Codigo: 15})
	tree.add(Producto{Codigo: 20})
	tree.add(Producto{Codigo: 25})
	tree.add(Producto{Codigo: 30})
	tree.add(Producto{Codigo: 35})
	tree.preorder(tree.root)
	if preorder != "20 10 5 15 30 25 35 " {
		t.Error("test not passed")
	}
}
func TestPostorder(t *testing.T) {
	tree := AVLtree{}
	tree.add(Producto{Codigo: 5})
	tree.add(Producto{Codigo: 10})
	tree.add(Producto{Codigo: 15})
	tree.add(Producto{Codigo: 20})
	tree.add(Producto{Codigo: 25})
	tree.add(Producto{Codigo: 30})
	tree.add(Producto{Codigo: 35})
	tree.postorder(tree.root)
	if postorder != "5 15 10 25 35 30 20 " {
		t.Error("test not passed")
	}
}
func TestInorder(t *testing.T) {
	tree := AVLtree{}
	tree.add(Producto{Codigo: 5})
	tree.add(Producto{Codigo: 10})
	tree.add(Producto{Codigo: 15})
	tree.add(Producto{Codigo: 20})
	tree.add(Producto{Codigo: 25})
	tree.add(Producto{Codigo: 30})
	tree.add(Producto{Codigo: 35})
	tree.inorder(tree.root)
	if inorder != "5 10 15 20 25 30 35 " {
		t.Error("test not passed")
	}
}
