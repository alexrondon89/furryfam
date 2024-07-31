package cloud

import (
	"fmt"
	"github.com/alexrondon89/furryfam/infrastructure/config"
	"github.com/alexrondon89/furryfam/infrastructure/internal/cloud/aws"
	"github.com/spf13/viper"
	"log"
)

type CloudConfig struct {
	AwsInst *aws.AwsInstance
}

func GetCloudConfig(configName, configPath, configType string) config.PlatformConfig {
	conf := GetConfig(configName, configPath, configType)
	return conf
}

func GetCloudInstance(platformConfig config.PlatformConfig) CloudConfig {
	var cloudConfig CloudConfig
	if platformConfig.Aws != nil {
		cloudConfig.AwsInst = aws.CreateAwsInstance(platformConfig)
	}
	return cloudConfig
}

func GetConfig(configName, configPath, configType string) config.PlatformConfig {
	conf := &config.PlatformConfig{}
	viper.SetConfigName(configName)
	viper.AddConfigPath(fmt.Sprintf("./infrastructure/internal/cloud/%s", configPath))
	viper.SetConfigType(configType)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(conf)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return *conf
}
