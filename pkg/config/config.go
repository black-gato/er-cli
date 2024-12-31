package config

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Repos []Repo `yaml:"repos"`
}
type Repo struct {
	Name        string `yaml:"name"`
	RepoURL     string `yaml:"repoURL"`
	Description string `yaml:"description"`
	Path        string `yaml:"path"`
}

func checkFileExists(path string) bool {
	_, error := os.Stat(path)
	return !errors.Is(error, os.ErrNotExist)
}
func NewConfig(dir string) Config {
	return Config{
		Repos: []Repo{
			{Name: "example-repo", RepoURL: "https://github.com/black-gato/err-cli-example.git", Description: "This is a sample repo to show how this works", Path: dir}}}
}

func CreateConfigDir(dir string) (string, error) {
	var err error
	if dir == "" {
		dir, err = os.UserHomeDir()
		if err != nil {
			return "", err
		}
	}

	errDir := filepath.Join(dir, ".err-cli")
	if checkFileExists(errDir) {
		return "", errors.New("err-cli config folder already exists")
	}
	err = os.Mkdir(errDir, 0750)
	if err != nil && !os.IsExist(err) {
		return "", err
	}
	return errDir, nil
}
func (c *Config) WriteConfigFile(dir string) (path string, err error) {
	configPath := filepath.Join(dir, "default.yaml")
	if checkFileExists(configPath) {
		return "", errors.New("err-cli config file already exists")
	}

	b, err := yaml.Marshal(&c)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(configPath, b, 0644)
	if err != nil {
		return "", err
	}
	return configPath, nil

}
