package utils

import (
	"io/fs"
	"os"
	"server/global"

	"gopkg.in/yaml.v3"
)

// 定义全局配置文件
const configFile = "config.yaml"

// 加载配置文件
func LoadYAML() ([]byte, error) {
	return os.ReadFile(configFile)
}

func SaveYAML() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}

	return os.WriteFile(configFile, byteData, fs.ModePerm)
}
