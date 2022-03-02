package interfaces

import (
	"fmt"
	"reflect"
	"strconv"
)

// funcion para  hacer para convertir de string  un entero
func ConvertTextInt(valor string) int {
	temp, err := strconv.Atoi(valor)
	if err != nil {
		fmt.Println("Erro Semantico: tipo pirmitivo no existe", err)
	}
	fmt.Println(reflect.TypeOf(temp))
	return temp
}

// funcion para convertir de un string a un float
func ConvertTextFloat(valor string) float64 {
	temp, err := strconv.ParseFloat(valor, 64)
	if err != nil {
		fmt.Println("Erro Semantico: tipo pirmitivo no existe", err)
	}
	fmt.Println(reflect.TypeOf(temp))
	return temp
}

//funcion para convertir de un string a booleano
