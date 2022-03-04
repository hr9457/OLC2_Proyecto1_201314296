package expressiones

import (
	"Proyecto1/src/environment"
	"Proyecto1/src/interfaces"
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
func (variable RetornoValorVariable) Ejecutar(entorno interface{}) interfaces.Simbolo {
	resultado := entorno.(environment.Entornos).GetVariables(variable.Id)
	return resultado.Valor.(interfaces.Simbolo)
}
