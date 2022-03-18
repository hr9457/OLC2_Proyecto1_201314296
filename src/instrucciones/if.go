package instrucciones

import (
	"Proyecto1/src/environment"
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"
	"reflect"

	arrayList "github.com/colegno/arraylist"
)

// estructura para guardar los elementos del IF
type If struct {
	Expresion     interfaces.Expresion //para guaradar la expresion en el if y pueda ser evaluada
	Contenido     *arrayList.List      // contenido dentro del if
	ContenidoElse interface{}          //bloque de contenido para el else
	// Elements      []If                 //
}

// consturctor para ingresar datos par utilizad del if
func NewIf(exp interfaces.Expresion, bloque *arrayList.List, bloqueElse interface{}) If {
	// fmt.Println("IF:   ", bloqueElse)
	tempIf := If{exp, bloque, bloqueElse}
	return tempIf
}

// implementado su metodo Ejecutar
func (firmaif If) Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interface{} {
	var resultado interfaces.Simbolo

	resultado = firmaif.Expresion.Ejecutar(entorno, traductor)

	// entra en la condicional para ejectuar el if
	if resultado.Valor == true {

		// creacion de un etorno para el if
		var entornoIf environment.Entornos

		// transofrmacion de tipos
		entornoIf = environment.NewEntorno(entorno.(environment.Entornos), "Entorno If")
		for _, s := range firmaif.Contenido.ToArray() {

			// creacion de una variable para compara con los typeof si es un sentencia break
			// creand una busqueda y retorno de un break
			b := BreakExp{Valor: nil, Tipo: interfaces.BREAK}
			// si es igual a una sentencia break se corrompe
			if reflect.TypeOf(s) == reflect.TypeOf(b) {
				return b
			}
			ejecucion := s.(interfaces.Instruction).Ejecutar(entornoIf, traductor)
			if reflect.TypeOf(ejecucion) == reflect.TypeOf(b) {
				return b
			}
		}
		// fmt.Println("IF:  retorno->", resultado.Valor)
		// fmt.Println(firmaif.Contenido)
	} else {

		// validacion para el Else
		if firmaif.ContenidoElse != nil {
			// fmt.Println("Este if contiene un else para ejecutar")
			// fmt.Println("ELSE:  ", firmaif.ContenidoElse)
			var entornoElse = environment.NewEntorno(entorno.(environment.Entornos), "Entorno Else")
			for _, s := range firmaif.ContenidoElse.(*arrayList.List).ToArray() {

				// creacion de una variable para compara con los typeof si es un sentencia break
				// creand una busqueda y retorno de un break
				b := BreakExp{Valor: nil, Tipo: interfaces.BREAK}
				// si es igual a una sentencia break se corrompe
				if reflect.TypeOf(s) == reflect.TypeOf(b) {
					return b
				}
				ejecucion := s.(interfaces.Instruction).Ejecutar(entornoElse, traductor)
				if reflect.TypeOf(ejecucion) == reflect.TypeOf(b) {
					return b
				}
			}
		} else {
			// fmt.Println("IF:  retorno->", resultado.Valor)
		}
	}

	// fmt.Println("IF:  retorno->", resultado.Valor)
	return resultado.Valor
}
