package dataStructures

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Producto struct {
	Nombre         string
	Codigo         int
	Descripcion    string
	Precio         float32
	Cantidad       int
	Imagen         string
	Almacenamiento string
}

type NodoAVL struct {
	Valor     Producto
	izquierdo *NodoAVL
	derecho   *NodoAVL
	altura    int
}

type AVLtree struct {
	Root *NodoAVL
}

func (tree *AVLtree) Add(valor Producto) {
	tree.Root = tree._add(valor, tree.Root)
}

func (tree AVLtree) _add(valor Producto, tmp *NodoAVL) *NodoAVL {
	if tmp == nil {
		return &NodoAVL{Valor: valor}
	} else if valor.Codigo < tmp.Valor.Codigo {
		tmp.izquierdo = tree._add(valor, tmp.izquierdo)
		if (tree.altura(tmp.izquierdo) - tree.altura(tmp.derecho)) == 2 {
			if valor.Codigo < tmp.izquierdo.Valor.Codigo {
				tmp = tree.srl(tmp)
			} else {
				tmp = tree.drl(tmp)
			}
		}
	} else if valor.Codigo > tmp.Valor.Codigo {
		tmp.derecho = tree._add(valor, tmp.derecho)
		if tree.altura(tmp.derecho)-tree.altura(tmp.izquierdo) == 2 {
			if valor.Codigo > tmp.derecho.Valor.Codigo {
				tmp = tree.srr(tmp)
			} else {
				tmp = tree.drr(tmp)
			}
		}
	}
	d := tree.altura(tmp.derecho)
	i := tree.altura(tmp.izquierdo)
	m := tree.maxi(d, i)
	tmp.altura = m + 1
	return tmp
}

func (tree AVLtree) Find(codigo int, tmp *NodoAVL) *NodoAVL {
	if tmp != nil {
		if codigo < tmp.Valor.Codigo {
			return tree.Find(codigo, tmp.izquierdo)
		} else if codigo > tmp.Valor.Codigo {
			return tree.Find(codigo, tmp.derecho)
		} else if codigo == tmp.Valor.Codigo {
			return tmp
		}
	}
	return tmp
}

func (tree AVLtree) srl(tmp1 *NodoAVL) *NodoAVL {
	tmp2 := tmp1.izquierdo
	tmp1.izquierdo = tmp2.derecho
	tmp2.derecho = tmp1
	tmp1.altura = tree.maxi(tree.altura(tmp1.izquierdo), tree.altura(tmp1.derecho)) + 1
	tmp2.altura = tree.maxi(tree.altura(tmp2.izquierdo), tmp1.altura) + 1
	return tmp2
}

func (tree AVLtree) drl(tmp *NodoAVL) *NodoAVL {
	tmp.izquierdo = tree.srr(tmp.izquierdo)
	return tree.srl(tmp)
}

func (tree AVLtree) srr(tmp1 *NodoAVL) *NodoAVL {
	tmp2 := tmp1.derecho
	tmp1.derecho = tmp2.izquierdo
	tmp2.izquierdo = tmp1
	tmp1.altura = tree.maxi(tree.altura(tmp1.izquierdo), tree.altura(tmp1.derecho)) + 1
	tmp2.altura = tree.maxi(tree.altura(tmp2.derecho), tmp1.altura) + 1
	return tmp2
}

func (tree AVLtree) drr(tmp *NodoAVL) *NodoAVL {
	tmp.derecho = tree.srl(tmp.derecho)
	return tree.srr(tmp)
}

func (tree AVLtree) maxi(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	} else {
		return val2
	}
}

func (tree AVLtree) altura(tmp *NodoAVL) int {
	if tmp == nil {
		return -1
	} else {
		return tmp.altura
	}
}

var preorder string

func (tree AVLtree) preorder(tmp *NodoAVL) {
	if tmp != nil {
		preorder += strconv.Itoa(tmp.Valor.Codigo) + " "
		tree.preorder(tmp.izquierdo)
		tree.preorder(tmp.derecho)
	}
}

var inorder string

func (tree AVLtree) inorder(tmp *NodoAVL) {
	if tmp != nil {
		tree.inorder(tmp.izquierdo)
		inorder += strconv.Itoa(tmp.Valor.Codigo) + " "
		tree.inorder(tmp.derecho)
	}
}

var postorder string

func (tree AVLtree) postorder(tmp *NodoAVL) {
	if tmp != nil {
		tree.postorder(tmp.izquierdo)
		tree.postorder(tmp.derecho)
		postorder += strconv.Itoa(tmp.Valor.Codigo) + " "
	}
}

var ListElements []Producto

func (tree AVLtree) ListAllProducts(tmp *NodoAVL) {
	if tmp != nil {
		tree.ListAllProducts(tmp.izquierdo)
		tree.ListAllProducts(tmp.derecho)
		ListElements = append(ListElements, tmp.Valor)
	}
}
func (tree AVLtree) ClearList() {
	ListElements = nil
}

var cuerpo string

func (tree AVLtree) MakeGraphviz(tmp *NodoAVL) {
	cuerpo = "digraph arbol{\n  node [shape=record]\n"
	tree.MakeCuerpo(tmp)
	cuerpo += "}"
	err := ioutil.WriteFile("../Cliente/src/assets/graphviz/inventario.dot", []byte(cuerpo), 0644)
	if err != nil {
		log.Fatal(err)
	}
	s := "dot.exe -Tpng ../Cliente/src/assets/graphviz/inventario.dot -o ../Cliente/src/assets/arboles/inventario.png"
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
}

func (tree AVLtree) MakeCuerpo(tmp *NodoAVL) {
	if tmp != nil {
		tree.MakeCuerpo(tmp.izquierdo)
		cuerpo += "\"" + strconv.Itoa(tmp.Valor.Codigo) + "\"" + " [label=\"" + strconv.Itoa(tmp.Valor.Codigo) + "|{" + tmp.Valor.Nombre + "|" + strconv.Itoa(tmp.Valor.Cantidad) + "}|" + fmt.Sprintf("%.2f", tmp.Valor.Precio) + "\"]\n"
		if tmp.izquierdo != nil {
			cuerpo += "\"" + strconv.Itoa(tmp.Valor.Codigo) + "\"->\"" + strconv.Itoa(tmp.izquierdo.Valor.Codigo) + "\"\n"
		}
		if tmp.derecho != nil {
			cuerpo += "\"" + strconv.Itoa(tmp.Valor.Codigo) + "\"->\"" + strconv.Itoa(tmp.derecho.Valor.Codigo) + "\"\n"
		}
		tree.MakeCuerpo(tmp.derecho)
	}
}
