package instrucciones

import (
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"
	"fmt"
	"reflect"
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
func (imprimir Imprimir) Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interface{} {
	var resultado interfaces.Simbolo
	// fmt.Println("Contendio a imprimir en el println: ", imprimir.Contenido)
	resultado = imprimir.Contenido.Ejecutar(entorno, traductor)
	if resultado.Tipo == interfaces.STRING {
		// fmt.Println("IMPRIMIR:  lo que se va imprimir es un string")
		if imprimir.Listado != nil {

			// fmt.Println("IMPRIMIR:  el print trae un lisado de variable")
			countVectores := strings.Count(resultado.Valor.(string), "{:?}")
			countVariables := strings.Count(resultado.Valor.(string), "{}")
			// junto los resultados tanto de variables como de vectores
			countResultado := countVectores + countVariables

			// fmt.Println("IMPRIMIR: se encontraron {}", countResultado)
			if countResultado == (imprimir.Listado.(*arrayList.List).Len()) {

				// buscando dentro del print las posibilidades de signo
				// casteamos el valor del print en un string para poder imprimirlo
				var textNew string
				textNew = resultado.Valor.(string)

				// reviso si tengo algo en la parte de los listado para imprimir
				if imprimir.Listado != nil {

					// verificacion de tipo ara un listado de
					lista := arrayList.New()

					

					for _, s := range imprimir.Listado.(*arrayList.List).ToArray() {
						// fmt.Println(s.(interfaces.Expresion).Ejecutar(entorno).Valor)
						// recupero el dato que me retorno la ejecuccion
						a := s.(interfaces.Expresion).Ejecutar(entorno, traductor).Valor
						// fmt.Println(reflect.TypeOf(s.(interfaces.Expresion).Ejecutar(entorno).Valor))
						// fmt.Println(reflect.TypeOf(strconv.Itoa(a.(int))))
						// remplazo := s.(interfaces.Expresion).Ejecutar(entorno).Valor

						// fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(lista))
						fmt.Println(a)

						// si viene un arreglo
						if reflect.TypeOf(a) == reflect.TypeOf(lista) {
							// fmt.Println(a)
							var concatenacion string
							concatenacion += "["
							for _, s := range a.(*arrayList.List).ToArray() {
								// fmt.Println(s.(interfaces.Simbolo).Valor)
								concatenacion += fmt.Sprintf("%v", s.(interfaces.Simbolo).Valor) + " ,"
								// fmt.Println(reflect.TypeOf(s.(interfaces.Simbolo).Valor))
							}
							concatenacion += "]"
							ns := strings.Replace(textNew, "{:?}", concatenacion, 1)
							textNew = ns

						} else {

							// si lo que viene son variable
							dato := fmt.Sprintf("%v", a)
							ns := strings.Replace(textNew, "{}", dato, 1)
							textNew = ns
						}
					}
				}
				// fmt.Println("IMPRIMIR:  Mostrando lo que tiene que estar dentro del println: ->", textNew)
				traductor.Contenido += textNew + "\n"
				// return textNew

			} else {
				// fmt.Println("IMPRIMIR: 	ERROR-> listado no conicide con la cantidad de {}")
				// return "IMPRIMIR: 	ERROR-> listado no conicide con la cantidad de {}"
				traductor.Contenido += "IMPRIMIR: 	ERROR-> listado no conicide con la cantidad de {}" + "\n"
			}

		} else {

			// fmt.Println("IMPRIMIR: 	ERROR-> no hay un listado de variables")
			// fmt.Println("IMPRIMIR:  Mostrando lo que tiene que estar dentro del println: ", resultado.Valor)
			dato := fmt.Sprintf("%v", resultado.Valor)
			traductor.Contenido += dato + "\n"
		}
	} else {
		// fmt.Println("IMPRIMIR:  ERROR-> Error Lexico")
		traductor.Contenido += "IMPRIMIR:  ERROR-> Error Lexico" + "\n"
		// return "IMPRIMIR:  ERROR-> Error Lexico"
	}

	// fmt.Println("IMPRIMIR:  Mostrando lo que tiene que estar dentro del println: ", resultado.Valor)
	return resultado.Valor
}
