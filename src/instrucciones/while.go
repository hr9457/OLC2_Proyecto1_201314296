package instrucciones

import (
	"Proyecto1/src/environment"
	"Proyecto1/src/interfaces"
	"fmt"
	"reflect"

	arrayList "github.com/colegno/arraylist"
)

type While struct {
	Expresion   interfaces.Expresion // expresion de evalcuion del while
	BloqueWhile interface{}          // lista de instruccion para el bloque while
}

// funcio para la construcion de un nuevo while
func NewWhile(exp interfaces.Expresion, bloque interface{}) While {
	tempWhile := While{exp, bloque}
	return tempWhile
}

// metodo ejcutar del while
func (firmaWhile While) Ejecutar(entorno interface{}) interface{} {
	var resultadoExp interfaces.Simbolo
	for {
		resultadoExp = firmaWhile.Expresion.Ejecutar(entorno)
		if resultadoExp.Valor == true {
			var entornoWhile environment.Entornos
			entornoWhile = environment.NewEntorno(entorno.(environment.Entornos), "Entorno While")
			for _, s := range firmaWhile.BloqueWhile.(*arrayList.List).ToArray() {
				// creacion de una variable para compara con los typeof si es un sentencia break
				b := BreakExp{}
				// si es igual a una sentencia break se corrompe
				if reflect.TypeOf(s) == reflect.TypeOf(b) {
					fmt.Println("Se econtro un break")
					return nil
				}
				s.(interfaces.Instruction).Ejecutar(entornoWhile)
			}
		} else {
			break
		}
	}
	return resultadoExp.Valor
}
