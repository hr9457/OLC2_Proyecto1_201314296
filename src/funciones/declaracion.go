package funciones

// estructura para guardado de la funcion
type BlockFuncion struct {
	Id        string
	Contenido interface{} // contendio de instruccione de un funcion
}

func NewFuncion(id string, contenido interface{}) *BlockFuncion {
	return &BlockFuncion{Id: id, Contenido: contenido}
}
