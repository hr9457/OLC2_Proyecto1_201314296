package arreglos

import (
	"Proyecto1/src/environment"
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"
)

type DeclaracionArreglo struct {
	Id       string
	Mut      bool
	Tipo     interfaces.TipoExpression
	Tamanio  interfaces.Expresion
	IsArray  bool
	ListaExp interfaces.Expresion
}

// cronstruccion de un arreglo
func NewArreglo(id string, mut bool, tipo interfaces.TipoExpression, tamanio interfaces.Expresion, isArray bool, lista interfaces.Expresion) DeclaracionArreglo {
	temp := DeclaracionArreglo{id, mut, tipo, tamanio, isArray, lista}
	return temp
}

func (arreglo DeclaracionArreglo) Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interface{} {
	// var valor interfaces.Simbolo
	var resultado = arreglo.ListaExp.Ejecutar(entorno, traductor)
	entorno.(environment.Entornos).AddVariable(arreglo.Id, resultado, arreglo.Mut, arreglo.Tipo)
	return nil
}
