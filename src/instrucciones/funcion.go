package instrucciones

type BlockFuncion struct {
	Identificador string
	// Contenido     *arrayList.List // contendio de instruccione de un funcion
}

func NewFuncion(id string) BlockFuncion {
	tempFuncion := BlockFuncion{id}
	return tempFuncion
}

// func (firmaFuncion BlockFuncion) Ejecutar(entorno interface{}) interface{} {

// 	fmt.Println("FUNCION:   retorno->")

// 	// var temp interfaces.Simbolo
// 	// temp := interfaces.Simbolo{Id: "funcion"}
// 	return 0
// }
