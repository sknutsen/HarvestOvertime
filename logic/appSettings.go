package logic

import "HarvestOvertime/logic/models"

type AppSettings struct {
	AccessToken   string        `json:"accessToken"`
	AccountId     string        `json:"accountId"`
	CarryOverTime float64       `json:"carryOverTime"`
	CurrentYear   bool          `json:"CurrentYear"`
	TimeOffTasks  []models.Task `json:"timeOffTasks"`
}
