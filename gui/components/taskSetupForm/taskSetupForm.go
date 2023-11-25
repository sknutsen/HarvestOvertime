package taskSetupForm

import (
	"HarvestOvertime/constants"
	"HarvestOvertime/logic"

	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/sknutsen/harvestovertimelib"
	"github.com/sknutsen/harvestovertimelib/models"
)

var Tasks []models.Task = []models.Task{}
var SelectedTasks []models.Task = []models.Task{}

var SelectedTask models.Task

func CreateTaskSetupForm(client *http.Client, window fyne.Window) fyne.CanvasObject {
	settings, err := logic.ReadDetailsFromFile()
	if err != nil {
		println(err.Error())
	}

	Tasks = settings.TimeOffTasks
	SelectedTasks = settings.TimeOffTasks

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

	getTasksButton := widget.NewButton(constants.GetTasksButtonText, func() {
		settings, err := logic.ReadDetailsFromFile()
		if err != nil {
			println(err.Error())
		}

		newTasks, err := harvestovertimelib.ListTasks(client, settings)
		if err != nil {
			dialog.ShowError(err, window)
		}

		Tasks = newTasks

		taskSelect.Refresh()
	})

	addSelectedButton := widget.NewButton(constants.AddSelectedButtonText, func() {
		settings, err := logic.ReadDetailsFromFile()
		if err != nil {
			println(err.Error())
		}

		SelectedTasks = append(SelectedTasks, SelectedTask)

		settings.TimeOffTasks = SelectedTasks
		err = logic.SaveDetailsToFile(settings)
		if err != nil {
			println(err.Error())
		}

		list.Refresh()
	})

	clearSelectedButton := widget.NewButton(constants.ClearSelectedButtonText, func() {
		settings, err := logic.ReadDetailsFromFile()
		if err != nil {
			println(err.Error())
		}

		SelectedTasks = []models.Task{}

		settings.TimeOffTasks = SelectedTasks
		err = logic.SaveDetailsToFile(settings)
		if err != nil {
			println(err.Error())
		}

		list.Refresh()
	})

	hbox := container.New(layout.NewHBoxLayout(), getTasksButton, addSelectedButton, clearSelectedButton)
	borderLayout := container.NewBorder(hbox, nil, nil, nil, taskSelect)

	form := container.NewHSplit(list, borderLayout)

	return form
}
