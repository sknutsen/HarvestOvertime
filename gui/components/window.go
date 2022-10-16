package components

import (
	"HarvestOvertime/constants"
	"HarvestOvertime/gui/components/mainForm"
	"HarvestOvertime/gui/components/settingsForm"
	tasksetupform "HarvestOvertime/gui/components/taskSetupForm"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func CreateWindow(client *http.Client, app fyne.App) fyne.Window {
	win := app.NewWindow(constants.TitleText)
	win.Resize(fyne.NewSize(800, 400))

	mainForm := mainForm.CreateMainForm(client, win)

	taskSetupForm := tasksetupform.CreateTaskSetupForm(client, win)

	settingsForm := settingsForm.CreateSettingsForm(win)

	tabs := container.NewAppTabs(
		container.NewTabItem(constants.MainTabText, mainForm),
		container.NewTabItem(constants.TaskSetupTabText, taskSetupForm),
		container.NewTabItem(constants.SettingsTabText, settingsForm),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	win.SetContent(tabs)

	return win
}
