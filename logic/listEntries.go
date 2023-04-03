package logic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"HarvestOvertime/logic/models"
)

func ListEntries(client *http.Client) (models.TimeEntries, error) {
	var entries models.TimeEntries
	var url string = fmt.Sprintf("https://api.harvestapp.com/api/v2/time_entries?from=%d-01-01", time.Now().Year())
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

	settings, err := InitSettingsFromFile()
	if err != nil {
		println("Error reading settings: " + err.Error())
	}

	req.Header.Add("Harvest-Account-ID", settings.AccountId)
	req.Header.Add("Authorization", "Bearer "+settings.AccessToken)
	req.Header.Add("User-Agent", "Harvest API Example")

	resp, err := client.Do(req)
	if err != nil {
		println("Error sending request: " + err.Error())
		return models.TimeEntries{}, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&entries)
	if err != nil {
		println("Error decoding response: " + err.Error())
		return models.TimeEntries{}, err
	}

	return entries, nil
}
