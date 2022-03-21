package arreglos

import (
	"Proyecto1/src/environment"
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"
	"fmt"
)

type DeclaracionArreglo struct {
	Id        string
	Mut       bool
	Tipo      interfaces.TipoExpression
	Tamanio   string
	IsArray   bool
	Expresion interfaces.Expresion
}

// cronstruccion de un arreglo
func NewArreglo(id string, mut bool, tipo interfaces.TipoExpression, tamanio string, isArray bool, lista interfaces.Expresion) DeclaracionArreglo {
	temp := DeclaracionArreglo{id, mut, tipo, tamanio, isArray, lista}
	return temp
}

func (arreglo DeclaracionArreglo) Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interface{} {
	// var resultado interfaces.Simbolo
	//va ser igual a un objeto tipo simbolo
	var resultado interfaces.Simbolo
	resultado = arreglo.Expresion.Ejecutar(entorno, traductor)

	// verificacion si la declaracion tiene un tip declarado para guardar
	if resultado.Tipo == arreglo.Tipo {
		// fmt.Println("la variable no  tiene tipo la cual sera dada por la opoeracion")
		entorno.(environment.Entornos).AddVariable(arreglo.Id, resultado, arreglo.Mut, resultado.Tipo)

	} else if arreglo.IsArray {

		entorno.(environment.Entornos).AddVariable(arreglo.Id, resultado, arreglo.Mut, interfaces.ARREGLO)
	} else {

		fmt.Println("LOS TIPOS NO COINCIDEN")
	}
	return resultado.Valor
}
