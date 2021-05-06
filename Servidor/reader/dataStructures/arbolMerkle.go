package dataStructures

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os/exec"
	"strings"
)

type Prueba struct {
	numero int
}

type NodoMerkle struct {
	izquierda *NodoMerkle
	derecha   *NodoMerkle
	padre     *NodoMerkle
	valor     string
}

type MerkleTree struct {
	Root *NodoMerkle
}

func InitMerkle() *MerkleTree {
	tree := MerkleTree{}
	tree.Root = &NodoMerkle{padre: nil}
	return &tree
}

func (tree *MerkleTree) FullTree(list []string) {
	size := len(list)
	sizeExpected := 0
	exponent := 0
	for sizeExpected < size {
		sizeExpected = int(math.Pow(2, float64(exponent)))
		exponent++
	}

	var lista []*NodoMerkle
	tree.CreateNodes(tree.Root, int(math.Log2(float64(sizeExpected))), &lista)
	for i := 0; i < len(lista); i++ {
		if i < size {
			lista[i].valor = fmt.Sprintf("%v", sha256.Sum256([]byte(fmt.Sprintf("%v", list[i]))))
		} else {
			lista[i].valor = fmt.Sprintf("%v", sha256.Sum256([]byte("vacio")))
		}
	}
	tree.SumarHijos(tree.Root)
}

func (tree *MerkleTree) CreateNodes(node *NodoMerkle, contador int, allNodes *[]*NodoMerkle) {
	node.izquierda = &NodoMerkle{padre: node}
	node.derecha = &NodoMerkle{padre: node}
	contador--
	if contador >= 0 {
		tree.CreateNodes(node.izquierda, contador, allNodes)
		tree.CreateNodes(node.derecha, contador, allNodes)
	}
	if contador == -1 {
		*allNodes = append(*allNodes, node)
	}
}

func (tree *MerkleTree) SumarHijos(node *NodoMerkle) {
	for node.izquierda.valor == "" || node.derecha.valor == "" {
		tree.SumarHijos(node.izquierda)
		tree.SumarHijos(node.derecha)
	}
	if node.izquierda.valor != "" && node.derecha.valor != "" {
		node.valor = node.izquierda.valor + node.derecha.valor
	}
}

func (node *NodoMerkle) AddValue(value string) {
	node.valor = value
}

var cuerpoMerkle string

func (tree *MerkleTree) makeCuerpo(nodo *NodoMerkle) {
	if nodo != nil {
		if nodo.valor != "" {
			tree.makeCuerpo(nodo.izquierda)
			tree.makeCuerpo(nodo.derecha)
			cuerpoMerkle += "\"" + nodo.valor + "\"" + " [label=\"" + nodo.valor + "\"]\n"
			if nodo.izquierda != nil {
				if nodo.izquierda.valor != "" {
					cuerpoMerkle += "\"" + nodo.valor + "\"->\"" + nodo.izquierda.valor + "\"\n"
				}
			}
			if nodo.derecha != nil {
				if nodo.derecha.valor != "" {
					cuerpoMerkle += "\"" + nodo.valor + "\"->\"" + nodo.derecha.valor + "\"\n"
				}
			}
		}
	}
}

func (tree *MerkleTree) MakeGraph(name string) {
	cuerpoMerkle = "digraph arbol{\n  node [shape=record]\n"
	tree.makeCuerpo(tree.Root)
	cuerpoMerkle += "}"
	err := ioutil.WriteFile("../Cliente/src/assets/graphviz/"+name+".dot", []byte(cuerpoMerkle), 0644)
	if err != nil {
		log.Fatal(err)
	}
	s := "dot.exe -Tpng ../Cliente/src/assets/graphviz/" + name + ".dot -o ../Cliente/src/assets/arboles/" + name + ".png"
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
