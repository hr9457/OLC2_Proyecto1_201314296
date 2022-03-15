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
- identificador = 6
- logico
*/
const (
	INTEGER       TipoExpression = iota //0
	FLOAT                               //1
	BOOLEAN                             //2
	CHAR                                //3
	STRING                              //4
	NULL                                //5
	IDENTIFICADOR                       //6
	LOGICO                              //7
	BREAK                               //8
)
