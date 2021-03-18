package reader

import (
	"./dataStructures"
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
var imagenPredeterminada string = "https://es.jumpseller.com/images/learn/choosing-platform/laptop.jpg"
var arregloListas []Lista
var Carrito []ElementoCarrito
var Years []dataStructures.Year

//todo tengo que hacer lo del calendario, armar la estructura
//tengo que verificar si el a;o existe ya
//tengo que ver si el mes ya existe
//tengo que ver si en la matriz ya existe el departamento y dia
//si existe tengo que agregar a la cola los productos
//si no existe tengo que hacer la cola y agregar los productos
//por ultimo tengo que graficar

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

type ElementoCarrito struct {
	IdTienda       int
	CodigoProducto int
	Cantidad       int
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
	router.HandleFunc("/tiendaUnica/{numero}", findTiendaUnica).Methods("GET")
	router.HandleFunc("/listaInventario/{numero}", getListaInventario).Methods("GET")
	router.HandleFunc("/inventario/{numero}", showInventario).Methods("GET")
	router.HandleFunc("/inventario", meterElementos).Methods("POST")
	router.HandleFunc("/Eliminar", eliminarTienda).Methods("DELETE")
	router.HandleFunc("/imagen", imagenSubida)
	router.HandleFunc("/guardar", GuardarRoutes)
	router.HandleFunc("/todasTiendas", todasTiendas).Methods("GET")
	router.HandleFunc("/addCarrito", addCarrito).Methods("POST")
	http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router))
}

func todasTiendas(response http.ResponseWriter, request *http.Request) {
	fmt.Println("se accedio a todasTiendas")
	//page1 := TiendaTransicional{"Gerardo Weco", "Hola Gerardo", "gerardoWeco@gmail.com", 69, "https://scontent.fgua3-2.fna.fbcdn.net/v/t1.0-9/88339887_2840382472710717_3169115093359132672_o.jpg?_nc_cat=107&ccb=1-3&_nc_sid=09cbfe&_nc_ohc=UJjNuNTyouoAX-SAP0G&_nc_ht=scontent.fgua3-2.fna&oh=c2ae32f06c0c7522c5cd90b60b134608&oe=606E33FF"}
	//page2 := TiendaTransicional{"Gerardo Hueco", "Como estas", "holaquetal@gmail.com", 69, "https://scontent.fgua3-2.fna.fbcdn.net/v/t1.0-9/88339887_2840382472710717_3169115093359132672_o.jpg?_nc_cat=107&ccb=1-3&_nc_sid=09cbfe&_nc_ohc=UJjNuNTyouoAX-SAP0G&_nc_ht=scontent.fgua3-2.fna&oh=c2ae32f06c0c7522c5cd90b60b134608&oe=606E33FF"}
	//page3 := TiendaTransicional{"Gerardo Gar", "Gerardo es gay", "gerardoWeco@gmail.com", 69, "https://scontent.fgua3-2.fna.fbcdn.net/v/t1.0-9/88339887_2840382472710717_3169115093359132672_o.jpg?_nc_cat=107&ccb=1-3&_nc_sid=09cbfe&_nc_ohc=UJjNuNTyouoAX-SAP0G&_nc_ht=scontent.fgua3-2.fna&oh=c2ae32f06c0c7522c5cd90b60b134608&oe=606E33FF"}
	//page4 := TiendaTransicional{"Gerardo Homosexual", "Le gusta Edson", "gerardoWeco@gmail.com", 69, "https://scontent.fgua3-2.fna.fbcdn.net/v/t1.0-9/88339887_2840382472710717_3169115093359132672_o.jpg?_nc_cat=107&ccb=1-3&_nc_sid=09cbfe&_nc_ohc=UJjNuNTyouoAX-SAP0G&_nc_ht=scontent.fgua3-2.fna&oh=c2ae32f06c0c7522c5cd90b60b134608&oe=606E33FF"}
	//page5 := TiendaTransicional{"Gerardo Wapo", "Gerardo weco", "gerardoWeco@gmail.com", 69, "https://scontent.fgua3-2.fna.fbcdn.net/v/t1.0-9/88339887_2840382472710717_3169115093359132672_o.jpg?_nc_cat=107&ccb=1-3&_nc_sid=09cbfe&_nc_ohc=UJjNuNTyouoAX-SAP0G&_nc_ht=scontent.fgua3-2.fna&oh=c2ae32f06c0c7522c5cd90b60b134608&oe=606E33FF"}
	//var pages []TiendaTransicional
	//pages = append(pages, page1)
	//pages = append(pages, page2)
	//pages = append(pages, page3)
	//pages = append(pages, page4)
	//pages = append(pages, page5)
	//data, err := json.Marshal(pages)
	data, err := json.Marshal(todasLasTiendas())
	if err != nil {
		response.Write([]byte("ocurrio un error"))
	}
	response.Write(data)
	//response.Write([]byte("Recuerda que lo primero que debes hacer es cargar tu archivo "))

}

func todasLasTiendas() []Tienda {
	matrix := MakeMatrix(JsonData)
	if len(arregloListas) == 0 {
		arregloListas = Linealizar(matrix, JsonData)
	}
	var listaCompleta []Tienda
	for i := 0; i < len(arregloListas); i++ {
		lista := arregloListas[i].ShowJson()
		for j := 0; j < len(lista); j++ {
			listaCompleta = append(listaCompleta, lista[j])
		}
	}
	return listaCompleta
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
	if len(arregloListas) == 0 {
		arregloListas = Linealizar(matrix, JsonData)
	}
	paraEnviar := ShowArray(arregloListas[:])
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
	if len(arregloListas) == 0 {
		arregloListas = Linealizar(matrix, JsonData)
	}
	nombre := arregloListas[numero-1]
	tienda := nombre.ShowJson()
	salida, err3 := json.Marshal(tienda)
	if err3 != nil {
		response.Write([]byte("error en la conversion de json"))
	}
	response.Write(salida)
}

func findTiendaUnica(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	numero, err := strconv.Atoi(vars["numero"])
	if err != nil {
		response.Write([]byte("el id que ingreso es invalido"))
	}
	//este es la linea que tengo que usar para cambiar datos
	tiendaUnica := FindWithId(numero, &arregloListas)
	//tiendaUnica.tienda = "hola"
	salida, _ := json.Marshal(Tienda{tiendaUnica.tienda, tiendaUnica.descripcion, tiendaUnica.contacto, tiendaUnica.logo, tiendaUnica.inventario, tiendaUnica.departamento, tiendaUnica.calificacion, tiendaUnica.id})
	response.Write(salida)
}

func getListaInventario(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	numero, err := strconv.Atoi(vars["numero"])
	if err != nil {
		response.Write([]byte("el id que ingreso es invalido"))
	}
	tiendaUnica := FindWithId(numero, &arregloListas)
	tiendaUnica.inventario.ClearList()
	tiendaUnica.inventario.ListAllProducts(tiendaUnica.inventario.Root)
	salida, _ := json.Marshal(dataStructures.ListElements)
	response.Write(salida)
}

func showInventario(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	numero, err := strconv.Atoi(vars["numero"])
	if err != nil {
		response.Write([]byte("el id que ingreso es invalido"))
	}
	tiendaUnica := FindWithId(numero, &arregloListas)
	tiendaUnica.inventario.MakeGraphviz(tiendaUnica.inventario.Root)
}

func meterElementos(response http.ResponseWriter, request *http.Request) {
	data, err1 := ioutil.ReadAll(request.Body)
	if err1 != nil {
		response.Write([]byte("entrada no valida"))
	}
	var inventarioCompleto []DevolucionInventario
	var entradaMetodo InventariosData
	err := json.Unmarshal(data, &entradaMetodo)
	if err != nil {
		fmt.Println("ocurrio un error")
	}
	inventarioCompleto = FindParaInventario(entradaMetodo, &arregloListas)
	for i := 0; i < len(inventarioCompleto); i++ {
		tienda := inventarioCompleto[i].Nodo
		inventario := inventarioCompleto[i].Lista
		for j := 0; j < len(inventario); j++ {
			var nodoAIngresar *dataStructures.NodoAVL
			if tienda.inventario.Root != nil {
				nodoAIngresar = tienda.inventario.Find(inventario[j].Codigo)
			}
			if nodoAIngresar == nil {
				tienda.inventario.Add(inventario[j])
			} else {
				nodoAIngresar.Valor.Cantidad += inventario[j].Cantidad
			}
		}
	}
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

func addCarrito(response http.ResponseWriter, request *http.Request) {
	//todo tengo que hacer que este metodo jale los datos para guardarlos en el carrito
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
					auxiliar.Logo = tienda.Logo
					return auxiliar
				}
			}
		}
	}

	return Tienda{"Su tienda no se encuentra", "Ingrese una tienda valida", "Algun dato no es correcto", imagenPredeterminada, dataStructures.AVLtree{}, "", 0, 0}
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
