package instrucciones

import (
	"Proyecto1/src/environment"
	"Proyecto1/src/interfaces"
	"fmt"

	arrayList "github.com/colegno/arraylist"
)

// estructura para guardar los elementos del IF
type If struct {
	Expresion interfaces.Expresion //para guaradar la expresion en el if y pueda ser evaluada
	Contenido *arrayList.List      // contenido dentro del if
}

// consturctor para ingresar datos par utilizad del if
func NewIf(exp interfaces.Expresion, bloque *arrayList.List) If {
	tempIf := If{exp, bloque}
	return tempIf
}

// implementado su metodo Ejecutar
func (firmaif If) Ejecutar(entorno interface{}) interface{} {
	var resultado interfaces.Simbolo

	resultado = firmaif.Expresion.Ejecutar(entorno)

	if resultado.Valor == true {

		// creacion de un etorno para el if
		var entornoIf environment.Entornos

		// transofrmacion de tipos
		entornoIf = environment.NewEntorno(entorno.(environment.Entornos), "if")
		for _, s := range firmaif.Contenido.ToArray() {
			s.(interfaces.Instruction).Ejecutar(entornoIf)
		}

		// fmt.Println(firmaif.Contenido)
	}

	fmt.Println("IF:  retorno->", resultado.Valor)
	return resultado.Valor
}
