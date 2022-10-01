package support

import "grock/src/foundation"

type ConfigRepository struct {
	App foundation.Application
}

var Config *ConfigRepository

func init() {
	Config = &ConfigRepository{}
}
