package instrucciones

import (
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"
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

				for _, s := range imprimir.Listado.(*arrayList.List).ToArray() {

					// fmt.Println(s.(interfaces.Expresion).Ejecutar(entorno).Valor)
					// recupero el dato que me retorno la ejecuccion
					// fmt.Println("IMPRIMIR:  ENTRO AL FOR")
					// fmt.Println("IMPRIMIR:  ", s)
					// fmt.Println(reflect.TypeOf(s))
					a := s.(interfaces.Expresion).Ejecutar(entorno, traductor)

					// fmt.Println("-->", a)

					if a.Tipo == interfaces.ARREGLO {

						// fmt.Println("Lo que se va imprimir es un arreglo")
						var result interfaces.Simbolo
						result = s.(interfaces.Expresion).Ejecutar(entorno, traductor)
						// fmt.Println(reflect.TypeOf(result))
						var tempConcat string
						tempConcat += "[ "
						for _, s := range result.Valor.(*arrayList.List).ToArray() {
							if s.(interfaces.Simbolo).Tipo == interfaces.ARREGLO {
								fmt.Println("se eccontro un arreglo dentro del arreglo")
								retorno := printArreglos(s, entorno, traductor)
								tempConcat += retorno
								// fmt.Println(reflect.TypeOf(s.(interfaces.Simbolo).Valor))
							} else {
								dato := fmt.Sprintf("%v", s.(interfaces.Simbolo).Valor)
								// fmt.Println(reflect.TypeOf(s))
								// fmt.Println(s.(interfaces.Simbolo).Valor)
								tempConcat += dato + " "
							}
						}
						tempConcat += " ]"
						ns := strings.Replace(textNew, "{:?}", tempConcat, 1)
						textNew = ns

					} else {

						fmt.Println("IMPRIMIR: ", a)
						// si lo que viene son variable
						dato := fmt.Sprintf("%v", a.Valor)
						ns := strings.Replace(textNew, "{}", dato, 1)
						textNew = ns
					}
				}
				// }
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

func printArreglos(s interface{}, entorno interface{}, traductor *traduccion.Traductor) string {
	var tempConcat string
	tempConcat += "[ "
	// fmt.Println(s.(interfaces.Simbolo).Valor)
	for _, s := range s.(interfaces.Simbolo).Valor.(*arrayList.List).ToArray() {
		// dato := fmt.Sprintf("%v", s.(interfaces.Simbolo).Valor)
		if s.(interfaces.Simbolo).Tipo == interfaces.ARREGLO {
			retorno := printArreglos(s, entorno, traductor)
			tempConcat += retorno
		} else {
			dato := fmt.Sprintf("%v", s.(interfaces.Simbolo).Valor)
			// fmt.Println(reflect.TypeOf(s))
			// fmt.Println(s.(interfaces.Simbolo).Valor)
			tempConcat += dato + " "
		}
	}
	tempConcat += " ]"
	return tempConcat
}
