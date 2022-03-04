package instrucciones

import (
	"Proyecto1/src/environment"
	"Proyecto1/src/interfaces"
	"fmt"
)

type Declaracion struct {
	Id        string                    //identificador
	Tipo      interfaces.TipoExpression // el tipo de declaracion
	Expresion interfaces.Expresion      // la expresion 	que se va guardar
}

//creo la declarcion con los parmamentros que necesito
func NewDeclaracion(id string, tipo interfaces.TipoExpression, expresion interfaces.Expresion) Declaracion {
	declaracionTemporal := Declaracion{id, tipo, expresion}
	return declaracionTemporal
}

func (declaracion Declaracion) Ejecutar(entorno interface{}) interface{} {
	// var resultado interfaces.Simbolo
	//va ser igual a un objeto tipo simbolo
	var resultado = declaracion.Expresion.Ejecutar(entorno)
	if resultado.Tipo == declaracion.Tipo {
		entorno.(environment.Entornos).AddVariable(declaracion.Id, resultado, false, declaracion.Tipo)
	}
	fmt.Println("El valor de la variable a guardar es: ", resultado.Valor)
	return resultado.Valor
}
