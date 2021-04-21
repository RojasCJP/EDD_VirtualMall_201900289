package reader

import (
	"./dataStructures"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

var JsonData = Datos{}
var imagenPredeterminada string = "https://es.jumpseller.com/images/learn/choosing-platform/laptop.jpg"
var arregloListas []Lista
var Carrito []ElementoCarrito
var Years []dataStructures.Year
var Btree *dataStructures.BTree
var Grafo dataStructures.Grafo

//todo tengo que ver que onda con la encriptacion

type htmltemplate struct {
	Name   string
	Carnet int
	Json   string
}

type Usuarios struct {
	Usuarios []dataStructures.Usuario
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
	NombreProducto string
	PrecioProducto float32
	Cantidad       int
	Almacenamiento string
}

func InYear(year int) bool {
	for i := 0; i < len(Years); i++ {
		if year == Years[i].Year {
			return true
		}
	}
	return false
}

func GetYear(year int) *dataStructures.Year {
	for i := 0; i < len(Years); i++ {
		if year == Years[i].Year {
			return &Years[i]
		}
	}
	return nil
}

func GetYearIndex(year int) int {
	for i := 0; i < len(Years); i++ {
		if year == Years[i].Year {
			return i
		}
	}
	return 0
}

func InMeses(mes int, year int) bool {
	for i := 0; i < len(Years[year].Meses); i++ {
		if mes == Years[year].Meses[i].Number {
			return true
		}
	}
	return false
}

func GetMesIndex(mes int, year int) int {
	for i := 0; i < len(Years[year].Meses); i++ {
		if mes == Years[year].Meses[i].Number {
			return i
		}
	}
	return 0
}

func GetMes(mes int, year *dataStructures.Year) *dataStructures.Mes {
	for i := 0; i < len(year.Meses); i++ {
		if mes == year.Meses[i].Number {
			return &year.Meses[i]
		}
	}
	return nil
}

func FindInCarrito(carrito []ElementoCarrito, tienda int, producto int) *ElementoCarrito {
	for i := 0; i < len(carrito); i++ {
		if tienda == carrito[i].IdTienda && producto == carrito[i].CodigoProducto {
			return &carrito[i]
		}
	}
	return nil
}

func indexCarrito(carrito []ElementoCarrito, tienda int, producto int) int {
	for i := 0; i < len(carrito); i++ {
		if tienda == carrito[i].IdTienda && producto == carrito[i].CodigoProducto {
			return i
		}
	}
	return 0
}

func remove(s []ElementoCarrito, i int) []ElementoCarrito {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func SetJsonData(jsonData Datos) {
	JsonData = jsonData

}

func LevantarServer() {
	Btree = dataStructures.BTree_()
	//a;adir datos solo en este btree
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
	router.HandleFunc("/addCarrito/{tienda}/{producto}", addCarrito).Methods("GET")
	router.HandleFunc("/deleteCarrito/{tienda}/{producto}", deleteCarrito).Methods("GET")
	router.HandleFunc("/verCarrito", verCarrito).Methods("GET")
	router.HandleFunc("/comprar", comprar).Methods("GET")
	router.HandleFunc("/calendario", addCalendario).Methods("POST")
	router.HandleFunc("/calendario", verYears).Methods("GET")
	router.HandleFunc("/calendario/{year}", verMeses).Methods("GET")
	router.HandleFunc("/calendario/{year}/{mes}", verCalendario).Methods("GET")
	router.HandleFunc("/calendario/{year}/{mes}/{dia}/{departamento}", verProductos).Methods("GET")
	router.HandleFunc("/calendarioImage", YearsImage).Methods("GET")
	router.HandleFunc("/calendarioImage/{year}", MonthsImage).Methods("GET")
	router.HandleFunc("/calendarioImage/{year}/{mes}", CalendarioImage).Methods("GET")
	router.HandleFunc("/calendarioImage/{year}/{mes}/{dia}/{departamento}", ProductosImage).Methods("GET")
	router.HandleFunc("/usuarios/cargar", CargarUsuario).Methods("POST")
	router.HandleFunc("/usuarios/consultar", ConsultarUsuario).Methods("POST")
	router.HandleFunc("/usuarios/graficar", GrafoUsuario).Methods("GET")
	router.HandleFunc("/usuarios/encriptado", UsuariosEncriptado).Methods("GET")
	router.HandleFunc("/grafo/cargar", CargarGrafo).Methods("POST")
	router.HandleFunc("/grafo/graficar", GrafoGraficar).Methods("GET")
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
	GraphAngular(paraEnviar)
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
				nodoAIngresar = tienda.inventario.Find(inventario[j].Codigo, tienda.inventario.Root)
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
	vars := mux.Vars(request)
	tienda, _ := strconv.Atoi(vars["tienda"])
	producto, _ := strconv.Atoi(vars["producto"])
	elemento := FindInCarrito(Carrito, tienda, producto)
	punteroTienda := FindWithId(tienda, &arregloListas)
	valor := punteroTienda.inventario.Find(producto, punteroTienda.inventario.Root).Valor
	precioProducto := valor.Precio
	nombreProducto := valor.Nombre
	almacenamientoProducto := valor.Almacenamiento
	if elemento == nil {
		Carrito = append(Carrito, ElementoCarrito{tienda, producto, nombreProducto, precioProducto, 1, almacenamientoProducto})
	} else {
		elemento.Cantidad++
	}
}

func deleteCarrito(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	tienda, _ := strconv.Atoi(vars["tienda"])
	producto, _ := strconv.Atoi(vars["producto"])
	index := indexCarrito(Carrito, tienda, producto)
	Carrito = remove(Carrito, index)
}

func verCarrito(response http.ResponseWriter, request *http.Request) {
	salida, _ := json.Marshal(Carrito)
	response.Write(salida)
}

func comprar(response http.ResponseWriter, request *http.Request) {
	for i := 0; i < len(Carrito); i++ {
		tienda := Carrito[i].IdTienda
		producto := Carrito[i].CodigoProducto
		cantidad := Carrito[i].Cantidad
		tiendaCambiar := FindWithId(tienda, &arregloListas)
		inventarioTienda := tiendaCambiar.inventario.Find(producto, tiendaCambiar.inventario.Root)
		inventarioTienda.Valor.Cantidad -= cantidad
	}
	var pedidoEspecifico dataStructures.ValorCola
	for i := 0; i < len(Carrito); i++ {
		tienda := FindWithId(Carrito[i].IdTienda, &arregloListas)
		nombreTienda := tienda.tienda
		year := 2021
		mes := 4
		dia := 20
		departamento := tienda.departamento
		if !InYear(year) {
			var meses []dataStructures.Mes
			Years = append(Years, dataStructures.Year{year, meses})
		}
		if !InMeses(mes, GetYearIndex(year)) {
			yearUse := GetYear(year)
			matriz := dataStructures.Matriz{}
			matriz.Init()
			yearUse.Meses = append(yearUse.Meses, dataStructures.Mes{Number: mes, Matriz: matriz})
		}
		yearComprobation := GetYear(year)
		mesComprobation := GetMes(mes, yearComprobation)
		if mesComprobation.Matriz.Find(dia, departamento) == nil {
			mesComprobation.Matriz.Add(departamento, dia, dataStructures.Cola{Len: 0})
		}
		codigo := dataStructures.Codigo{Codigo: Carrito[i].CodigoProducto}
		var codigos []dataStructures.Codigo
		codigos = append(codigos, codigo)
		pedidoEspecifico = dataStructures.ValorCola{Fecha: "20-04-2021", Tienda: nombreTienda, Departamento: departamento, Calificacion: tienda.calificacion, Productos: codigos}
		mesComprobation.Matriz.Find(dia, departamento).Valor.Add(dataStructures.NodoCola{Valor: dataStructures.ValorCola{pedidoEspecifico.Fecha, pedidoEspecifico.Tienda, pedidoEspecifico.Departamento, pedidoEspecifico.Calificacion, pedidoEspecifico.Productos}})
	}
	Carrito = make([]ElementoCarrito, 0)
}

//todo agregar a calendario

func addCalendario(response http.ResponseWriter, request *http.Request) {
	data, errRead := ioutil.ReadAll(request.Body)
	if errRead != nil {
		response.Write([]byte("error en la carga del json"))
	}
	var calendario = dataStructures.Pedidos{}
	err := json.Unmarshal(data, &calendario)
	if err != nil {
		log.Fatal("error al convertir a estructura " + err.Error())
	}
	for j := 0; j < len(calendario.Pedidos); j++ {
		year := calendario.Pedidos[j].Year()
		mes := calendario.Pedidos[j].Mes()
		dia := calendario.Pedidos[j].Dia()
		departamento := calendario.Pedidos[j].Departamento
		//var tienda *Nodo
		//for k := 0; k < len(arregloListas); k++ {
		//	tienda = arregloListas[k].FindParaInventario(paraInventario{calendario.Pedidos[j].Tienda, departamento, calendario.Pedidos[j].Calificacion})
		//}
		if !InYear(year) {
			var meses []dataStructures.Mes
			Years = append(Years, dataStructures.Year{year, meses})
		}
		if !InMeses(mes, GetYearIndex(year)) {
			yearUse := GetYear(year)
			matriz := dataStructures.Matriz{}
			matriz.Init()
			yearUse.Meses = append(yearUse.Meses, dataStructures.Mes{Number: mes, Matriz: matriz})
		}
		yearComprobation := GetYear(year)
		mesComprobation := GetMes(mes, yearComprobation)
		if mesComprobation.Matriz.Find(dia, departamento) == nil {
			mesComprobation.Matriz.Add(departamento, dia, dataStructures.Cola{Len: 0})
		}
		pedidoEspecifico := calendario.Pedidos[j]

		mesComprobation.Matriz.Find(dia, departamento).Valor.Add(dataStructures.NodoCola{Valor: dataStructures.ValorCola{pedidoEspecifico.Fecha, pedidoEspecifico.Tienda, pedidoEspecifico.Departamento, pedidoEspecifico.Calificacion, pedidoEspecifico.Productos}})
	}
}

//todo ver como se agregan los pedidos

func verYears(response http.ResponseWriter, request *http.Request) {
	var allYears []int
	for i := 0; i < len(Years); i++ {
		allYears = append(allYears, Years[i].Year)
	}
	data, _ := json.Marshal(allYears)
	response.Write(data)
}

func verMeses(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	year, _ := strconv.Atoi(vars["year"])
	yearIndex := GetYearIndex(year)
	var allMonths []int
	for i := 0; i < len(Years[yearIndex].Meses); i++ {
		allMonths = append(allMonths, Years[yearIndex].Meses[i].Number)
	}
	data, _ := json.Marshal(allMonths)
	response.Write(data)
}

func verCalendario(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	year, _ := strconv.Atoi(vars["year"])
	month, _ := strconv.Atoi(vars["mes"])
	yearUse := GetYear(year)
	monthUse := GetMes(month, yearUse)
	data, _ := json.Marshal(monthUse.Matriz.ReturnListNodes())
	response.Write(data)
}

func verProductos(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	year, _ := strconv.Atoi(vars["year"])
	month, _ := strconv.Atoi(vars["mes"])
	dia, _ := strconv.Atoi(vars["dia"])
	departamento := vars["departamento"]
	yearUse := GetYear(year)
	monthUse := GetMes(month, yearUse)
	cola := monthUse.Matriz.Find(dia, departamento).Valor
	respuesta := cola.AllProducts()
	data, _ := json.Marshal(respuesta)
	response.Write(data)
}

func YearsImage(response http.ResponseWriter, request *http.Request) {
	cuerpo := "digraph matriz{\n    rankdir = \"LR\"\n"
	for i := 0; i < len(Years); i++ {
		if len(Years) == 1 {
			cuerpo += strconv.Itoa(Years[i].Year) + "\n"
		}
		if i < len(Years)-1 {
			cuerpo += strconv.Itoa(Years[i].Year) + "->" + strconv.Itoa(Years[i+1].Year) + "\n"
		}
		if i > 0 {
			cuerpo += strconv.Itoa(Years[i].Year) + "->" + strconv.Itoa(Years[i-1].Year) + "\n"
		}
	}
	cuerpo += "}"
	err := ioutil.WriteFile("../Cliente/src/assets/graphviz/matriz.dot", []byte(cuerpo), 0644)
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

func MonthsImage(response http.ResponseWriter, request *http.Request) {
	cuerpo := "digraph matriz{\n    rankdir = \"LR\"\n"
	vars := mux.Vars(request)
	year, _ := strconv.Atoi(vars["year"])
	yearIndex := GetYearIndex(year)
	for i := 0; i < len(Years[yearIndex].Meses); i++ {
		if len(Years[yearIndex].Meses) == 1 {
			cuerpo += strconv.Itoa(Years[yearIndex].Meses[i].Number) + "\n"
		}
		if i < len(Years[yearIndex].Meses)-1 {
			cuerpo += strconv.Itoa(Years[yearIndex].Meses[i].Number) + "->" + strconv.Itoa(Years[yearIndex].Meses[i+1].Number) + "\n"
		}
		if i > 0 {
			cuerpo += strconv.Itoa(Years[yearIndex].Meses[i].Number) + "->" + strconv.Itoa(Years[yearIndex].Meses[i-1].Number) + "\n"
		}
	}
	cuerpo += "}"
	err := ioutil.WriteFile("../Cliente/src/assets/graphviz/matriz.dot", []byte(cuerpo), 0644)
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

func CalendarioImage(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	year, _ := strconv.Atoi(vars["year"])
	month, _ := strconv.Atoi(vars["mes"])
	yearUse := GetYear(year)
	monthUse := GetMes(month, yearUse)
	monthUse.Matriz.MakeGraph()
}

func ProductosImage(response http.ResponseWriter, request *http.Request) {
	cuerpo := "digraph matriz{\n    rankdir = \"LR\"\n"
	vars := mux.Vars(request)
	year, _ := strconv.Atoi(vars["year"])
	month, _ := strconv.Atoi(vars["mes"])
	dia, _ := strconv.Atoi(vars["dia"])
	departamento := vars["departamento"]
	yearUse := GetYear(year)
	monthUse := GetMes(month, yearUse)
	cola := monthUse.Matriz.Find(dia, departamento).Valor
	respuesta := cola.AllProducts()
	for i := 0; i < len(respuesta); i++ {
		if len(respuesta) == 1 {
			cuerpo += strconv.Itoa(respuesta[i]) + "\n"
		}
		if i < len(respuesta)-1 {
			cuerpo += strconv.Itoa(respuesta[i]) + "->" + strconv.Itoa(respuesta[i+1]) + "\n"
		}
		if i > 0 {
			cuerpo += strconv.Itoa(respuesta[i]) + "->" + strconv.Itoa(respuesta[i-1]) + "\n"
		}
	}
	cuerpo += "}"
	err := ioutil.WriteFile("../Cliente/src/assets/graphviz/matriz.dot", []byte(cuerpo), 0644)
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

func CargarUsuario(response http.ResponseWriter, request *http.Request) {
	data, _ := ioutil.ReadAll(request.Body)
	usuarios := Usuarios{}
	err := json.Unmarshal(data, &usuarios)
	if err != nil {
		response.Write([]byte("ocurrio un error en la carga"))
	}
	for i := 0; i < len(usuarios.Usuarios); i++ {
		if Btree.Find(usuarios.Usuarios[i].Dpi, Btree.Root).Dpi == 0 {
			Btree.Insert(&dataStructures.Usuario{Dpi: usuarios.Usuarios[i].Dpi, Nombre: usuarios.Usuarios[i].Nombre,
				Correo: usuarios.Usuarios[i].Correo, Password: usuarios.Usuarios[i].Password, Cuenta: usuarios.Usuarios[i].Cuenta, ContraEncriptada: fmt.Sprintf("x", sha256.Sum256([]byte(usuarios.Usuarios[i].Password)))})
		}
	}
	response.Write(data)
}

type ConsultaSesion struct {
	Dpi      int
	Password string
}
type RespuestaSesion struct {
	Consulta bool
	Cuenta   string
}

func ConsultarUsuario(response http.ResponseWriter, request *http.Request) {
	data, _ := ioutil.ReadAll(request.Body)
	consultaSesion := ConsultaSesion{}
	json.Unmarshal(data, &consultaSesion)
	respuestaSesion := RespuestaSesion{Consulta: false, Cuenta: "Usuario"}
	usuarioBusqueda := Btree.Find(consultaSesion.Dpi, Btree.Root)
	if usuarioBusqueda.Password == consultaSesion.Password {
		respuestaSesion.Consulta = true
		respuestaSesion.Cuenta = usuarioBusqueda.Cuenta
	}
	res, _ := json.Marshal(respuestaSesion)
	response.Write(res)
}

func CargarGrafo(response http.ResponseWriter, request *http.Request) {
	data, _ := ioutil.ReadAll(request.Body)
	grafoEntrada := dataStructures.Grafo{}
	err := json.Unmarshal(data, &grafoEntrada)
	if err != nil {
		response.Write([]byte("ocurrio un error en la carga"))
	}
	Grafo.PosicionInicialRobot = grafoEntrada.PosicionInicialRobot
	Grafo.Entrega = grafoEntrada.Entrega
	for i := 0; i < len(grafoEntrada.Nodos); i++ {
		Grafo.AddNode(grafoEntrada.Nodos[i].Nombre, grafoEntrada.Nodos[i].Enlaces)
	}
	response.Write(data)
}

func GrafoUsuario(response http.ResponseWriter, request *http.Request) {
	Btree.MakeGraphviz()
}

func UsuariosEncriptado(response http.ResponseWriter, request *http.Request) {
	Btree.MakeGraphvizEncriptado()
}

type GrafoCamino struct {
	Inicio string
	Final  string
}

func GrafoGraficar(response http.ResponseWriter, request *http.Request) {
	var camino []string
	camino = append(camino, Grafo.PosicionInicialRobot)
	for i := 0; i < len(Carrito); i++ {
		camino = append(camino, Carrito[i].Almacenamiento)
	}
	camino = append(camino, Grafo.Entrega)
	Grafo.MakeFileGrafo(camino)
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
