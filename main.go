package main

import (
	"./reader"
	"encoding/json"
	"fmt"
	"log"
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
	reader.MakeMatrix(mainJson)
	reader.LevantarServer(mainJson)
}
