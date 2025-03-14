package expressiones

import (
	"Proyecto1/src/interfaces"
	"Proyecto1/src/traduccion"
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
func (op Aritmetica) Ejecutar(entorno interface{}, traductor *traduccion.Traductor) interfaces.Simbolo {
	var resultado interface{}
	var resultadoTipo interfaces.TipoExpression

	if op.LeftOperation == nil {
		// CREACION DE LA COMPUERTA NOT
		// fmt.Println("El operador de la izquierda no existe")
		if op.SymbolOperation == "!" {
			if op.RighthOpertion.Ejecutar(entorno, traductor).Tipo == interfaces.BOOLEAN {
				resultado = !op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(bool)
				resultadoTipo = interfaces.BOOLEAN
			}
		} else if op.SymbolOperation == "-" {
			if op.RighthOpertion.Ejecutar(entorno, traductor).Tipo == interfaces.INTEGER {
				resultado = -op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int)
				resultadoTipo = interfaces.INTEGER
			} else if op.RighthOpertion.Ejecutar(entorno, traductor).Tipo == interfaces.FLOAT {
				resultado = -op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64)
				resultadoTipo = interfaces.FLOAT
			}
		}

	} else {

		//verificacion de tipos
		if op.LeftOperation.Ejecutar(entorno, traductor).Tipo == op.RighthOpertion.Ejecutar(entorno, traductor).Tipo {
			// comparamos el tipo de valor para operar

			switch op.LeftOperation.Ejecutar(entorno, traductor).Tipo {

			// cuando las operaciones son de tipo integer
			case interfaces.INTEGER:
				{
					//creacon del swtich case para ver el tipo de operacion
					switch op.SymbolOperation {
					case "+":
						{
							resultado = op.LeftOperation.Ejecutar(entorno, traductor).Valor.(int) + op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int)
							resultadoTipo = interfaces.INTEGER
						}
					case "-":
						{
							resultado = op.LeftOperation.Ejecutar(entorno, traductor).Valor.(int) - op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int)
							resultadoTipo = interfaces.INTEGER
						}
					case "*":
						{
							resultado = op.LeftOperation.Ejecutar(entorno, traductor).Valor.(int) * op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int)
							resultadoTipo = interfaces.INTEGER
						}
					case "/":
						{
							if op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int) == 0 {
								resultado = 0
								resultadoTipo = interfaces.INTEGER
							} else {
								resultado = op.LeftOperation.Ejecutar(entorno, traductor).Valor.(int) / op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int)
								resultadoTipo = interfaces.INTEGER
							}
						}
					case "%":
						{
							resultado = op.LeftOperation.Ejecutar(entorno, traductor).Valor.(int) % op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int)
							resultadoTipo = interfaces.INTEGER
						}
					case "i64::pow":
						{
							primero := float64(op.LeftOperation.Ejecutar(entorno, traductor).Valor.(int))
							segundo := float64(op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int))
							resultado = int(math.Pow(primero, segundo))
							resultadoTipo = interfaces.INTEGER
						}

					case "==":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(int) == op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case "!=":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(int) != op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case ">":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(int) > op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case "<":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(int) < op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case ">=":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(int) >= op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case "<=":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(int) <= op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(int) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
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
							resultado = op.LeftOperation.Ejecutar(entorno, traductor).Valor.(float64) + op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64)
							resultadoTipo = interfaces.FLOAT
						}
					case "-":
						{
							resultado = op.LeftOperation.Ejecutar(entorno, traductor).Valor.(float64) - op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64)
							resultadoTipo = interfaces.FLOAT
						}
					case "*":
						{
							resultado = op.LeftOperation.Ejecutar(entorno, traductor).Valor.(float64) * op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64)
							resultadoTipo = interfaces.FLOAT
						}
					case "/":
						{
							if op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64) == 0 {
								resultado = 0
								resultadoTipo = interfaces.FLOAT
							} else {
								resultado = op.LeftOperation.Ejecutar(entorno, traductor).Valor.(float64) / op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64)
								resultadoTipo = interfaces.FLOAT
							}
						}
					case "%":
						{
							resultado = math.Mod(op.LeftOperation.Ejecutar(entorno, traductor).Valor.(float64), op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64))
							resultadoTipo = interfaces.FLOAT
						}
					case "f64::pow":
						{
							resultado = math.Pow(op.LeftOperation.Ejecutar(entorno, traductor).Valor.(float64), op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64))
							resultadoTipo = interfaces.INTEGER
						}
					case "==":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(float64) == op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case "!=":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(float64) != op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case ">":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(float64) > op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case "<":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(float64) < op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case ">=":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(float64) >= op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case "<=":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(float64) <= op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(float64) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					}
				}

			case interfaces.BOOLEAN:
				{
					//creacon del swtich case para ver el tipo de operacion
					switch op.SymbolOperation {
					case "==":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(bool) == op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(bool) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case "!=":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(bool) != op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(bool) {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case ">":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(bool) == true {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case "<":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(bool) == true {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case ">=":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(bool) == true {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case "<=":
						{
							if op.LeftOperation.Ejecutar(entorno, traductor).Valor.(bool) == true {
								resultado = true
							} else {
								resultado = false
							}
							resultadoTipo = interfaces.BOOLEAN
						}
					case "&&":
						{
							Left := op.LeftOperation.Ejecutar(entorno, traductor).Valor.(bool)
							Right := op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(bool)
							if Left == true && Right == true {
								resultado = true
								resultadoTipo = interfaces.BOOLEAN
							} else {
								resultado = false
								resultadoTipo = interfaces.BOOLEAN
							}
						}
					case "||":
						{
							Left := op.LeftOperation.Ejecutar(entorno, traductor).Valor.(bool)
							Right := op.RighthOpertion.Ejecutar(entorno, traductor).Valor.(bool)
							if Left == true && Right == true {
								resultado = true
								resultadoTipo = interfaces.BOOLEAN
							} else if Left == true && Right == false {
								resultado = true
								resultadoTipo = interfaces.BOOLEAN
							} else if Left == false && Right == true {
								resultado = true
								resultadoTipo = interfaces.BOOLEAN
							} else if Left == false && Right == false {
								resultado = false
								resultadoTipo = interfaces.BOOLEAN
							}
						}
					}
				}
			}

		} else {
			fmt.Println("OPERACIONES:  Error Semantico-> operadores no son del mismo tipo")
			resultado = 0
			resultadoTipo = interfaces.NULL
		}

	}
	//retorno del simbolo como tal
	// todos los casos de operacion
	tempSimbolo := interfaces.Simbolo{Id: "", Valor: resultado, Mut: "", Tipo: resultadoTipo}
	// fmt.Println("OPERACION:  retonro realizado->", tempSimbolo)
	return tempSimbolo
}
