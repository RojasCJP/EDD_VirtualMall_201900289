package reader

import (
	"sort"
)

func MakeMatrix(dataJson Datos) [][][][]string {

	var matrix [][][][]string
	var sub3 [][][]string
	var sub2 [][]string
	var sub1 []string

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
						matrix[i][j][l] = append(matrix[i][j][l], dataJson.Datos[i].Departamentos[j].Tiendas[k].Nombre)
					}
				}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			for k := 0; k < len(matrix[i][j]); k++ {
				sort.Strings(matrix[i][j][k])
			}
		}
	}

	//fmt.Println(matrix)
	//fmt.Println(matrix[0])
	//fmt.Println(matrix[0][0])
	//fmt.Println(matrix[0][0][0])
	return matrix
}

func Linealizar(matrix [][][][]string, dataJson Datos) []Lista {
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
					tiendaAux := FindTiendaWithNombre(matrix[i][j][k][l], dataJson)
					lista.Insert(tiendaAux.Nombre, id, tiendaAux.Descripcion, tiendaAux.Contacto, tiendaAux.Calificacion)
					id++
				}
				//fmt.Println(len(matrix[i][j]))
				linealizada[i*(len(matrix[i])*len(matrix[i][j]))+j*(len(matrix[i][j]))+k] = lista
			}
		}
	}
	return linealizada
}

func FindWithId(id int, linealizada []Lista) string {
	var nombre string
	for i := 0; i < len(linealizada); i++ {
		nodo := linealizada[i].FindId(id)
		if nodo != nil {
			return nodo.Value()
		}
	}
	return nombre
}