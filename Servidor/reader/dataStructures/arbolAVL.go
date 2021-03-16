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
	Nombre      string
	Codigo      int
	Descripcion string
	Precio      float32
	Cantidad    int
}

type NodoAVL struct {
	valor     Producto
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
		return &NodoAVL{valor: valor}
	} else if valor.Codigo < tmp.valor.Codigo {
		tmp.izquierdo = tree._add(valor, tmp.izquierdo)
		if (tree.altura(tmp.izquierdo) - tree.altura(tmp.derecho)) == 2 {
			if valor.Codigo < tmp.izquierdo.valor.Codigo {
				tmp = tree.srl(tmp)
			} else {
				tmp = tree.drl(tmp)
			}
		}
	} else if valor.Codigo > tmp.valor.Codigo {
		tmp.derecho = tree._add(valor, tmp.derecho)
		if tree.altura(tmp.derecho)-tree.altura(tmp.izquierdo) == 2 {
			if valor.Codigo > tmp.derecho.valor.Codigo {
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
	// todo lo tenia diferente espino, en lugar de derecho tenia izquierdo, hay que comprobar
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
		preorder += strconv.Itoa(tmp.valor.Codigo) + " "
		tree.preorder(tmp.izquierdo)
		tree.preorder(tmp.derecho)
	}
}

var inorder string

func (tree AVLtree) inorder(tmp *NodoAVL) {
	if tmp != nil {
		tree.inorder(tmp.izquierdo)
		inorder += strconv.Itoa(tmp.valor.Codigo) + " "
		tree.inorder(tmp.derecho)
	}
}

var postorder string

func (tree AVLtree) postorder(tmp *NodoAVL) {
	if tmp != nil {
		tree.postorder(tmp.izquierdo)
		tree.postorder(tmp.derecho)
		postorder += strconv.Itoa(tmp.valor.Codigo) + " "
	}
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
		cuerpo += "\"" + strconv.Itoa(tmp.valor.Codigo) + "\"" + " [label=\"" + strconv.Itoa(tmp.valor.Codigo) + "|{" + tmp.valor.Nombre + "|" + strconv.Itoa(tmp.valor.Cantidad) + "}|" + fmt.Sprintf("%.2f", tmp.valor.Precio) + "\"]\n"
		if tmp.izquierdo != nil {
			cuerpo += "\"" + strconv.Itoa(tmp.valor.Codigo) + "\"->\"" + strconv.Itoa(tmp.izquierdo.valor.Codigo) + "\"\n"
		}
		if tmp.derecho != nil {
			cuerpo += "\"" + strconv.Itoa(tmp.valor.Codigo) + "\"->\"" + strconv.Itoa(tmp.derecho.valor.Codigo) + "\"\n"
		}
		tree.MakeCuerpo(tmp.derecho)
	}
}
