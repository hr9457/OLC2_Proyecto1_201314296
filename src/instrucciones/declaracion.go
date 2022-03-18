package instrucciones

import (
	"Proyecto1/src/environment"
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"
)

type Declaracion struct {
	Id        string                    //identificador
	Mut       bool                      //mutabilidad de la variable
	Tipo      interfaces.TipoExpression // el tipo de declaracion
	Expresion interfaces.Expresion      // la expresion 	que se va guardar
}

//creo la declarcion con los parmamentros que necesito
func NewDeclaracion(id string, mut bool, tipo interfaces.TipoExpression, expresion interfaces.Expresion) Declaracion {
	// fmt.Println("DECLARACION:  exprsion que se guardara :", expresion.(expressiones.Primitivo).TipoPrimitivo)
	// fmt.Println("DECLARACION:  tipo->", reflect.TypeOf(expresion))
	declaracionTemporal := Declaracion{id, mut, tipo, expresion}
	// fmt.Println("DECLARACION:  ", declaracionTemporal)
	return declaracionTemporal
}

func (declaracion Declaracion) Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interface{} {
	// var resultado interfaces.Simbolo
	//va ser igual a un objeto tipo simbolo
	var resultado = declaracion.Expresion.Ejecutar(entorno, traductor)

	// verificacion si la declaracion tiene un tip declarado para guardar
	if declaracion.Tipo == interfaces.NULL {
		// fmt.Println("la variable no  tiene tipo la cual sera dada por la opoeracion")
		// fmt.Println("tipo", resultado.Tipo)
		entorno.(environment.Entornos).AddVariable(declaracion.Id, resultado, declaracion.Mut, resultado.Tipo)

	} else {

		//verificacon de tipo de primitivo y el tipo de la variable donde se quiere guardar
		if resultado.Tipo == declaracion.Tipo {
			// fmt.Println("DECLARACION:  ", declaracion)
			// fmt.Println("DECLARACION: resultado: ", resultado)
			entorno.(environment.Entornos).AddVariable(declaracion.Id, resultado, declaracion.Mut, declaracion.Tipo)
		}
		// fmt.Println("DECLARACION:  El valor de la variable a guardar es: ", resultado.Valor)
		// fmt.Println("DECLARACION:  Variable tiene el valor de: ", resultado.Valor)
	}
	return resultado.Valor
}
