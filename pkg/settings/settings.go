package settings

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	UseServiceAccount bool     `envconfig:"USE_SERVICE_ACCOUNT" default:"false"`
	ExcludedRegistry  []string `envconfig:"EXCLUDED_REGISTRY"`
}

func NewSettings() Settings {
	var settings Settings

	err := envconfig.Process("", &settings)
	if err != nil {
		log.Fatalln(err)
	}

	return settings
}
