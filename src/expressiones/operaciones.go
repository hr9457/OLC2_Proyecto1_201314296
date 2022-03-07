package expressiones

import (
	"Proyecto1/src/interfaces"
	"fmt"
	"math"
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

	//verificacion de tipos
	if op.LeftOperation.Ejecutar(entorno).Tipo == op.RighthOpertion.Ejecutar(entorno).Tipo {
		// comparamos el tipo de valor para operar

		switch op.LeftOperation.Ejecutar(entorno).Tipo {

		// cuando las operaciones son de tipo integer
		case interfaces.INTEGER:
			{
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
				case "%":
					{
						resultado = op.LeftOperation.Ejecutar(entorno).Valor.(int) % op.RighthOpertion.Ejecutar(entorno).Valor.(int)
					}
				case "==":
					{
						if op.LeftOperation.Ejecutar(entorno).Valor.(int) == op.RighthOpertion.Ejecutar(entorno).Valor.(int) {
							resultado = true
						} else {
							resultado = false
						}
					}
				case "!=":
					{
						if op.LeftOperation.Ejecutar(entorno).Valor.(int) != op.RighthOpertion.Ejecutar(entorno).Valor.(int) {
							resultado = true
						} else {
							resultado = false
						}
					}
				case ">":
					{
						if op.LeftOperation.Ejecutar(entorno).Valor.(int) > op.RighthOpertion.Ejecutar(entorno).Valor.(int) {
							resultado = true
						} else {
							resultado = false
						}
					}
				case "<":
					{
						if op.LeftOperation.Ejecutar(entorno).Valor.(int) < op.RighthOpertion.Ejecutar(entorno).Valor.(int) {
							resultado = true
						} else {
							resultado = false
						}
					}
				case ">=":
					{
						if op.LeftOperation.Ejecutar(entorno).Valor.(int) >= op.RighthOpertion.Ejecutar(entorno).Valor.(int) {
							resultado = true
						} else {
							resultado = false
						}
					}
				case "<=":
					{
						if op.LeftOperation.Ejecutar(entorno).Valor.(int) > op.RighthOpertion.Ejecutar(entorno).Valor.(int) {
							resultado = true
						} else {
							resultado = false
						}
					}
				case "&&":
					{
						fmt.Println("OPERACION:  ->", op.LeftOperation.Ejecutar(entorno))
						resultado = true
					}
				}
			}

			//cuando las operaciones son de tipo flotante
		case interfaces.FLOAT:
			{
				//creacon del swtich case para ver el tipo de operacion
				switch op.SymbolOperation {
				case "+":
					{
						resultado = op.LeftOperation.Ejecutar(entorno).Valor.(float64) + op.RighthOpertion.Ejecutar(entorno).Valor.(float64)
					}
				case "-":
					{
						resultado = op.LeftOperation.Ejecutar(entorno).Valor.(float64) - op.RighthOpertion.Ejecutar(entorno).Valor.(float64)
					}
				case "*":
					{
						resultado = op.LeftOperation.Ejecutar(entorno).Valor.(float64) * op.RighthOpertion.Ejecutar(entorno).Valor.(float64)
					}
				case "/":
					{
						resultado = op.LeftOperation.Ejecutar(entorno).Valor.(float64) / op.RighthOpertion.Ejecutar(entorno).Valor.(float64)
					}
				case "%":
					{
						resultado = math.Mod(op.LeftOperation.Ejecutar(entorno).Valor.(float64), op.RighthOpertion.Ejecutar(entorno).Valor.(float64))
					}
				case "==":
					{
						if op.LeftOperation.Ejecutar(entorno).Valor.(float64) == op.RighthOpertion.Ejecutar(entorno).Valor.(float64) {
							resultado = true
						} else {
							resultado = false
						}
					}
				case "!=":
					{
						if op.LeftOperation.Ejecutar(entorno).Valor.(float64) != op.RighthOpertion.Ejecutar(entorno).Valor.(float64) {
							resultado = true
						} else {
							resultado = false
						}
					}
				case ">":
					{
						if op.LeftOperation.Ejecutar(entorno).Valor.(float64) > op.RighthOpertion.Ejecutar(entorno).Valor.(float64) {
							resultado = true
						} else {
							resultado = false
						}
					}
				case "<":
					{
						if op.LeftOperation.Ejecutar(entorno).Valor.(float64) < op.RighthOpertion.Ejecutar(entorno).Valor.(float64) {
							resultado = true
						} else {
							resultado = false
						}
					}
				case ">=":
					{
						if op.LeftOperation.Ejecutar(entorno).Valor.(float64) >= op.RighthOpertion.Ejecutar(entorno).Valor.(float64) {
							resultado = true
						} else {
							resultado = false
						}
					}
				case "<=":
					{
						if op.LeftOperation.Ejecutar(entorno).Valor.(float64) > op.RighthOpertion.Ejecutar(entorno).Valor.(float64) {
							resultado = true
						} else {
							resultado = false
						}
					}
				}
			}

		}

	} else {
		fmt.Println("OPERACIONES:  Error Semantico-> operadores no son del mismo tipo")
		resultado = 0
	}

	//retorno del simbolo como tal
	// todos los casos de operacion
	return interfaces.Simbolo{Id: "", Valor: resultado, Mut: "", Tipo: 0}
}
