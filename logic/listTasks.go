package logic

import (
	"HarvestOvertime/logic/models"
	"fmt"
	"net/http"
)

func ListTasks(client *http.Client) ([]models.Task, error) {
	var tasks []models.Task
	var counter uint64 = 0

	var m = make(map[uint64]models.Task)
	newEntries, err := ListEntries(client)
	if err == nil {
		for _, e := range newEntries.TimeEntries {
			_, exists := m[e.Task.Id]

			if !exists {
				m[e.Task.Id] = e.Task
				tasks = append(tasks, e.Task)
			}
		}
	}

	counter++

	println(counter)

	fmt.Printf("Tasks: %d", len(tasks))

	return tasks, nil
}
