package logic

import (
	"HarvestOvertime/logic/models"

	"github.com/spf13/viper"
)

func ReadDetailsFromFile() (string, string, error) {
	var vp *viper.Viper = GetViperConfig()
	var appSettings models.AppSettings

	err := vp.ReadInConfig()
	if err != nil {
		return "", "", err
	}

	err = vp.Unmarshal(&appSettings)
	if err != nil {
		return "", "", err
	}

	return appSettings.AccountId, appSettings.AccessToken, nil
}
