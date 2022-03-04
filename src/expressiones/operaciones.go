package expressiones

import (
	"Proyecto1/src/interfaces"
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
func (op Aritmetica) Ejecutar(entorno interface{}) interfaces.Simbolo {
	var resultado interface{}
	//creacon del swtich case para ver el tipo de operacion
	switch op.SymbolOperation {
	case "+":
		{
			resultado = op.LeftOperation.Ejecutar(entorno).Valor.(int) + op.RighthOpertion.Ejecutar(entorno).Valor.(int)
		}
	case "-":
		{
			resultado = op.LeftOperation.Ejecutar(entorno).Valor.(int) - op.RighthOpertion.Ejecutar(entorno).Valor.(int)
		}
	case "*":
		{
			resultado = op.LeftOperation.Ejecutar(entorno).Valor.(int) * op.RighthOpertion.Ejecutar(entorno).Valor.(int)
		}
	case "/":
		{
			resultado = op.LeftOperation.Ejecutar(entorno).Valor.(int) / op.RighthOpertion.Ejecutar(entorno).Valor.(int)
		}

	}
	//retorno del simbolo como tal
	// todos los casos de operacion
	return interfaces.Simbolo{Id: "", Valor: resultado, Mut: "", Tipo: 0}
}
