package instrucciones

import (
	"Proyecto1/src/interfaces"
	"fmt"

	arrayList "github.com/colegno/arraylist"
)

// estructura para guardar los elementos del IF
type If2 struct {
	Expresion1    interfaces.Expresion //
	Expresion     interfaces.Expresion //para guaradar la expresion en el if y pueda ser evaluada
	Contenido     *arrayList.List      // contenido dentro del if
	ContenidoElse interface{}          //bloque de contenido para el else
}

// consturctor para ingresar datos par utilizad del if
func NewIf2(exp1 interfaces.Expresion, exp interfaces.Expresion, bloque *arrayList.List, bloqueElse interface{}) If2 {
	// fmt.Println("IF:   ", bloqueElse)
	tempIf2 := If2{exp1, exp, bloque, bloqueElse}
	return tempIf2
}

// implementado su metodo Ejecutar
func (firmaif If2) Ejecutar(entorno interface{}) interface{} {
	var resultado interfaces.Simbolo
	resultado = firmaif.Expresion1.Ejecutar(entorno)

	var resultado2 interfaces.Simbolo
	resultado2 = firmaif.Expresion.Ejecutar(entorno)
	fmt.Println(resultado)
	if resultado2.Valor == true {
		fmt.Println("true")
		fmt.Println("IF:  retorno->", resultado.Valor)
		return resultado.Valor
	} else {
		fmt.Println("false")
		fmt.Println("IF:  retorno->", resultado.Valor)
		return resultado2.Valor

	}

}
