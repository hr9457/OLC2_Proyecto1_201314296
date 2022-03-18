package expressiones

import (
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"
)

// necsito una estructura para poder resivir los datos
// lo cual me intersa saber el cual es el valor y el tipo
type Primitivo struct {
	Valor         interface{}               // que puede recibir cualquier valor para todos los existentes en el lenguaje
	TipoPrimitivo interfaces.TipoExpression // el cual me guarda el tipo de primitivo que puede ser
}

// funcion para crear un nuevo primitivo
// el cual me permitir crear, guadar y estructurar el primitivo de
//  a partir del structur primitivo
func NewPrimito(valor interface{}, tipo interfaces.TipoExpression) Primitivo {
	primitivoTemporal := Primitivo{valor, tipo}
	return primitivoTemporal
}

// funcion para retornar el primitivo con la estructura de un simbolo/variable
func (primate Primitivo) Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interfaces.Simbolo {
	// retrono el primitivo con una estructrua definida
	// fmt.Println("PRIMATE:  tipo->", primate.TipoPrimitivo)
	return interfaces.Simbolo{
		Valor: primate.Valor,
		Tipo:  primate.TipoPrimitivo,
	}
}
