package logic

import (
	"github.com/spf13/viper"
)

func SaveDetailsToFile(accountId string, accessToken string) error {
	var vp *viper.Viper = GetViperConfig()

	vp.Set("accountId", accountId)
	vp.Set("accessToken", accessToken)

	err := vp.WriteConfig()

	return err
}
