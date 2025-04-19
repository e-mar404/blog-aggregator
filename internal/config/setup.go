package config

import (
	"encoding/json"
	"io"
	"os"
)

const configFile = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	configPath, err := configFilePath()
	if err != nil {
		return &Config{}, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return &Config{}, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return &Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return &Config{}, err
	}

	return &config, nil
}

func (c Config) SetUser(name string) error {
	c.CurrentUserName = name
	if err := write(c); err != nil {
		return err
	}
	return nil
}

func configFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	
	return home + "/" + configFile, nil
}

func write(c Config) error {
	raw, err := json.Marshal(c)
	if err != nil {
		return err
	}
	path, err := configFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(path, raw, 0644)
	if err != nil {
		return err
	}
	return nil
}
