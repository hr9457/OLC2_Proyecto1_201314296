package traduccion

// ------ captura de la traduccion -----------------
type Traductor struct {
	Contenido string
}

// contructor
func NewTraductor(contenido string) Traductor {
	temp := Traductor{contenido}
	return temp
}

// var traduccion Traduccion = Traduccion{Contenido: ""}

// --------------------------------------------------
