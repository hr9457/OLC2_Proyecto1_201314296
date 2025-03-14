package interfaces

import "Proyecto1/src/traduccion"

// creacion de plantillas para darle estructura a los primitivos
// de los cuales solo necesito saber el tipo que es  y el valor que tiene
// para lo cual es la misma escrtura que mis simbolos o mis variables
// con la diferencia del id la cual sirve solo para para tener la identicacion en variables

// donde se hereda se tenga la posibilidad de teneer un metodo ejecutar
type Expresion interface {
	Ejecutar(entorno interface{}, traductor *traduccion.Traductor) Simbolo
}

type Instruction interface {
	Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interface{}
}
