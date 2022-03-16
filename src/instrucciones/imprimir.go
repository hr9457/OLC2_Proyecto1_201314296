package instrucciones

import (
	"Proyecto1/src/interfaces"
	"fmt"
	"strings"

	arrayList "github.com/colegno/arraylist"
)

type Imprimir struct {
	Contenido interfaces.Expresion
	Listado   interface{}
}

// construccion del metodo para construir la funcion imprimir y almacenar el contenido de lo que se quiere imprimir y la funcion
func NewImprimir(contendio interfaces.Expresion, list interface{}) Imprimir {
	tempImprimir := Imprimir{contendio, list}
	return tempImprimir
}

// ejecutando lo metodos que ncesita ejectura dentro del print
func (imprimir Imprimir) Ejecutar(entorno interface{}) interface{} {
	var resultado interfaces.Simbolo
	// fmt.Println("Contendio a imprimir en el println: ", imprimir.Contenido)
	resultado = imprimir.Contenido.Ejecutar(entorno)
	if resultado.Tipo == interfaces.STRING {
		// fmt.Println("IMPRIMIR:  lo que se va imprimir es un string")
		if imprimir.Listado != nil {

			// fmt.Println("IMPRIMIR:  el print trae un lisado de variable")
			countResultado := strings.Count(resultado.Valor.(string), "{}")
			// fmt.Println("IMPRIMIR: se encontraron {}", countResultado)
			if countResultado == (imprimir.Listado.(*arrayList.List).Len()) {

				// fmt.Println("IMPRIMIR: 	listado si conicide con la cantidad de {}")
				// buscando dentro del print las posibilidades de signo
				var textNew string
				textNew = resultado.Valor.(string)

				if imprimir.Listado != nil {
					for _, s := range imprimir.Listado.(*arrayList.List).ToArray() {
						// fmt.Println(s.(interfaces.Expresion).Ejecutar(entorno).Valor)
						a := s.(interfaces.Expresion).Ejecutar(entorno).Valor
						// fmt.Println(reflect.TypeOf(s.(interfaces.Expresion).Ejecutar(entorno).Valor))
						// fmt.Println(reflect.TypeOf(strconv.Itoa(a.(int))))
						// remplazo := s.(interfaces.Expresion).Ejecutar(entorno).Valor
						dato := fmt.Sprintf("%v", a)
						ns := strings.Replace(textNew, "{}", dato, 1)
						textNew = ns
					}
				}
				// fmt.Println("IMPRIMIR:  ", resultado)
				fmt.Println("IMPRIMIR:  Mostrando lo que tiene que estar dentro del println: ", textNew)

			} else {
				fmt.Println("IMPRIMIR: 	ERROR-> listado no conicide con la cantidad de {}")
			}

		} else {

			// fmt.Println("IMPRIMIR: 	ERROR-> no hay un listado de variables")
			fmt.Println("IMPRIMIR:  Mostrando lo que tiene que estar dentro del println: ", resultado.Valor)

		}
	} else {
		fmt.Println("IMPRIMIR:  ERROR-> Error Lexico")
	}

	// fmt.Println("IMPRIMIR:  Mostrando lo que tiene que estar dentro del println: ", resultado.Valor)
	return resultado.Valor
}
