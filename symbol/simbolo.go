package symbol

// variable para saber que tipo de expresion se esta manejando
type TipoExpression int

/*
-enumeracion para los varios de tipo de error
- intenger = 0
- float = 1....
- .............
- null = 5
*/
const (
	INTERGER TipoExpression = iota
	FLOAT
	STRING
	BOOLEAN
	NULL
)

// creacion de la clase para los simbolos
// uso el := para declara, asignar y asignar el tipo que es
type Simbolo struct {
	Id    interface{}
	Valor interface{}
	Mut   interface{}
	Tipo  TipoExpression
}

// metodo constructor del simbolo
func NewSimbolo(id interface{}, valor interface{}, mut interface{}, tipo TipoExpression) *Simbolo {
	return &Simbolo{Id: id, Valor: valor, Mut: mut, Tipo: tipo}
}

// get y set
func (simbolo *Simbolo) GetId() interface{} {
	return simbolo.Id
}

func (simbolo *Simbolo) SetId(id interface{}) {
	simbolo.Id = id
}

func (simbolo *Simbolo) GetValor() interface{} {
	return simbolo.Valor
}

func (simbolo *Simbolo) SetValor(valor interface{}) {
	simbolo.Valor = valor
}

func (simbolo *Simbolo) GetMut() interface{} {
	return simbolo.Mut
}

func (simbolo *Simbolo) SetMut(mutable interface{}) {
	simbolo.Mut = mutable
}

func (simbolo *Simbolo) GetTipo() interface{} {
	return simbolo.Tipo
}

func (simbolo *Simbolo) SetTipo(tipo TipoExpression) {
	simbolo.Tipo = tipo
}

// -> constructor
// l := Simbolo{Id:"",Valor;"",Mut:"",Tip:""}
