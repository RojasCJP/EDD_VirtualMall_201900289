package dataStructures

import "fmt"

type NodoMatriz struct {
	valor        Cola
	departamento string
	dia          int
	up           *NodoMatriz
	down         *NodoMatriz
	left         *NodoMatriz
	right        *NodoMatriz
}

type Matriz struct {
	head *NodoMatriz
}

func (matriz *Matriz) init() {
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
	newNode := &NodoMatriz{valor: valor}
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

func (matriz *Matriz) add(row string, col int, valor Cola) {
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

func (m *Matriz) show() {
	tmprow := m.head
	for tmprow != nil {
		tmpcol := m.head
		fmt.Print(tmprow.departamento, ",", tmpcol.dia, " ")
		tmpcol = tmprow.right
		if tmprow.departamento == "0" {
			for tmpcol != nil {
				fmt.Print(tmprow.departamento, ",", tmpcol.dia, " ")
				tmpcol = tmpcol.right
			}
		} else {
			for tmpcol != nil {
				fmt.Print(tmprow.departamento, ",", getCol(tmpcol), "(", tmpcol.valor.Len, ") ")
				tmpcol = tmpcol.right
			}
		}
		tmprow = tmprow.down
		fmt.Println()
	}
}
