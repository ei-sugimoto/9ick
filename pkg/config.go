package config

import (
	"fmt"
	"os"
	"path/filepath"
)

type IConfig interface {
	GetRootDir() (string, error)
}

type Config struct{}

func NewConfig() IConfig {
	return &Config{}
}

func (c Config) GetRootDir() (string, error) {
	executable, err := os.Executable()
	if err != nil {
		fmt.Println("error :", err)
		return "", err
	}

	execDir := filepath.Dir(executable)

	rootDir := filepath.Dir(execDir)

	return rootDir, nil
}
