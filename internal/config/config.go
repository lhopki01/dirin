package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	yaml "gopkg.in/yaml.v3"
)

func configDir() string {
	appName := "dirin"
	if xdgDir := os.Getenv("XDG_CONFIG_HOME"); xdgDir != "" {
		return path.Join(xdgDir, appName)
	}
	// Default to `~/.kk` rather than `~/.config/kk` to match the kubectl default of `~/.kube`
	return path.Join(homeDir(), ".config", appName)
}

func collectionsDir() string {
	return path.Join(configDir(), "collections")
}

func homeDir() string {
	h, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return h
}

func EnsureConfigDir() error {
	err := os.MkdirAll(configDir(), 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll(collectionsDir(), 0755)
	if err != nil {
		return err
	}
	return nil
}

func (c *Collection) CreateCollection() error {
	fileName := path.Join(collectionsDir(), c.Name)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		if os.IsExist(err) {
			return errors.New(fmt.Sprintf("Collection %s already exists", c.Name))
		}
		return errors.New(fmt.Sprintf("Unknown error creating collection: %s", err))
	}
	bytes, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(bytes)
	return err
}
