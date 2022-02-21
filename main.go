package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

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
		fmt.Println("Comilando!!!!")
		data := txtEntrada.Text
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
