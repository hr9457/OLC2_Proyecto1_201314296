package symbol

import (
	"Proyecto1/src/interfaces"
	"fmt"
)

// creacion de la clase para los simbolos
// uso el := para declara, asignar y asignar el tipo que es
type Simbolo struct {
	Id    interface{}
	Valor interface{}
	Mut   interface{}
	Tipo  interfaces.TipoExpression
}

// metodo constructor del simbolo
func NewSimbolo(id interface{}, valor interface{}, mut interface{}, tipo interfaces.TipoExpression) *Simbolo {
	if tipo == 5 {
		fmt.Println("La variable declarada  no tiene un tipo definido")
		fmt.Println("->")
		fmt.Println(valor)
		fmt.Println("->")
	}
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

func (simbolo *Simbolo) SetTipo(tipo interfaces.TipoExpression) {
	simbolo.Tipo = tipo
}

// -> constructor
// l := Simbolo{Id:"",Valor;"",Mut:"",Tip:""}
