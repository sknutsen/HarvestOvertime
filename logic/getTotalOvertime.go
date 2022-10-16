package logic

import (
	"HarvestOvertime/logic/models"
	"fmt"
)

func GetTotalOvertime(entries models.TimeEntries) float64 {
	var sum float64 = 0.0
	var overtime float64 = 0.0
	var dates []string = []string{}

	for i := 0; i < len(entries.TimeEntries); i++ {
		sum = entries.TimeEntries[i].Hours + sum

		dates = appendDate(dates, entries.TimeEntries[i].SpentDate)
	}

	fmt.Printf("Number of dates: %d\n", len(dates))

	overtime = sum - (float64(len(dates)) * 7.5)

	return overtime
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
