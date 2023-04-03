package logic

import (
	"HarvestOvertime/logic/models"
	"fmt"
)

func GetTotalOvertime(entries models.TimeEntries) float64 {
	var sum float64 = 0.0
	var overtime float64 = 0.0
	var dates []string = []string{}

	filteredList := filterTimeOff(entries)

	for i := 0; i < len(filteredList); i++ {
		sum = filteredList[i].Hours + sum

		dates = appendDate(dates, filteredList[i].SpentDate)
	}

	fmt.Printf("Number of dates: %d\n", len(dates))

	overtime = sum - (float64(len(dates)) * 7.5)

	return addCarryOver(overtime)
}

func appendDate(dates []string, date string) []string {
	var exists bool = false

	for i := 0; i < len(dates); i++ {
		if dates[i] == date {
			exists = true
			break
		}
	}

	if !exists {
		dates = append(dates, date)
	}

	return dates
}

func addCarryOver(overtime float64) float64 {
	settings, err := InitSettingsFromFile()
	if err != nil {
		return overtime
	}

	return overtime + settings.CarryOverTime
}

func filterTimeOff(entries models.TimeEntries) []models.TimeEntry {
	var filteredList []models.TimeEntry = []models.TimeEntry{}

	settings, err := InitSettingsFromFile()
	if err != nil {
		return entries.TimeEntries
	}

	for i := 0; i < len(entries.TimeEntries); i++ {
		exists := false

		for j := 0; j < len(settings.TimeOffTasks); j++ {
			if entries.TimeEntries[i].Task.Id == settings.TimeOffTasks[j].Id {
				exists = true
				break
			}
		}

		if !exists {
			filteredList = append(filteredList, entries.TimeEntries[i])
		}
	}

	return filteredList
}
