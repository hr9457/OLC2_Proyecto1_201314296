package instrucciones

import (
	"Proyecto1/src/environment"
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"
)

type Asignacion struct {
	Id        string               // guarda el identificar de la variable
	Expresion interfaces.Expresion // guarda la expresion
}

// metodo para contruir
func NewAsignacion(id string, exp interfaces.Expresion) Asignacion {
	tempAsigancion := Asignacion{id, exp}
	return tempAsigancion
}

//metodo ejecutar para la asignacion
func (asignacion Asignacion) Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interface{} {
	// variable para guardar el resultado
	// var resultado interfaces.Simbolo
	//polimorfismo
	var resultado = asignacion.Expresion.Ejecutar(entorno, traductor)
	// fmt.Println("ASIGNACION: resultado:", resultado)
	// fmt.Println("ASIGNACION: asignacion:", asignacion)
	// fmt.Println("ASIGNACION:  Tipo de mutabilidad de la variables ", resultado.Mut)
	entorno.(environment.Entornos).AlterVariable(asignacion.Id, resultado)
	// fmt.Println("ASIGNACION:  Tipo de mutabilidad de la variables ", resultado.Mut)
	// fmt.Println("ASIGNACION: ", resultado.Valor)
	return resultado.Valor
}
