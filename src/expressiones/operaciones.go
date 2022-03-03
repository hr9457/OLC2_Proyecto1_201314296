package expressiones

import (
	"Proyecto1/src/interfaces"
	"fmt"
)

// estructura de una oepraicon artimetica
type Aritmetica struct {
	LeftOperation   interfaces.Expresion
	SymbolOperation string
	RighthOpertion  interfaces.Expresion
}

// metodo constructor para construir una operacion aritmetica
func NewOperacion(leftOperation interfaces.Expresion, symbolOperation string, rightOperation interfaces.Expresion) Aritmetica {
	tempOperacion := Aritmetica{leftOperation, symbolOperation, rightOperation}
	return tempOperacion
}

// funcion para realizar operacion aritemticas suma
func (op Aritmetica) Ejecutar() interfaces.Simbolo {
	v := op.LeftOperation.Ejecutar().Valor.(int) + op.RighthOpertion.Ejecutar().Valor.(int)
	fmt.Println(v)
	// todos los casos de operacion
	return interfaces.Simbolo{Id: "", Valor: v, Mut: "", Tipo: 0}

}
