package instrucciones

import (
	"Proyecto1/src/environment"
	"Proyecto1/src/interfaces"
	"fmt"
	"reflect"

	arrayList "github.com/colegno/arraylist"
)

type Loop struct {
	Contenido interface{} // contenido
}

// metodo para cconstruir el Loop
func NewLoop(contenido interface{}) Loop {
	tempLoop := Loop{contenido}
	return tempLoop
}

func (firmaLoop Loop) Ejecutar(entorno interface{}) interface{} {
	// ejecuccion del for
	// guardadfo y evalucacion de la expresion del loop

	// creacion del entorno del loop
	var entornoLoop environment.Entornos
	entornoLoop = environment.NewEntorno(entorno.(environment.Entornos), "Etorno Loop")
	// ejecuccion en un loop
	for {
		// ejecucion de instruccion por instrucuon dentro del loop
		for _, s := range firmaLoop.Contenido.(*arrayList.List).ToArray() {
			// ejecutar instruccion
			// creacion de una variable para compara con los typeof si es un sentencia break
			b := BreakExp{Valor: nil, Tipo: interfaces.BREAK}
			// si es igual a una sentencia break se corrompe
			if reflect.TypeOf(s) == reflect.TypeOf(b) {
				fmt.Println("Se econtro un break")
				return nil
			}
			ejecucion := s.(interfaces.Instruction).Ejecutar(entornoLoop)
			if reflect.TypeOf(ejecucion) == reflect.TypeOf(b) {
				return nil
			}
		}
	}
}
