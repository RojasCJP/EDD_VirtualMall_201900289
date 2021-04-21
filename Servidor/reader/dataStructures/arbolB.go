package dataStructures

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

var nodosArbolB string
var conexionesArbolB string

type Usuario struct {
	Dpi              int
	Nombre           string
	Correo           string
	Password         string
	Cuenta           string
	ContraEncriptada string
	DatosSensibles   string
}

type BTree struct {
	Root *NodoB
}

type NodoB struct {
	leaf   bool
	n      int
	Child  [6]*NodoB
	User   [5]*Usuario
	Parent *NodoB
}

func _Nodo(Parent *NodoB) *NodoB {
	return &NodoB{Parent: Parent, leaf: true, n: 0}
}

func BTree_() *BTree {
	return &BTree{Root: _Nodo(nil)}
}

func (tree *BTree) Insert(user *Usuario) {
	tree.Root = tree.Insert_(user, tree.Root)
}

func (tree *BTree) Insert_(user *Usuario, node *NodoB) *NodoB {

	if node.leaf {
		node.Insert(user)
	} else {
		found := false
		for i := 0; i < node.n-1; i++ {
			if user.Dpi < node.User[i].Dpi {
				found = true
				tree.Insert_(user, node.Child[i])
				break
			}
		}
		if !found {
			tree.Insert_(user, node.Child[node.n])
		}
	}
	if node.n == 5 {
		if node.Parent == nil {
			c := node
			node = _Nodo(nil)
			node.Insert(c.User[2])
			node.Child[0] = _Nodo(node)
			node.Child[1] = _Nodo(node)
			node.Child[0].Child[0] = c.Child[0]
			if node.Child[0].Child[0] != nil {
				node.Child[0].Child[0].Parent = node.Child[0]
				node.Child[0].leaf = false
			}
			node.Child[1].Child[0] = c.Child[3]
			if node.Child[1].Child[0] != nil {
				node.Child[1].Child[0].Parent = node.Child[1]
				node.Child[1].leaf = false
			}
			for i := 0; i < 2; i++ {
				node.Child[0].Insert(c.User[i])
				node.Child[0].Child[i+1] = c.Child[i+1]
				if node.Child[0].Child[i+1] != nil {
					node.Child[0].Child[i+1].Parent = node.Child[0]
					node.Child[0].leaf = false
				}
			}
			for i := 3; i < 5; i++ {
				node.Child[1].Insert(c.User[i])
				node.Child[1].Child[i-2] = c.Child[i+1]
				if node.Child[1].Child[i-2] != nil {
					node.Child[1].Child[i-2].Parent = node.Child[1]
					node.Child[1].leaf = false
				}
			}
			node.leaf = false
		} else {
			mkey := node.User[2]
			node.Parent.Insert(mkey)
			var index int
			for index = 0; index < node.Parent.n; index++ {
				if node.Parent.User[index] == mkey {
					break
				}
			}
			for i := node.Parent.n; i > index+1; i-- {
				if node.Parent.n < 5 {
					node.Parent.Child[i] = node.Parent.Child[i-1]
				} else {
					tree.Insert_(user, node.Parent)
				}
			}
			node.Parent.Child[index+1] = _Nodo(node.Parent)
			node.Parent.Child[index+1].Child[0] = node.Child[3]
			if node.Parent.Child[index+1].Child[0] != nil {
				node.Parent.Child[index+1].Child[0].Parent = node.Parent.Child[index+1]
				node.Parent.Child[index+1].leaf = false
			}
			for i := 3; i < 5; i++ {
				node.Parent.Child[index+1].Insert(node.User[i])
				node.Parent.Child[index+1].Child[i-2] = node.Child[i+1]
				if node.Parent.Child[index+1].Child[i-2] != nil {
					node.Parent.Child[index+1].Child[i-2].Parent = node.Parent.Child[index+1]
					node.Parent.Child[index+1].leaf = false
				}
			}
			aux := node
			node.Parent.Child[index] = _Nodo(node.Parent)
			node.Parent.Child[index].Child[0] = aux.Child[0]
			if node.Parent.Child[index].Child[0] != nil {
				node.Parent.Child[index].Child[0].Parent = node.Parent.Child[index]
				node.Parent.Child[index].leaf = false
			}
			for i := 0; i < 2; i++ {
				node.Parent.Child[index].Insert(aux.User[i])
				node.Parent.Child[index].Child[i+1] = aux.Child[i+1]
				if node.Parent.Child[index].Child[i+1] != nil {
					node.Parent.Child[index].Child[i+1].Parent = node.Parent.Child[index]
					node.Parent.Child[index].leaf = false
				}
			}
		}
	}

	return node
}

func (tree *BTree) Find(dpi int, tmp *NodoB) *Usuario {
	if tmp.Find(dpi) == nil {
		var index int
		for index = 0; index < tmp.n; index++ {
			if tmp.User[index].Dpi > dpi {
				break
			}
		}
		if tmp.Child[index] != nil {
			return tree.Find(dpi, tmp.Child[index])
		} else {
			return &Usuario{Dpi: 0, Nombre: "", Correo: "", Password: "", Cuenta: ""}
		}

	} else {
		return tmp.Find(dpi)
	}

}

func (tree *BTree) AllNodes(tmp *NodoB) {
	nodosArbolB += "\""
	for i := 0; i < tmp.n; i++ {
		if tmp.User[i] != nil {
			nodosArbolB += strconv.Itoa(tmp.User[i].Dpi)
		}
	}
	nodosArbolB += "\" [label=\""
	for i := 0; i < tmp.n; i++ {
		if tmp.User[i] != nil {
			if i < tmp.n-1 {
				nodosArbolB += strconv.Itoa(tmp.User[i].Dpi) + "|"
			} else {
				nodosArbolB += strconv.Itoa(tmp.User[i].Dpi)
			}
		}
	}
	nodosArbolB += "\"];\n"
	if tmp.Parent != nil {
		conexionesArbolB += "\""
		for i := 0; i < tmp.n; i++ {
			conexionesArbolB += strconv.Itoa(tmp.User[i].Dpi)
		}
		conexionesArbolB += "\"--\""
		for i := 0; i < tmp.Parent.n; i++ {
			conexionesArbolB += strconv.Itoa(tmp.Parent.User[i].Dpi)
		}
		conexionesArbolB += "\";\n"
	}
	for i := 0; i < len(tmp.Child); i++ {
		if tmp.Child[i] != nil {
			tree.AllNodes(tmp.Child[i])
		}
	}
}

func (tree *BTree) AllNodesEncriptado(tmp *NodoB) {
	nodosArbolB += "\""
	for i := 0; i < tmp.n; i++ {
		if tmp.User[i] != nil {
			nodosArbolB += strconv.Itoa(tmp.User[i].Dpi)
		}
	}
	nodosArbolB += "\" [label=\""
	for i := 0; i < tmp.n; i++ {
		if tmp.User[i] != nil {
			if i < tmp.n-1 {
				nodosArbolB += tmp.User[i].ContraEncriptada + "|"
			} else {
				nodosArbolB += tmp.User[i].ContraEncriptada
			}
		}
	}
	nodosArbolB += "\"];\n"
	if tmp.Parent != nil {
		conexionesArbolB += "\""
		for i := 0; i < tmp.n; i++ {
			conexionesArbolB += strconv.Itoa(tmp.User[i].Dpi)
		}
		conexionesArbolB += "\"--\""
		for i := 0; i < tmp.Parent.n; i++ {
			conexionesArbolB += strconv.Itoa(tmp.Parent.User[i].Dpi)
		}
		conexionesArbolB += "\";\n"
	}
	for i := 0; i < len(tmp.Child); i++ {
		if tmp.Child[i] != nil {
			tree.AllNodesEncriptado(tmp.Child[i])
		}
	}
}

func (tree *BTree) MakeGraphviz() {
	tree.AllNodes(tree.Root)
	cuerpoDelGrafo := "graph g {\nnode [shape=\"record\"];\ngraph [rankdir=\"BT\"];"
	cuerpoDelGrafo += nodosArbolB
	cuerpoDelGrafo += conexionesArbolB
	cuerpoDelGrafo += "}"
	err := ioutil.WriteFile("../Cliente/src/assets/graphviz/usuarios.dot", []byte(cuerpoDelGrafo), 0644)
	if err != nil {
		log.Fatal(err)
	}
	s := "dot.exe -Tpng ../Cliente/src/assets/graphviz/usuarios.dot -o ../Cliente/src/assets/arboles/usuarios.png"
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
	cuerpoDelGrafo = ""
	nodosArbolB = ""
	conexionesArbolB = ""
}

func (tree *BTree) MakeGraphvizEncriptado() {
	tree.AllNodesEncriptado(tree.Root)
	cuerpoDelGrafo := "graph g {\nnode [shape=\"record\"];\ngraph [rankdir=\"BT\"];"
	cuerpoDelGrafo += nodosArbolB
	cuerpoDelGrafo += conexionesArbolB
	cuerpoDelGrafo += "}"
	err := ioutil.WriteFile("../Cliente/src/assets/graphviz/usuariosEncriptado.dot", []byte(cuerpoDelGrafo), 0644)
	if err != nil {
		log.Fatal(err)
	}
	s := "dot.exe -Tpng ../Cliente/src/assets/graphviz/usuariosEncriptado.dot -o ../Cliente/src/assets/arboles/usuariosEncriptado.png"
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
	cuerpoDelGrafo = ""
	nodosArbolB = ""
	conexionesArbolB = ""
}

func (node *NodoB) Find(dpi int) *Usuario {
	for i := 0; i < node.n; i++ {
		if node.User[i].Dpi == dpi {
			return node.User[i]
		}
	}
	return nil
}

func (node *NodoB) Insert(user *Usuario) {
	node.n++
	for i := 0; i < node.n; i++ {
		if node.User[i] == nil {
			node.User[i] = user
			break
		} else if node.User[i].Dpi > user.Dpi {
			for j := node.n; j > i; j-- {
				node.User[j] = node.User[j-1]
			}
			node.User[i] = user
		}
	}
}
