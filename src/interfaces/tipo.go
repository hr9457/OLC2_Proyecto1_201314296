package interfaces

// variable para saber que tipo de expresion se esta manejando
type TipoExpression int

/*
-enumeracion para los varios de tipo de error
- intenger = 0
- float = 1
- boolean = 2
- char = 3
- string = 4
- null = 5
*/
const (
	INTEGER TipoExpression = iota
	FLOAT
	BOOLEAN
	CHAR
	STRING
	NULL
	IDENTIFICADOR
)
