package dataStructures

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type NodoMatriz struct {
	Valor        Cola
	departamento string
	dia          int
	up           *NodoMatriz
	down         *NodoMatriz
	left         *NodoMatriz
	right        *NodoMatriz
}

type Devoluciones struct {
	Departamento string
	Dia          int
}

type Matriz struct {
	head *NodoMatriz
}

func (matriz *Matriz) Init() {
	matriz.head = &NodoMatriz{departamento: "0", dia: 0}
}

func (matriz *Matriz) addRow(row string) {
	tmp := matriz.head
	if tmp.down == nil {
		newNode := &NodoMatriz{departamento: row}
		tmp.down = newNode
		newNode.up = tmp
	} else {
		for tmp.down != nil && tmp.down.departamento < row {
			tmp = tmp.down
		}
		if tmp.down == nil && tmp.departamento != row {
			newNode := &NodoMatriz{departamento: row}
			tmp.down = newNode
			newNode.up = tmp
		} else if tmp.down != nil && tmp.down.departamento != row {
			aux := tmp.down
			newNode := &NodoMatriz{departamento: row}
			tmp.down = newNode
			newNode.up = tmp
			newNode.down = aux
			aux.up = newNode
		}
	}
}

func (matriz *Matriz) addCol(col int) {
	tmp := matriz.head
	if tmp.right == nil {
		newNode := &NodoMatriz{dia: col}
		tmp.right = newNode
		newNode.left = tmp
	} else {
		for tmp.right != nil && tmp.right.dia < col {
			tmp = tmp.right
		}
		if tmp.right == nil && tmp.dia != col {
			newNode := &NodoMatriz{dia: col}
			tmp.right = newNode
			newNode.left = tmp
		} else if tmp.right != nil && tmp.right.dia != col {
			aux := tmp.right
			newNode := &NodoMatriz{dia: col}
			tmp.right = newNode
			newNode.left = tmp
			newNode.right = aux
			aux.left = newNode
		}
	}
}

func (matriz *Matriz) addNode(row string, col int, valor Cola) {
	newNode := &NodoMatriz{Valor: valor}
	tmprow := matriz.head
	tmpcol := matriz.head
	filasBajadas := 0
	for tmprow.down != nil {
		tmprow = tmprow.down
		filasBajadas++
		if tmprow.departamento == row {
			for tmpcol.right != nil {
				tmpcol = tmpcol.right
				if tmprow.right != nil {
					tmprow = tmprow.right
				}
				if tmpcol.dia == col {
					for i := 0; i < filasBajadas; i++ {
						if tmpcol.down != nil {
							tmpcol = tmpcol.down
						}
					}
					newNode.departamento = row
					newNode.dia = col
					newNode.left = tmprow
					newNode.right = tmprow.right
					newNode.up = tmpcol
					newNode.down = tmpcol.down
					newNode.left.right = newNode
					if newNode.right != nil {
						newNode.right.left = newNode
					}
					newNode.up.down = newNode
					if newNode.down != nil {
						newNode.down.up = newNode
					}
				}
			}
		}
	}
}

func (matriz *Matriz) Add(row string, col int, valor Cola) {
	matriz.addRow(row)
	matriz.addCol(col)
	matriz.addNode(row, col, valor)
}

func getCol(tmp *NodoMatriz) int {
	var col int
	for tmp.up != nil {
		tmp = tmp.up
		col = tmp.dia
	}
	return col
}

var cuerpoGraph string
var ranks string

func (m *Matriz) Show() {
	ranks = "{rank = same;"
	cuerpoGraph = "digraph matriz{\n    node [shape=box]\n    /* esto no se elimina, es para evitar el posicionamiento a lo loco */\n    e0[ shape = point, width = 0 ];\n    e1[ shape = point, width = 0 ];\n"
	tmprow := m.head
	for tmprow != nil {
		tmpcol := m.head
		//todo aqui recupero todos los departamentos con &tmprow
		cuerpoGraph += "\"" + tmprow.departamento + strconv.Itoa(tmprow.dia) + "\" [label = \"" + tmprow.departamento + "\"    width = 1.5 style = filled, fillcolor = bisque1, group = 0 ];\n"
		if tmprow.up != nil {
			cuerpoGraph += "\"" + tmprow.departamento + strconv.Itoa(tmprow.dia) + "\"->\"" + tmprow.up.departamento + strconv.Itoa(tmprow.up.dia) + "\";\n"
		}
		if tmprow.down != nil {
			cuerpoGraph += "\"" + tmprow.departamento + strconv.Itoa(tmprow.dia) + "\"->\"" + tmprow.down.departamento + strconv.Itoa(tmprow.down.dia) + "\";\n"
		}
		if tmprow.left != nil {
			cuerpoGraph += "\"" + tmprow.departamento + strconv.Itoa(tmprow.dia) + "\"->\"" + tmprow.left.departamento + strconv.Itoa(tmprow.left.dia) + "\";\n"
		}
		if tmprow.right != nil {
			cuerpoGraph += "\"" + tmprow.departamento + strconv.Itoa(tmprow.dia) + "\"->\"" + tmprow.right.departamento + strconv.Itoa(tmprow.right.dia) + "\";\n"
		}
		ranks += "\"" + tmprow.departamento + strconv.Itoa(tmprow.dia) + "\";"
		fmt.Print(tmprow.departamento, ",", tmpcol.dia, " ")
		tmpcol = tmprow.right
		if tmprow.departamento == "0" {
			for tmpcol != nil {
				//todo aqui recupero todos los dias con &tmpcol
				cuerpoGraph += "\"" + tmpcol.departamento + strconv.Itoa(tmpcol.dia) + "\" [label = \"" + strconv.Itoa(tmpcol.dia) + "\"    width = 1.5 style = filled, fillcolor = lightskyblue, group = " + strconv.Itoa(tmpcol.dia) + " ];\n"
				if tmpcol.up != nil {
					cuerpoGraph += "\"" + tmpcol.departamento + strconv.Itoa(tmpcol.dia) + "\"->\"" + tmpcol.up.departamento + strconv.Itoa(tmpcol.up.dia) + "\";\n"
				}
				if tmpcol.down != nil {
					cuerpoGraph += "\"" + tmpcol.departamento + strconv.Itoa(tmpcol.dia) + "\"->\"" + tmpcol.down.departamento + strconv.Itoa(tmpcol.down.dia) + "\";\n"
				}
				if tmpcol.left != nil {
					cuerpoGraph += "\"" + tmpcol.departamento + strconv.Itoa(tmpcol.dia) + "\"->\"" + tmpcol.left.departamento + strconv.Itoa(tmpcol.left.dia) + "\";\n"
				}
				if tmpcol.right != nil {
					cuerpoGraph += "\"" + tmpcol.departamento + strconv.Itoa(tmpcol.dia) + "\"->\"" + tmpcol.right.departamento + strconv.Itoa(tmpcol.right.dia) + "\";\n"
				}
				ranks += "\"" + tmpcol.departamento + strconv.Itoa(tmpcol.dia) + "\";"
				fmt.Print(tmprow.departamento, ",", tmpcol.dia, " ")
				tmpcol = tmpcol.right
			}
		} else {
			for tmpcol != nil {
				//todo aqui recupero todos los demas nodos
				cuerpoGraph += "\"" + tmpcol.departamento + strconv.Itoa(tmpcol.dia) + "\" [label = \"" + tmpcol.departamento + strconv.Itoa(tmpcol.dia) + "\", width=1.5 style= filled, group=" + strconv.Itoa(tmpcol.dia) + "];\n"
				if tmpcol.up != nil {
					cuerpoGraph += "\"" + tmpcol.departamento + strconv.Itoa(tmpcol.dia) + "\"->\"" + tmpcol.up.departamento + strconv.Itoa(tmpcol.up.dia) + "\";\n"
				}
				if tmpcol.down != nil {
					cuerpoGraph += "\"" + tmpcol.departamento + strconv.Itoa(tmpcol.dia) + "\"->\"" + tmpcol.down.departamento + strconv.Itoa(tmpcol.down.dia) + "\";\n"
				}
				if tmpcol.left != nil {
					cuerpoGraph += "\"" + tmpcol.departamento + strconv.Itoa(tmpcol.dia) + "\"->\"" + tmpcol.left.departamento + strconv.Itoa(tmpcol.left.dia) + "\";\n"
				}
				if tmpcol.right != nil {
					cuerpoGraph += "\"" + tmpcol.departamento + strconv.Itoa(tmpcol.dia) + "\"->\"" + tmpcol.right.departamento + strconv.Itoa(tmpcol.right.dia) + "\";\n"
				}
				ranks += "\"" + tmpcol.departamento + strconv.Itoa(tmpcol.dia) + "\";"
				fmt.Print(tmpcol.departamento, ",", tmpcol.dia, "(", tmpcol.Valor.Len, ") ")
				tmpcol = tmpcol.right
			}
		}
		ranks += "}\n"
		cuerpoGraph += ranks
		ranks = "{rank = same;"
		tmprow = tmprow.down
		fmt.Println()
	}
	cuerpoGraph += "}"
}

func (m *Matriz) Find(dia int, departamento string) *NodoMatriz {
	tmprow := m.head
	for tmprow != nil {
		tmpcol := m.head
		tmpcol = tmprow.right
		if tmprow.departamento == "0" {
			for tmpcol != nil {
				tmpcol = tmpcol.right
			}
		} else {
			for tmpcol != nil {
				if (tmpcol.departamento == departamento) && (tmpcol.dia == dia) {
					return tmpcol
				}
				tmpcol = tmpcol.right
			}
		}
		tmprow = tmprow.down
		fmt.Println()
	}
	return nil
}

func (m *Matriz) ReturnListNodes() []Devoluciones {
	var devoluciones []Devoluciones
	tmprow := m.head
	for tmprow != nil {
		tmpcol := m.head
		tmpcol = tmprow.right
		if tmprow.departamento == "0" {
			for tmpcol != nil {
				tmpcol = tmpcol.right
			}
		} else {
			for tmpcol != nil {
				devoluciones = append(devoluciones, Devoluciones{tmpcol.departamento, tmpcol.dia})
				tmpcol = tmpcol.right
			}
		}
		tmprow = tmprow.down
		fmt.Println()
	}
	return devoluciones
}

func (matriz Matriz) MakeGraph() {
	matriz.Show()
	err := ioutil.WriteFile("../Cliente/src/assets/graphviz/matriz.dot", []byte(cuerpoGraph), 0644)
	if err != nil {
		log.Fatal(err)
	}
	s := "dot.exe -Tpng ../Cliente/src/assets/graphviz/matriz.dot -o ../Cliente/src/assets/arboles/matriz.png"
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
