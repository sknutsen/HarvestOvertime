package mainForm

import (
	"HarvestOvertime/constants"
	"HarvestOvertime/logic"
	"fmt"
	"image/color"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateMainForm(client *http.Client, window fyne.Window) *fyne.Container {
	title := canvas.NewText(constants.TitleText, color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	hoursText := widget.NewLabel("")
	hoursText.Alignment = fyne.TextAlignCenter
	hoursText.Wrapping = fyne.TextWrapWord

	button := widget.NewButton(constants.GetHoursButtonText, func() {
		entries, err := logic.ListEntries(client)
		if err != nil {
			println("Error: " + err.Error())

			dialog.ShowError(err, window)
		}

		fmt.Printf("Number of entries: %d\n", len(entries.TimeEntries))

		hours := fmt.Sprint(logic.GetTotalOvertime(entries))
		fmt.Printf("Overtime: %s\n", hours)

		hoursText.SetText(hours + " hours of overtime")
	})

	hboxButton := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	vbox := container.New(layout.NewVBoxLayout(), title, hboxButton, widget.NewSeparator(), hoursText)

	return vbox
}
