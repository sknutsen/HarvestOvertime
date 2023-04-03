package logic

import (
	"github.com/spf13/viper"
)

func InitSettingsFromFile() (AppSettings, error) {
	var settings AppSettings

	err := settings.ReadDetailsFromFile()

	return settings, err
}

func (settings *AppSettings) ReadDetailsFromFile() error {
	var vp *viper.Viper = GetViperConfig()

	err := vp.ReadInConfig()
	if err != nil {
		err = settings.SaveDetailsToFile()

		return err
	}

	err = vp.Unmarshal(&settings)
	if err != nil {
		return err
	}

	return nil
}
