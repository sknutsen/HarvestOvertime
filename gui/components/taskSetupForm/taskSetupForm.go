package tasksetupform

import (
	"HarvestOvertime/constants"
	"HarvestOvertime/logic"
	"HarvestOvertime/logic/models"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var Tasks []models.Task = []models.Task{}
var SelectedTasks []models.Task = []models.Task{}

var SelectedTask models.Task

// var data = []string{"a", "string", "list"}

func CreateTaskSetupForm(client *http.Client, window fyne.Window) *fyne.Container {
	taskSelect := widget.NewList(
		func() int {
			return len(Tasks)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(Tasks[i].Name)
		},
	)

	taskSelect.OnSelected = func(id widget.ListItemID) {
		SelectedTask = Tasks[id]
	}

	list := widget.NewList(
		func() int {
			return len(SelectedTasks)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(SelectedTasks[i].Name)
		},
	)

	list.Resize(fyne.NewSize(400, list.Size().Height))

	getTasksButton := widget.NewButton(constants.GetTasksButtonText, func() {
		newTasks, err := logic.ListTasks(client)
		if err != nil {
			dialog.ShowError(err, window)
		}

		Tasks = newTasks

		taskSelect.Refresh()
	})

	addSelectedButton := widget.NewButton(constants.AddSelectedButtonText, func() {
		SelectedTasks = append(SelectedTasks, SelectedTask)

		list.Refresh()
	})

	clearSelectedButton := widget.NewButton(constants.ClearSelectedButtonText, func() {
		SelectedTasks = []models.Task{}

		list.Refresh()
	})

	hbox := container.New(layout.NewHBoxLayout(), addSelectedButton, clearSelectedButton)
	vbox := container.New(layout.NewVBoxLayout(), getTasksButton, taskSelect, hbox)
	form := container.New(layout.NewHBoxLayout(), list, vbox)

	return form
}
