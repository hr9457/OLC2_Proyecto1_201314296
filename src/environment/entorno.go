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
func (entorno *Entornos) AddVariable(id string, valor interface{}, mut bool, tipo interfaces.TipoExpression) {
	// busqueda de la variable en el entorno
	variable, ok := entorno.Variables[id]
	if ok {
		fmt.Println("Variable ya existe en el entorno: ", variable)
		return
	}
	// agrega un nuevo valor ala lista ->
	fmt.Println("Agreacion de una nueva variable en el entorno ->")
	entorno.Variables[id] = interfaces.Simbolo{Id: id, Valor: valor, Mut: mut, Tipo: tipo}
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
