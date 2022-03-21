package arreglos

import (
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"
	"fmt"

	arrayList "github.com/colegno/arraylist"
)

type ArrayAccess struct {
	Arr   interfaces.Expresion
	Index interfaces.Expresion
}

func NewArrayAccess(lista interfaces.Expresion, index interfaces.Expresion) ArrayAccess {
	temp := ArrayAccess{Arr: lista, Index: index}
	return temp
}

func (arr ArrayAccess) Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interfaces.Simbolo {

	var tempArray interfaces.Simbolo
	tempArray = arr.Arr.Ejecutar(entorno, traductor)

	fmt.Println("ACCESO; ", tempArray)
	if tempArray.Tipo == interfaces.ARREGLO {

		var tempIndex interfaces.Simbolo
		tempIndex = arr.Index.Ejecutar(entorno, traductor)

		var tempValor interface{}
		tempValor = tempArray.Valor
		
		return tempValor.(*arrayList.List).GetValue(tempIndex.Valor.(int)).(interfaces.Simbolo)
	}
	return interfaces.Simbolo{
		Id:    "",
		Tipo:  interfaces.NULL,
		Valor: "ERROR",
	}
}
