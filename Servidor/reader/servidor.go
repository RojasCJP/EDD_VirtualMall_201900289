package reader

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var JsonData = Datos{}

type htmltemplate struct {
	Name   string
	Carnet int
	Json   string
}

type TiendaTransicional struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
	Logo         string
}

type getArregloTemplate struct {
	Cadena    string
	Direccion string
}

func SetJsonData(jsonData Datos) {
	JsonData = jsonData

}

func LevantarServer() {
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/", rutaInicial).Methods("GET")
	router.HandleFunc("/cargartienda", CargarJson).Methods("POST")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", tiendaEspecifica).Methods("POST")
	router.HandleFunc("/id/{numero}", idTienda).Methods("GET")
	router.HandleFunc("/Eliminar", eliminarTienda).Methods("DELETE")
	router.HandleFunc("/imagen", imagenSubida)
	router.HandleFunc("/guardar", GuardarRoutes)
	router.HandleFunc("/todasTiendas", todasTiendas).Methods("GET")
	//todo tengo que hacer una para todas las tiendas, ahorita la voy a hacer quemada
	http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router))
}

func todasTiendas(response http.ResponseWriter, request *http.Request) {
	fmt.Println("se accedio a todasTiendas")
	page1 := TiendaTransicional{"Gerardo Weco", "Hola Gerardo", "gerardoWeco@gmail.com", 69, "https://scontent.fgua3-2.fna.fbcdn.net/v/t1.0-9/88339887_2840382472710717_3169115093359132672_o.jpg?_nc_cat=107&ccb=1-3&_nc_sid=09cbfe&_nc_ohc=UJjNuNTyouoAX-SAP0G&_nc_ht=scontent.fgua3-2.fna&oh=c2ae32f06c0c7522c5cd90b60b134608&oe=606E33FF"}
	page2 := TiendaTransicional{"Gerardo Hueco", "Como estas", "holaquetal@gmail.com", 69, "https://scontent.fgua3-2.fna.fbcdn.net/v/t1.0-9/88339887_2840382472710717_3169115093359132672_o.jpg?_nc_cat=107&ccb=1-3&_nc_sid=09cbfe&_nc_ohc=UJjNuNTyouoAX-SAP0G&_nc_ht=scontent.fgua3-2.fna&oh=c2ae32f06c0c7522c5cd90b60b134608&oe=606E33FF"}
	page3 := TiendaTransicional{"Gerardo Gar", "Gerardo es gay", "gerardoWeco@gmail.com", 69, "https://scontent.fgua3-2.fna.fbcdn.net/v/t1.0-9/88339887_2840382472710717_3169115093359132672_o.jpg?_nc_cat=107&ccb=1-3&_nc_sid=09cbfe&_nc_ohc=UJjNuNTyouoAX-SAP0G&_nc_ht=scontent.fgua3-2.fna&oh=c2ae32f06c0c7522c5cd90b60b134608&oe=606E33FF"}
	page4 := TiendaTransicional{"Gerardo Homosexual", "Le gusta Edson", "gerardoWeco@gmail.com", 69, "https://scontent.fgua3-2.fna.fbcdn.net/v/t1.0-9/88339887_2840382472710717_3169115093359132672_o.jpg?_nc_cat=107&ccb=1-3&_nc_sid=09cbfe&_nc_ohc=UJjNuNTyouoAX-SAP0G&_nc_ht=scontent.fgua3-2.fna&oh=c2ae32f06c0c7522c5cd90b60b134608&oe=606E33FF"}
	page5 := TiendaTransicional{"Gerardo Wapo", "Gerardo weco", "gerardoWeco@gmail.com", 69, "https://scontent.fgua3-2.fna.fbcdn.net/v/t1.0-9/88339887_2840382472710717_3169115093359132672_o.jpg?_nc_cat=107&ccb=1-3&_nc_sid=09cbfe&_nc_ohc=UJjNuNTyouoAX-SAP0G&_nc_ht=scontent.fgua3-2.fna&oh=c2ae32f06c0c7522c5cd90b60b134608&oe=606E33FF"}
	var pages []TiendaTransicional
	pages = append(pages, page1)
	pages = append(pages, page2)
	pages = append(pages, page3)
	pages = append(pages, page4)
	pages = append(pages, page5)
	data, err := json.Marshal(pages)
	if err != nil {
		response.Write([]byte("ocurrio un error"))
	}
	response.Write(data)
	//response.Write([]byte("Recuerda que lo primero que debes hacer es cargar tu archivo "))

}

func imagenSubida(response http.ResponseWriter, request *http.Request) {
	imagen, err := ioutil.ReadFile("templates/imagen.png")
	if err != nil {
		response.Write([]byte("al cargar hubo error"))
	}
	response.Write(imagen)
}

func rutaInicial(response http.ResponseWriter, request *http.Request) {
	page := htmltemplate{"Juan Pablo Rojas Chinchilla", 201900289, "EDD 1er Semestre 2021"}
	temp, _ := template.ParseFiles("./templates/welcome-template.html")
	//response.Write([]byte("Recuerda que lo primero que debes hacer es cargar tu archivo "))
	temp.Execute(response, page)
	fmt.Println(JsonData)
}

func CargarJson(response http.ResponseWriter, request *http.Request) {
	fmt.Println("EDD primer semestre  2021")
	data, errRead := ioutil.ReadAll(request.Body)
	if errRead != nil {
		response.Write([]byte("error en la carga del json"))
	}
	var mainJson = Datos{}
	err := json.Unmarshal(data, &mainJson)
	if err != nil {
		log.Fatal("error al convertir a estructura " + err.Error())
	}
	JsonData = mainJson
	response.Write(data)
	//fmt.Println(string(data))
	//fmt.Println(JsonData)
	SetJsonData(JsonData)
	MakeMatrix(JsonData)
}

func getArreglo(response http.ResponseWriter, request *http.Request) {
	matrix := MakeMatrix(JsonData)
	linealizada := Linealizar(matrix, JsonData)
	paraEnviar := ShowArray(linealizada[:])
	GraphvizMethod(paraEnviar)
	fmt.Println(paraEnviar)
	//data, err2 := json.Marshal(paraEnviar)
	//if err2 != nil {
	//	log.Fatal("error al imprimir los datos" + err2.Error())
	//}
	//fmt.Println(data)
	//response.Write(data)
	page := getArregloTemplate{"Pagina de Get Arreglo", "./templates/imagen.png"}
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	temp, _ := template.ParseFiles("./templates/getArreglo.html")
	//response.Write([]byte("Recuerda que lo primero que debes hacer es cargar tu archivo "))
	temp.Execute(response, page)
	fmt.Println(JsonData)
}

func tiendaEspecifica(response http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Write([]byte("entrada no valida"))
	}
	var especifica Especifica
	err2 := json.Unmarshal(data, &especifica)
	if err2 != nil {
		response.Write([]byte("entrada no valida"))
	}
	tienda := FindTienda(especifica, JsonData)
	salida, err3 := json.Marshal(tienda)
	if err3 != nil {
		response.Write([]byte("error en la conversion de json"))
	}
	fmt.Println(data)
	response.Write(salida)

}

func idTienda(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	matrix := MakeMatrix(JsonData)
	numero, err := strconv.Atoi(vars["numero"])
	if err != nil {
		response.Write([]byte("el id que ingreso es invalido"))
	}
	linealizada := Linealizar(matrix, JsonData)
	nombre := linealizada[numero-1]
	tienda := nombre.ShowJson()
	salida, err3 := json.Marshal(tienda)
	if err3 != nil {
		response.Write([]byte("error en la conversion de json"))
	}
	response.Write(salida)
}

func eliminarTienda(response http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Write([]byte("entrada no valida"))
	}
	var especifica Especifica
	err2 := json.Unmarshal(data, &especifica)
	if err2 != nil {
		response.Write([]byte("entrada no valida"))
	}
	var mensaje = ""
	JsonData, mensaje = DeleteTienda(especifica, JsonData)
	salida, err3 := json.Marshal(JsonData)
	if err3 != nil {
		response.Write([]byte("error en la conversion de json"))
	}
	fmt.Println(JsonData)
	if mensaje == "" {
		response.Write(salida)
	} else {
		response.Write([]byte(mensaje))
	}
}

func GuardarRoutes(response http.ResponseWriter, request *http.Request) {
	GuardarArchivo()
	response.Write([]byte("Json guardado"))
}

func FindTienda(especifica Especifica, data Datos) Tienda {
	auxiliar := Tienda{}
	for i := 0; i < len(data.Datos); i++ {
		for j := 0; j < len(data.Datos[i].Departamentos); j++ {
			for k := 0; k < len(data.Datos[i].Departamentos[j].Tiendas); k++ {
				tienda := data.Datos[i].Departamentos[j].Tiendas[k]
				departamento := data.Datos[i].Departamentos[j]
				if (tienda.Nombre == especifica.Nombre) && (tienda.Calificacion == especifica.Calificacion) && (departamento.Nombre == especifica.Departamento) {
					auxiliar.Nombre = tienda.Nombre
					auxiliar.Calificacion = tienda.Calificacion
					auxiliar.Descripcion = tienda.Descripcion
					auxiliar.Contacto = tienda.Contacto
					return auxiliar
				}
			}
		}
	}

	return Tienda{"Su tienda no se encuentra", "Ingrese una tienda valida", "Algun dato no es correcto", 0}
}

func DeleteTienda(especifica Especifica, data Datos) (Datos, string) {
	//auxiliar := Tienda{}
	for i := 0; i < len(data.Datos); i++ {
		for j := 0; j < len(data.Datos[i].Departamentos); j++ {
			for k := 0; k < len(data.Datos[i].Departamentos[j].Tiendas); k++ {
				tienda := data.Datos[i].Departamentos[j].Tiendas[k]
				departamento := data.Datos[i].Departamentos[j]
				if (tienda.Nombre == especifica.Nombre) && (tienda.Calificacion == especifica.Calificacion) && (departamento.Nombre == especifica.Departamento) {
					data.Datos[i].Departamentos[j].Tiendas = append(data.Datos[i].Departamentos[j].Tiendas[:k], data.Datos[i].Departamentos[j].Tiendas[k+1:]...)
					return data, ""
				}
			}
		}
	}

	return data, "no se encontro ninguna tienda con estos datos"
}

func FindTiendaWithNombre(nombre string, data Datos) Tienda { //todo me va a servir para meter los datos a la matriz
	auxiliar := Tienda{}
	for i := 0; i < len(data.Datos); i++ {
		for j := 0; j < len(data.Datos[i].Departamentos); j++ {
			for k := 0; k < len(data.Datos[i].Departamentos[j].Tiendas); k++ {
				tienda := data.Datos[i].Departamentos[j].Tiendas[k]
				if tienda.Nombre == nombre {
					auxiliar.Nombre = tienda.Nombre
					auxiliar.Calificacion = tienda.Calificacion
					auxiliar.Descripcion = tienda.Descripcion
					auxiliar.Contacto = tienda.Contacto
					return auxiliar
				}
			}
		}
	}

	return Tienda{"Su tienda no se encuentra", "Ingrese una tienda valida", "Algun dato no es correcto", 0}
}

func GuardarArchivo() {
	data, err := json.Marshal(JsonData)
	if err != nil {
		log.Fatal(err)
	}
	err1 := ioutil.WriteFile("templates/datosGuardados.json", data, 0644)
	if err1 != nil {
		log.Fatal(err1)
	}
}

type Especifica struct {
	Departamento string
	Nombre       string
	Calificacion int
}
