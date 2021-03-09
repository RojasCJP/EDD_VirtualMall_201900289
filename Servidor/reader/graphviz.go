package reader

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

var graphviz string
var rank string

func GraphvizMethod(valores string) {
	graphviz = "digraph getAll {\ngraph[splines=\"ortho\"];\n"
	numeroListas := hacerGraph(valores)
	rango(numeroListas)
	graphviz += rank
	graphviz += "}"
	err := ioutil.WriteFile("templates/getAll.dot", []byte(graphviz), 0644)
	if err != nil {
		log.Fatal(err)
	}
	s := "dot.exe -Tpng templates/getAll.dot -o templates/imagen.png"
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
func hacerGraph(cadena string) int {
	grafo := []byte(cadena)
	var numeroLista int
	var nombreTienda string
	var numeroTienda int
	var numeroTiendaLista int
	for _, letra := range grafo {
		if string(letra) == "[" {
			numeroTiendaLista = 0
			numeroLista++
			numero := strconv.Itoa(numeroLista)
			graphviz += "Lista" + numero + "[label=\"" + numero + "\",shape=\"box\"];\n"
		} else if string(letra) == "," {
			numero := strconv.Itoa(numeroLista)
			numeroAnterior := strconv.Itoa(numeroLista - 1)
			if numeroLista > 1 {
				graphviz += "Lista" + numeroAnterior + "->Lista" + numero + ";\n"
			}
		} else if (string(letra) == "]") || (string(letra) == ";") {
			numero := strconv.Itoa(numeroLista)
			if nombreTienda != "" {
				numeroTienda++
				numeroTiendaStr := strconv.Itoa(numeroTienda)
				numeroTiendaAnterior := strconv.Itoa(numeroTienda - 1)
				numeroTiendaLista++
				if numeroTiendaLista == 1 {
					graphviz += "elemento" + numeroTiendaStr + "[label=\"" + nombreTienda + "\"];\nLista" + numero + "->elemento" + numeroTiendaStr + ";\n elemento" + numeroTiendaStr + "->Lista" + numero + ";\n"
				} else {
					graphviz += "elemento" + numeroTiendaStr + "[label=\"" + nombreTienda + "\",];\nelemento" + numeroTiendaAnterior + "->elemento" + numeroTiendaStr + ";\n elemento" + numeroTiendaStr + "->elemento" + numeroTiendaAnterior + ";"
				}
			}
			if string(letra) == "]" {
				numeroTiendaLista = 0
			}
			nombreTienda = ""
		} else {
			nombreTienda += string(letra)
		}
	}
	return numeroLista
}

func rango(numeroLista int) {
	rank = "{rank=\"same\";"
	for i := 0; i < numeroLista; i++ {
		numero := strconv.Itoa(i + 1)
		rank += "Lista" + numero + ";"
	}
	rank += "}"
}
