package expressiones

import (
	"Proyecto1/src/environment"
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"
)

type RetornoValorVariable struct {
	Id string
}

// metodo constructor
func NewLLamadoVariable(id string) RetornoValorVariable {
	tempValorVariable := RetornoValorVariable{Id: id}
	return tempValorVariable
}

//conviertiend y implentando el metodo ejecutar de tipo expresion
func (variable RetornoValorVariable) Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interfaces.Simbolo {
	resultado := entorno.(environment.Entornos).GetVariables(variable.Id)
	// fmt.Println("CALLVARIABLE:  retorna->", resultado.Valor.(interfaces.Simbolo))
	return resultado.Valor.(interfaces.Simbolo)
}
