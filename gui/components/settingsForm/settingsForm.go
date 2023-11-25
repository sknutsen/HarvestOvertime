package settingsForm

import (
	"HarvestOvertime/constants"
	"HarvestOvertime/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/sknutsen/harvestovertimelib/models"
)

func CreateSettingsForm(window fyne.Window) *fyne.Container {
	var appSettings models.Settings

	accountIdInput := widget.NewEntry()
	accountIdInput.SetPlaceHolder("Enter Harvest account ID...")

	accessTokenInput := widget.NewEntry()
	accessTokenInput.SetPlaceHolder("Enter Harvest access token...")

	appSettings, err := logic.ReadDetailsFromFile()
	if err != nil {
		dialog.ShowError(err, window)
	} else {
		accountIdInput.SetText(appSettings.AccountId)
		accessTokenInput.SetText(appSettings.AccessToken)
	}

	button := widget.NewButton(constants.SaveDetailsButtonText, func() {
		updatedAppSettings, err := logic.ReadDetailsFromFile()
		if err != nil {
			dialog.ShowError(err, window)
		}

		updatedAppSettings.AccessToken = accessTokenInput.Text
		updatedAppSettings.AccountId = accountIdInput.Text

		err = logic.SaveDetailsToFile(updatedAppSettings)
		if err != nil {
			dialog.ShowError(err, window)
		}
	})

	form := container.New(layout.NewVBoxLayout(),
		accountIdInput,
		accessTokenInput,
		button,
		layout.NewSpacer())

	return form
}
