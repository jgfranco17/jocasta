package environment

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	variablePrefix string = "JOCASTA"
)

type Configurations struct {
	EnvLogLevel string `mapstructure:"JOCASTA_LOG_LEVEL"`
}

func (c *Configurations) LogLevel() string {
	return c.EnvLogLevel
}

func LoadConfigs() (*Configurations, error) {
	config := Configurations{}
	viper.SetEnvPrefix(variablePrefix)
	envVars := []string{"JOCASTA_LOG_LEVEL"}
	for _, variable := range envVars {
		viper.BindEnv(variable, variable)
	}
	viper.AutomaticEnv()
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse configuration: %w", err)
	}
	return &config, nil
}
