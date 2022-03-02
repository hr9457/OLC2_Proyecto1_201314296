package instrucciones

import (
	"Proyecto1/src/interfaces"
	"fmt"
)

type Imprimir struct {
	Contenido interfaces.Expresion
}

// construccion del metodo para construir la funcion imprimir y almacenar el contenido de lo que se quiere imprimir y la funcion
func NewImprimir(contendio interfaces.Expresion) Imprimir {
	tempImprimir := Imprimir{contendio}
	return tempImprimir
}

// ejecutando lo metodos que ncesita ejectura dentro del print
func (imprimir Imprimir) Ejecutar() interface{} {
	var resultado interfaces.Simbolo
	resultado = imprimir.Contenido.Ejecutar()
	fmt.Println("Mostrando lo que tiene que estar dentro del println")
	fmt.Println(resultado.Valor)
	return resultado
}
