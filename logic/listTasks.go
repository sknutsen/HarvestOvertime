package logic

import (
	"HarvestOvertime/logic/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func ListTasks(client *http.Client) ([]models.Task, error) {
	var tasks []models.Task
	var url string = "https://api.harvestapp.com/api/v2/tasks"
	var counter int32 = 0

	for url != "" {
		newEntries, err := listTasks(client, url)
		if err == nil {
			url = newEntries.Links.Next
			println("New url: " + newEntries.Links.Next)

			tasks = append(tasks, newEntries.Tasks...)
		}

		counter++

		println(counter)
	}

	fmt.Printf("Tasks: %d", len(tasks))

	return tasks, nil
}

func listTasks(client *http.Client, url string) (models.Tasks, error) {
	var entries models.Tasks

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		println("Error creating request: " + err.Error())
		return models.Tasks{}, err
	}

	id, token, err := ReadDetailsFromFile()
	if err != nil {
		println("Error reading settings: " + err.Error())
	}

	req.Header.Add("Harvest-Account-ID", id)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("User-Agent", "Harvest API Example")

	resp, err := client.Do(req)
	if err != nil {
		println("Error sending request: " + err.Error())
		return models.Tasks{}, err
	}

	defer resp.Body.Close()

	// content, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	println(err.Error())
	// }

	// println(string(content))

	err = json.NewDecoder(resp.Body).Decode(&entries)
	if err != nil {
		println("Error decoding response: " + err.Error())
		return models.Tasks{}, err
	}

	return entries, nil
}
