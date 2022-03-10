package interfaces

import (
	"fmt"
	"strconv"
)

// funcion para  hacer para convertir de string  un entero
func ConvertTextInt(valor string) int {
	temp, err := strconv.Atoi(valor)
	if err != nil {
		fmt.Println("Erro Semantico: tipo pirmitivo no existe", err)
	}
	// fmt.Println(reflect.TypeOf(temp))
	return temp
}

// funcion para convertir de un string a un float
func ConvertTextFloat(valor string) float64 {
	temp, err := strconv.ParseFloat(valor, 64)
	if err != nil {
		fmt.Println("Erro Semantico: tipo pirmitivo no existe", err)
	}
	// fmt.Println(reflect.TypeOf(temp))
	return temp
}

//funciones para extraer un los string y los char
func ConvertTextString(valor string) string {
	if (valor[0] == 34 && valor[len(valor)-1] == 34) || (valor[0] == 39 && valor[len(valor)-1] == 39) {
		// fmt.Println("Cumple con los paramentrso de un string/char")
		valor := valor[1 : len(valor)-1]
		return valor
	} else {
		// fmt.Println("PRIMITIVO:  ERROR SEMANTICO-> string/char")
		return "NULL"
	}
}

func ConvertTextBool(valor string) bool {
	temValor,_ := strconv.ParseBool(valor)
	return temValor
}
