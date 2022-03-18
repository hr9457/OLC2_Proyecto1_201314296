package instrucciones

import (
	"Proyecto1/src/environment"
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"
	"reflect"

	arrayList "github.com/colegno/arraylist"
)

// estructura para guardar los elementos del IF
type If2 struct {
	Expresion     interfaces.Expresion //para guaradar la expresion en el if y pueda ser evaluada
	Contenido     *arrayList.List      // contenido dentro del if
	ContenidoElse interface{}          //bloque de contenido para el else
	Listado       *arrayList.List      // listado
}

// consturctor para ingresar datos par utilizad del if
func NewIf2(exp interfaces.Expresion, bloque *arrayList.List, bloqueElse interface{}, listado *arrayList.List) If2 {
	// fmt.Println("IF:   ", bloqueElse)
	tempIf := If2{exp, bloque, bloqueElse, listado}
	return tempIf
}

// implementado su metodo Ejecutar
func (firmaif If2) Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interface{} {
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

		// revision de los bloque del else if
		if firmaif.Listado != nil {
			// primer recorrrido la lista principal
			for _, elif := range firmaif.Listado.ToArray() {
				// fmt.Println(reflect.TypeOf(elif))
				// fmt.Println("Else If: ", elif.(If).Expresion.Ejecutar(entorno).Valor)
				if elif.(If).Expresion.Ejecutar(entorno, traductor).Valor == true {
					// fmt.Println("ELSE IF:  condicion verdadera")
					var entornoElif environment.Entornos
					entornoElif = environment.NewEntorno(entorno.(environment.Entornos), "Entorno Else if")
					for _, s := range elif.(If).Contenido.ToArray() {

						// creacion de una variable para compara con los typeof si es un sentencia break
						// creand una busqueda y retorno de un break
						b := BreakExp{Valor: nil, Tipo: interfaces.BREAK}
						// si es igual a una sentencia break se corrompe
						if reflect.TypeOf(s) == reflect.TypeOf(b) {
							return b
						}
						ejecucion := s.(interfaces.Instruction).Ejecutar(entornoElif, traductor)
						// return resultado.Valor
						if reflect.TypeOf(ejecucion) == reflect.TypeOf(b) {
							return b
						}
					}
				}
			}

		}

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
		}
	}

	// fmt.Println("IF:  retorno->", resultado.Valor)
	return resultado.Valor
}
