package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Monitor struct {
		Types       []string `yaml:"types"`
		IncludeDirs []string `yaml:"includeDirs"`
		Events      []string `yaml:"events"`
		// convert to
		TypesMap       map[string]bool `yaml:"-"`
		IncludeDirsMap map[string]bool `yaml:"-"`
		DirsMap        map[string]bool `yaml:"-"`
		IncludeDirsRec map[string]bool `yaml:"-"`
	}
	Notifier struct {
		PublishMqUrl string `yaml:"publish_mq_url"`
	}
}

func LoadConfig(configFile string, result interface{}) {
	//加载
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	viper.Unmarshal(&result)
}
