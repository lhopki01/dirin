package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v3"
)

func configDir() string {
	appName := "dirin"
	if xdgDir := os.Getenv("XDG_CONFIG_HOME"); xdgDir != "" {
		return path.Join(xdgDir, appName)
	}
	return path.Join(homeDir(), ".config", appName)
}

func CollectionsDir() string {
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
	err = os.MkdirAll(CollectionsDir(), 0755)
	if err != nil {
		return err
	}
	err = EnsureConfigFile()
	if err != nil {
		return err
	}
	return nil
}

func EnsureConfigFile() error {
	fileName := path.Join(configDir(), "config.yaml")
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return errors.New(fmt.Sprintf("Unknown error creating collection: %s", err))
	}
	bytes, err := yaml.Marshal(Config{})
	if err != nil {
		return err
	}
	f.Truncate(0)
	f.Seek(0, 0)
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}
	err = f.Close()
	return err
}

func LoadConfigRead() (c *Config, err error) {
	c, f, err := LoadConfig()
	f.Close()
	return c, err
}

func LoadConfig() (c *Config, f *os.File, err error) {
	fileName := path.Join(configDir(), "config.yaml")
	f, err = os.OpenFile(fileName, os.O_RDWR, 0600)
	if err != nil {
		return nil, f, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, f, err
	}
	b := make([]byte, fi.Size())
	_, err = f.Read(b)
	if err != nil {
		return nil, f, err
	}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return nil, f, err
	}
	return c, f, err
}

func (c *Config) WriteConfig(f *os.File) (err error) {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	f.Truncate(0)
	f.Seek(0, 0)
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}
	err = f.Close()
	return err
}

func (c *Collection) CreateCollection() error {
	fileName := path.Join(CollectionsDir(), c.Name)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		if os.IsExist(err) {
			return errors.New(fmt.Sprintf("Collection %s already exists", c.Name))
		}
		return errors.New(fmt.Sprintf("Unknown error creating collection: %s", err))
	}
	err = c.WriteCollection(f)
	return err
}

func LoadCollectionRead(collectionName string) (c *Collection, err error) {
	c, f, err := LoadCollection(collectionName)
	f.Close()
	return c, err
}

func LoadCollection(collectionName string) (c *Collection, f *os.File, err error) {
	fileName := path.Join(CollectionsDir(), collectionName)
	f, err = os.OpenFile(fileName, os.O_RDWR, 0600)
	if err != nil {
		return nil, nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, nil, err
	}
	b := make([]byte, fi.Size())
	_, err = f.Read(b)
	if err != nil {
		return nil, nil, err
	}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return nil, nil, err
	}
	return c, f, err
}

func (c *Collection) WriteCollection(f *os.File) (err error) {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	f.Truncate(0)
	f.Seek(0, 0)
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}
	err = f.Close()
	return err
}

func (c *Collection) AddDirectoriesToCollection(dirs []*Dir, f *os.File) error {
	for _, dir := range dirs {
		if _, ok := c.Directories[dir.Path]; !ok {
			c.Directories[dir.Path] = dir
		}
	}
	c.WriteCollection(f)
	return nil
}

func (c *Collection) GetUsedColors() map[int]bool {
	usedColors := map[int]bool{}
	for _, dir := range c.Directories {
		usedColors[dir.Color] = true
	}
	return usedColors
}

func GetCollection(viperVar string) (string, error) {
	collection := viper.GetString(viperVar)
	if collection != "" {
		return collection, nil
	}
	c, err := LoadConfigRead()
	if err != nil {
		return "", err
	}
	if c.ActiveCollection == "" {
		fmt.Println("No active collection\nPlease use --collection flag or dirin switch [collection]")
		os.Exit(1)
	}
	return c.ActiveCollection, nil
}
