package environment

import (
	"Proyecto1/src/interfaces"
	"fmt"
)

// estructar de los entornos
type Entornos struct {
	Padre     interface{}
	Nombre    interface{}
	Variables map[string]interfaces.Simbolo
	// Funciones   interface{}
	// Estructuras interface{}
}

// metodo constructor para entornos
// otra forma de declarar el metodo constructor para un objeto en golang
func NewEntorno(padre interface{}, nombre interface{}) Entornos {
	// pasa por parametros quien es el padre, nombre y inicializa el listado de los simbolos
	nuevoEntorno := Entornos{padre, nombre, make(map[string]interfaces.Simbolo)}
	return nuevoEntorno
}

/*funciones de retorno de busqueda para
los listados de variables, funciones y estructuras puesta en los ambitos
*/
// funcion para agregar variables
func (entorno Entornos) AddVariable(id string, valor interfaces.Simbolo, mut bool, tipo interfaces.TipoExpression) {
	// busqueda de la variable en el entorno
	variable, ok := entorno.Variables[id]
	if ok {
		fmt.Println("ETORNO: ERROR->Variable ya existe en el entorno: ", variable)
		return
	}
	// agrega un nuevo valor ala lista ->
	// fmt.Println("ENTORNO:  ->Agreacion de una nueva variable en el entorno")
	entorno.Variables[id] = interfaces.Simbolo{Id: id, Valor: valor, Mut: mut, Tipo: tipo}
	// fmt.Println("ENTORNO:  valores de la variable", entorno.Variables[id])
}

//funcion para retornar variable
func (entorno Entornos) GetVariables(id string) interfaces.Simbolo {
	var tempEntorno Entornos
	tempEntorno = entorno
	for {
		if variable, ok := tempEntorno.Variables[id]; ok {
			// fmt.Println("ENTORNO:  El variable que se retorna es->", variable)
			return variable
		}
		if tempEntorno.Padre == nil {
			break
		} else {
			tempEntorno = tempEntorno.Padre.(Entornos)
		}
	}
	fmt.Println("ENTORNO:  ERROR->Variable no existe")
	return interfaces.Simbolo{Id: "", Tipo: interfaces.NULL, Valor: interfaces.Simbolo{Id: "", Tipo: interfaces.NULL, Valor: nil}}
}

//funcion para retorna el valor de una variable
func (entorno Entornos) AlterVariable(id string, valor interfaces.Simbolo) interfaces.Simbolo {
	var tempEntorno Entornos
	tempEntorno = entorno
	// forpara buscar
	for {
		if variable, ok := tempEntorno.Variables[id]; ok {
			// verificacion de la mutabilidad de una variable para ser alterada
			// fmt.Println("ENTORNO:  variable a alatera: ", tempEntorno.Variables[id])
			if tempEntorno.Variables[id].Mut == true {
				// fmt.Println("ENTORNO:  mutando variable")
				tempEntorno.Variables[id] = interfaces.Simbolo{Id: id, Valor: valor, Mut: variable.Mut, Tipo: variable.Tipo}
				return variable
			} else {
				fmt.Println("ENTORNO:  ERROR-> variable no es mutable", tempEntorno.Variables[id].Mut)
				return tempEntorno.Variables[id]
			}
		}
		if tempEntorno.Padre == nil {
			break
		} else {
			tempEntorno = tempEntorno.Padre.(Entornos)
		}
	}
	//la variable no se encontro en el arreglo de la variables
	fmt.Println("ENTORNO: ERROR->La variable no existe en el entorno")
	//retornamos un simbolo con datos por defecto
	return interfaces.Simbolo{Id: "", Valor: interfaces.Simbolo{Id: "", Tipo: interfaces.NULL, Valor: 0}, Mut: false, Tipo: interfaces.NULL}

}

// get y set para padre, nombre
func (entornos *Entornos) GetPadre() interface{} {
	return entornos.Padre
}

func (entornos *Entornos) SetPadre(padre interface{}) {
	entornos.Padre = padre
}

func (entornos *Entornos) GetNombre() interface{} {
	return entornos.Nombre
}

func (entornos *Entornos) SetEntorno(nombre interface{}) {
	entornos.Nombre = nombre
}

// func (entornos *Entornos) AddSymbol() {
// 	s := NewSimbolo("a", "a", false, 0)
// 	fmt.Println(s.GetId())
// }
