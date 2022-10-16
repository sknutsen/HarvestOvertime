package gui

import (
	"HarvestOvertime/gui/components"
	"net/http"

	"fyne.io/fyne/v2/app"
)

func GetGui(client *http.Client) {
	a := app.New()
	win := components.CreateWindow(client, a)
	win.ShowAndRun()
}
