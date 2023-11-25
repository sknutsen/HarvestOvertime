package logic

import (
	"github.com/sknutsen/harvestovertimelib/models"
	"github.com/spf13/viper"
)

func ReadDetailsFromFile() (models.Settings, error) {
	var settings models.Settings
	var vp *viper.Viper = GetViperConfig()

	err := vp.ReadInConfig()
	if err != nil {
		err = SaveDetailsToFile(settings)

		return models.Settings{}, err
	}

	err = vp.Unmarshal(&settings)
	if err != nil {
		return models.Settings{}, err
	}

	return settings, nil
}
