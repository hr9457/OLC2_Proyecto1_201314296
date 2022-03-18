package funciones

// estructura de un global
type Global struct {
	ContenidoGlobal interface{}
}

// constructor para global
func NewGlobal(contenido interface{}) Global {
	tempGlobal := Global{contenido}
	return tempGlobal
}

// ejecucion del meotod global
// func (global Global) Ejecutar(entorno interface{}) interface{} {
// 	if global.ContenidoGlobal != nil {

// 	} else {
// 		return 0
// 	}
// 	return 0
// }
