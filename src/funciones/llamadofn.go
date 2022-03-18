package funciones

// estructura para el llamado a funcion
type LlamadoFn struct {
	Id        string
	Contenido interface{}
}

// constructor
func NewLlamdoFn(id string, contenido interface{}) LlamadoFn {
	temp := LlamadoFn{id, contenido}
	return temp
}
