package logic

import (
	"bufio"
	"os"

	"github.com/spf13/viper"
)

func (settings *AppSettings) SaveDetailsToFile() error {
	var vp *viper.Viper = GetViperConfig()

	vp.Set("accessToken", settings.AccessToken)
	vp.Set("accountId", settings.AccountId)
	vp.Set("carryOverTime", settings.CarryOverTime)
	vp.Set("currentYear", settings.CurrentYear)
	vp.Set("timeOffTasks", settings.TimeOffTasks)

	err := vp.WriteConfig()
	if err != nil {
		path := "appsettings.json"

		file, err := os.Create(path)
		if err != nil {
			return err
		}
		writer := bufio.NewWriter(file)

		_, err = writer.WriteString("{}")
		if err != nil {
			return err
		}

		return settings.SaveDetailsToFile()
	}

	return nil
}
