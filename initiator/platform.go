package initiator

import (
	"dating/platform"
	"dating/platform/logger"
	"dating/platform/sms"

	"github.com/spf13/viper"
)

type PlatformLayer struct {
	sms platform.SMSClient
}

func InitPlatformLayer(logger logger.Logger) PlatformLayer {
	return PlatformLayer{
		sms: sms.InitSMS(
			platform.SMSConfig{
				UserName:  viper.GetString("sms.username"),
				Password:  viper.GetString("sms.password"),
				Server:    viper.GetString("sms.server"),
				Type:      viper.GetString("sms.type"),
				DCS:       viper.GetString("sms.dcs"),
				DLRMask:   viper.GetString("sms.dlrmask"),
				DLRURL:    viper.GetString("sms.dlrurl"),
				Sender:    viper.GetString("sms.sender"),
				Templates: viper.GetStringMapString("sms.templates"),
				APIKey:    viper.GetString("sms.api_key"),
			},
			logger),
	}
}

func InitMockPlatformLayer(logger logger.Logger) PlatformLayer {
	return PlatformLayer{}
}
