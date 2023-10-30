package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"thibautdepond.github.io/platano-goDice/utils"
)

func main() {
	a := app.New()
	w := a.NewWindow("goDice")

	dice := widget.NewLabel("Dice :")
	diceResult := widget.NewLabel("")

	w.SetContent(container.NewVBox(
		dice,
		diceResult,
		widget.NewButton("Roll!", func() {
			diceResult.SetText(strconv.Itoa(utils.Roll()))
		}),
	))

	w.ShowAndRun()
}
