package settingsForm

import (
	"HarvestOvertime/constants"
	"HarvestOvertime/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateSettingsForm(window fyne.Window) *fyne.Container {
	accountIdInput := widget.NewEntry()
	accountIdInput.SetPlaceHolder("Enter Harvest account ID...")

	accessTokenInput := widget.NewEntry()
	accessTokenInput.SetPlaceHolder("Enter Harvest access token...")

	button := widget.NewButton(constants.SaveDetailsButtonText, func() {
		err := logic.SaveDetailsToFile(accountIdInput.Text, accessTokenInput.Text)
		if err != nil {
			dialog.ShowError(err, window)
		}
	})

	accountId, accessToken, err := logic.ReadDetailsFromFile()
	if err != nil {
		dialog.ShowError(err, window)
	} else {
		accountIdInput.SetText(accountId)
		accessTokenInput.SetText(accessToken)
	}

	form := container.New(layout.NewVBoxLayout(), accountIdInput, accessTokenInput, button)

	return form
}
