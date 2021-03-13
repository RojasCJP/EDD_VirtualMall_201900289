package dataStructures

import "strconv"

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
	root *NodoAVL
}

func (tree *AVLtree) add(valor Producto) {
	tree.root = tree._add(valor, tree.root)
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
