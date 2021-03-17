package reader

import (
	"./dataStructures"
	"sort"
)

type InventariosData struct {
	//todo tengo que ver si le van a poner inventarios o invetarios
	Inventarios []Inventario
}

type Inventario struct {
	Tienda       string
	Departamento string
	Calificacion int
	Productos    []dataStructures.Producto
}

type DevolucionInventario struct {
	Nodo  *Nodo
	Lista []dataStructures.Producto
}

type multiSorter struct {
	changes []Tienda
	less    []lessFunc
}

func (ms *multiSorter) Swap(i, j int) {
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}
func (ms *multiSorter) Len() int {
	return len(ms.changes)
}
func (ms *multiSorter) Sort(changes []Tienda) {
	ms.changes = changes
	sort.Sort(ms)
}
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.changes[i], &ms.changes[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

type lessFunc func(p1, p2 *Tienda) bool

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func MakeMatrix(dataJson Datos) [][][][]Tienda {

	var matrix [][][][]Tienda
	var sub3 [][][]Tienda
	var sub2 [][]Tienda
	var sub1 []Tienda

	for i := 0; i < len(dataJson.Datos); i++ {
		matrix = append(matrix, sub3)
	}

	for i := 0; i < len(dataJson.Datos); i++ {
		for j := 0; j < len(dataJson.Datos[i].Departamentos); j++ {
			matrix[i] = append(matrix[i], sub2)
		}
	}

	for i := 0; i < len(dataJson.Datos); i++ {
		for j := 0; j < len(dataJson.Datos[i].Departamentos); j++ {
			for k := 0; k < 5; k++ {
				matrix[i][j] = append(matrix[i][j], sub1)
			}
		}
	}

	for i := 0; i < len(dataJson.Datos); i++ {
		for j := 0; j < len(dataJson.Datos[i].Departamentos); j++ {
			for k := 0; k < len(dataJson.Datos[i].Departamentos[j].Tiendas); k++ {
				for l := 0; l < 5; l++ {
					if dataJson.Datos[i].Departamentos[j].Tiendas[k].Calificacion == (l + 1) {
						tienda := dataJson.Datos[i].Departamentos[j].Tiendas[k]
						matrix[i][j][l] = append(matrix[i][j][l], Tienda{tienda.Nombre, tienda.Descripcion, tienda.Contacto, tienda.Logo, tienda.Inventario, dataJson.Datos[i].Departamentos[j].Nombre, tienda.Calificacion, tienda.Id})
					}
				}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			for k := 0; k < len(matrix[i][j]); k++ {
				nombre := func(c1, c2 *Tienda) bool {
					return c1.Nombre < c2.Nombre
				}
				OrderedBy(nombre).Sort(matrix[i][j][k])
			}
		}
	}

	//fmt.Println(matrix)
	//fmt.Println(matrix[0])
	//fmt.Println(matrix[0][0])
	//fmt.Println(matrix[0][0][0])
	return matrix
}

func Linealizar(matrix [][][][]Tienda, dataJson Datos) []Lista {
	var linealizada []Lista
	var id int
	id = 0
	for i := 0; i < (len(matrix)); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			for k := 0; k < len(matrix[i][j]); k++ {
				linealizada = append(linealizada, Lista{})
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			for k := 0; k < len(matrix[i][j]); k++ {
				lista := Lista{}
				for l := 0; l < len(matrix[i][j][k]); l++ {
					tiendaAux := matrix[i][j][k][l]
					lista.Insert(tiendaAux.Nombre, id, tiendaAux.Descripcion, tiendaAux.Contacto, tiendaAux.Calificacion, tiendaAux.Logo, tiendaAux.Inventario, tiendaAux.Departamento)
					id++
				}
				//fmt.Println(len(matrix[i][j]))
				linealizada[i*(len(matrix[i])*len(matrix[i][j]))+j*(len(matrix[i][j]))+k] = lista
			}
		}
	}
	return linealizada
}

func FindWithId(id int, linealizada *[]Lista) *Nodo {
	for i := 0; i < len(*linealizada); i++ {
		nodo := (*linealizada)[i].FindId(id)
		if nodo != nil {
			return nodo
		}
	}
	return nil
}

func FindParaInventario(data InventariosData, linealizada *[]Lista) []DevolucionInventario {
	var tiendaBusqueda paraInventario
	var devolucion []DevolucionInventario
	for i := 0; i < len(data.Inventarios); i++ {
		tiendaBusqueda = paraInventario{data.Inventarios[i].Tienda, data.Inventarios[i].Departamento, data.Inventarios[i].Calificacion}
		for j := 0; j < len(*linealizada); j++ {
			nodo := (*linealizada)[j].FindParaInventario(tiendaBusqueda)
			if nodo != nil {
				devolucion = append(devolucion, DevolucionInventario{Nodo: nodo, Lista: data.Inventarios[i].Productos})
			}
		}
	}
	return devolucion
}
