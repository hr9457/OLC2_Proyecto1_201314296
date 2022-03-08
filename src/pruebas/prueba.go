package pruebas

import (
	"Proyecto1/src/environment"
	"Proyecto1/src/interfaces"

	arrayList "github.com/colegno/arraylist"
)

func Probar(lista *arrayList.List) {
	var entornoPrueba environment.Entornos
	entornoPrueba = environment.NewEntorno(nil, "global")
	//retorno la lista
	for _, s := range lista.ToArray() {
		s.(interfaces.Instruction).Ejecutar(entornoPrueba)
	}
	//fmt.Println("Cantidad de instrucciones:", lista.Len())
}
