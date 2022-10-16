package logic

import (
	"encoding/json"
	"net/http"

	"HarvestOvertime/logic/models"
)

func ListEntries(client *http.Client) (models.TimeEntries, error) {
	var entries models.TimeEntries
	var url string = "https://api.harvestapp.com/api/v2/time_entries"
	var counter int32 = 0

	for url != "" {
		newEntries, err := listEntries(client, url)
		if err == nil {
			url = newEntries.Links.Next
			println("New url: " + newEntries.Links.Next)

			entries.TimeEntries = append(entries.TimeEntries, newEntries.TimeEntries...)
		}

		counter++

		println(counter)
	}

	return entries, nil
}

func listEntries(client *http.Client, url string) (models.TimeEntries, error) {
	var entries models.TimeEntries

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		println("Error creating request: " + err.Error())
		return models.TimeEntries{}, err
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
		return models.TimeEntries{}, err
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
		return models.TimeEntries{}, err
	}

	return entries, nil
}
