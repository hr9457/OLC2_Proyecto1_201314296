package instrucciones

import (
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"
)

// struct break
type BreakExp struct {
	Valor interface{}               //guaarda el valor del break
	Tipo  interfaces.TipoExpression // guardar el tipo de expresion
}

// metodo para contruir un Break
func NewBreak(tipo interfaces.TipoExpression) BreakExp {
	tempBreak := BreakExp{nil, tipo}
	return tempBreak
}

func (firmaBreak BreakExp) Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interface{} {
	return nil
}
