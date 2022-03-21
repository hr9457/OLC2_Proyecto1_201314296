package arreglos

import (
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"

	arrayList "github.com/colegno/arraylist"
)

type Array struct {
	ListaExp *arrayList.List
}

func NewArray(lista *arrayList.List) Array {
	temp := Array{lista}
	return temp
}

func (arr Array) Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interfaces.Simbolo {

	var temp *arrayList.List
	temp = arrayList.New()
	for _, s := range arr.ListaExp.ToArray() {
		temp.Add(s.(interfaces.Expresion).Ejecutar(entorno, traductor))
	}
	// interfaces.Simbolo{Id: "", Valor: resultado, Mut: "", Tipo: resultadoTipo}
	return interfaces.Simbolo{
		Id:    "",
		Tipo:  interfaces.ARREGLO,
		Valor: temp,
	}
}
