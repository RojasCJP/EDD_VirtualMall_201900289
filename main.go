package main

import (
	"./reader"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("EDD primer semestre  2021")
	var data = []byte(`{"Datos":[{"Indice":"A","Departamentos":[{"Nombre":"Deportes","Tiendas":[{"Nombre":"Aurora","Descripcion":"Es una empresa multinacional estadounidense dedicada al diseño, desarrollo, fabricación y comercialización de equipamiento deportivo: balones, calzado, ropa, equipo, accesorios y otros artículos deportivos","Contacto":"5544-3377","Calificacion":5},{"Nombre":"Amador","Descripcion":"es una empresa alemana fabricante de accesorios, ropa y calzado deportivo, cuya sede central está en Herzogenaurach, Alemania","Contacto":"5588-9988","Calificacion":4},{"Nombre":"Armados","Descripcion":"Equipo extremo","Contacto":"8995222","Calificacion":5}]},{"Nombre":"Comida","Tiendas":[{"Nombre":"A comer todo","Descripcion":"todo lo que puedan pedir por un dollar","Contacto":"559999","Calificacion":5}]},{"Nombre":"celulares","Tiendas":[]}]},{"Indice":"B","Departamentos":[{"Nombre":"Deportes","Tiendas":[]},{"Nombre":"Comida","Tiendas":[]},{"Nombre":"Celulares","Tiendas":[{"Nombre":"Bayoneta","Descripcion":"Telefonos militares","Contacto":"bayoneta@gmail.com","Calificacion":5}]}]}]}`)
	var mainJson = reader.Datos{}
	err := json.Unmarshal(data, &mainJson)
	if err != nil {
		log.Fatal("error al convertir a estructura " + err.Error())
	}
	fmt.Println(mainJson)
	makeMatrix(mainJson)
	levantarServer(mainJson)
}

func levantarServer(json reader.Datos) {
	http.HandleFunc("/", rutaInicial)
	http.HandleFunc("/getArreglo", getArreglo)
	http.HandleFunc("/TiendaEspecifica", tiendaEspecifica)
	http.HandleFunc("/id", idTienda)
	http.HandleFunc("/Eliminar", eliminarTienda)
	http.ListenAndServe(":3000", nil)
}

func rutaInicial(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("hello world"))
}
func getArreglo(response http.ResponseWriter, request *http.Request) {
	var data = []byte(`{"Datos":[{"Indice":"A","Departamentos":[{"Nombre":"Deportes","Tiendas":[{"Nombre":"Aurora","Descripcion":"Es una empresa multinacional estadounidense dedicada al diseño, desarrollo, fabricación y comercialización de equipamiento deportivo: balones, calzado, ropa, equipo, accesorios y otros artículos deportivos","Contacto":"5544-3377","Calificacion":5},{"Nombre":"Amador","Descripcion":"es una empresa alemana fabricante de accesorios, ropa y calzado deportivo, cuya sede central está en Herzogenaurach, Alemania","Contacto":"5588-9988","Calificacion":4},{"Nombre":"Armados","Descripcion":"Equipo extremo","Contacto":"8995222","Calificacion":5}]},{"Nombre":"Comida","Tiendas":[{"Nombre":"A comer todo","Descripcion":"todo lo que puedan pedir por un dollar","Contacto":"559999","Calificacion":5}]},{"Nombre":"celulares","Tiendas":[]}]},{"Indice":"B","Departamentos":[{"Nombre":"Deportes","Tiendas":[]},{"Nombre":"Comida","Tiendas":[]},{"Nombre":"Celulares","Tiendas":[{"Nombre":"Bayoneta","Descripcion":"Telefonos militares","Contacto":"bayoneta@gmail.com","Calificacion":5}]}]}]}`)
	var mainJson = reader.Datos{}
	err := json.Unmarshal(data, &mainJson)
	if err != nil {
		log.Fatal("error al convertir estructura de datos " + err.Error())
	}
	matrix := makeMatrix(mainJson)
	linealizada := linealizar(matrix)
	paraEnviar := reader.ShowArray(linealizada[:])
	fmt.Println(paraEnviar)
	//todo tengo que ver que onda con la linealizada por que fijo esta mal porque no puede imprimir ciertos punteros
	data, err2 := json.Marshal(paraEnviar)
	if err2 != nil {
		log.Fatal("error al imprimir los datos" + err.Error())
	}
	fmt.Println(data)
	response.Write(data)
}
func tiendaEspecifica(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Pagina de tiendaEspecifica"))
}
func idTienda(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Pagina de idTienda"))
}
func eliminarTienda(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Pagina de eliminarTienda"))
}

func makeMatrix(dataJson reader.Datos) [][][][]string {

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

	//fmt.Println(matrix)
	//fmt.Println(matrix[0])
	//fmt.Println(matrix[0][0])
	//fmt.Println(matrix[0][0][0])
	return matrix
}

func linealizar(matrix [][][][]string) []reader.Lista {
	var linealizada []reader.Lista

	for i := 0; i < (len(matrix)); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			for k := 0; k < len(matrix[i][j]); k++ {
				linealizada = append(linealizada, reader.Lista{})
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			for k := 0; k < len(matrix[i][j]); k++ {
				lista := reader.Lista{}
				for l := 0; l < len(matrix[i][j][k]); l++ {
					lista.Insert(matrix[i][j][k][l])
				}
				//todo tengo que revisar si tengo que multiplicar por la longitud o solo sumarla
				//fmt.Println(len(matrix[i][j]))
				linealizada[i*(len(matrix[i])*len(matrix[i][j]))+j*(len(matrix[i][j]))+k] = lista
			}
		}
	}
	return linealizada
}
