package main

import (
	"fmt"

	"Proyecto1/parser"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type rustListener struct {
	*parser.BaserustListener
}

func main() {

	// simbolo := tableSymbol.NewSimbolo("a", "a", false, 0)
	// fmt.Println(simbolo.GetTipo())

	a := app.New()
	window := a.NewWindow("OLC2")

	var ancho, alto float32 = 500, 500
	var btnPosicionX, btnPosicionY float32 = 460, 30
	var inicioX, inicioY float32 = 20, 100

	// tamanio de la ventana
	window.Resize(fyne.NewSize(1150, 700))

	// etiquetas de titulo
	lbEntrada := widget.NewLabel("Entrada de datos")
	lbEntrada.Move(fyne.NewPos(inicioX, inicioY))
	lbSalida := widget.NewLabel("Salida de datos")
	lbSalida.Move(fyne.NewPos(inicioX+600, inicioY))

	// entrada para analizar
	txtEntrada := widget.NewMultiLineEntry()
	txtEntrada.SetPlaceHolder("Ingrese su codigo.....")
	txtEntrada.Resize(fyne.NewSize(ancho, alto))
	txtEntrada.Move(fyne.NewPos(inicioX, 150))

	// salida del analizador
	txtSalida := widget.NewMultiLineEntry()
	txtSalida.SetPlaceHolder("Salida..........")
	txtSalida.Resize(fyne.NewSize(ancho, alto))
	txtSalida.Move(fyne.NewPos(inicioX+600, 150))

	// bonton de compilacion
	btnRun := widget.NewButton("Compilar", func() {
		fmt.Println("****************Comilando!!!!******************")
		data := txtEntrada.Text
		is := antlr.NewInputStream(data)

		// creacion del analizador lexico
		lexer := parser.NewrustLexer(is)
		// lecutra de los token del analizador
		// for {
		// 	t := lexer.NextToken()
		// 	if t.GetTokenType() == antlr.TokenEOF {
		// 		break
		// 	}
		// 	fmt.Printf("%s (%q)\n",
		// 		lexer.SymbolicNames[t.GetTokenType()], t.GetText())
		// }
		// creando analizador sintactico
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewrustParser(stream)
		antlr.ParseTreeWalkerDefault.Walk(&rustListener{}, p.Start())

		txtSalida.SetText(data)
	})
	btnRun.Resize(fyne.NewSize(100, 50))
	btnRun.Move(fyne.NewPos(btnPosicionX, btnPosicionY))

	// boton de reporte
	btnReport := widget.NewButton("Reporte", func() {
		fmt.Println("Reporte")
	})
	btnReport.Resize(fyne.NewSize(100, 50))
	btnReport.Move(fyne.NewPos(btnPosicionX+120, btnPosicionY))

	// agregacion de contenedores a la ventana
	window.SetContent(
		container.NewWithoutLayout(
			lbEntrada,
			lbSalida,
			btnRun,
			btnReport,
			txtEntrada,
			txtSalida,
		),
	)

	// visibilidad
	window.Show()
	a.Run()
}
