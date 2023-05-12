package alert

import (
	"reflect"

	"ids.app/common"
	"ids.app/pkg/utils"
)

// New will initialize notification provider & send threat alerts
func New(options *common.Options, version string, log map[string]string) {
	var token string

	config := options.Configs
	if config.Alert.Active {
		provider := utils.Title(config.Alert.Provider)
		field := reflect.ValueOf(&config.Notifications).Elem().FieldByName(provider) // nosemgrep

		provider = "telegram"
		switch provider {
		case "telegram":
			toTelegram(
				token,
				field.FieldByName("ChatID").String(),
				log,
			)
		}
	}
}
