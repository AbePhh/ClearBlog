package core

import (
	"log"
	"server/config"
	"server/utils"

	"gopkg.in/yaml.v3"
)

// 从 config.yaml读取配置
func InitConf() *config.Config {
	c := &config.Config{}
	yamlConf, err := utils.LoadYAML()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v ", err)
	}
	if err := yaml.Unmarshal(yamlConf, c); err != nil {
		log.Fatalf("Failed to unmarshal configuration: %v ", err)
	}
	return c
}
