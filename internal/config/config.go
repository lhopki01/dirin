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
	err = c.WriteCollection(f)
	return err
}

func LoadCollection(collectionName string) (c *Collection, f *os.File, err error) {
	fileName := path.Join(collectionsDir(), collectionName)
	f, err = os.OpenFile(fileName, os.O_RDWR, 0600)
	if err != nil {
		fmt.Println("can't open config")
		log.Fatal(err)
		return nil, nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	b := make([]byte, fi.Size())
	_, err = f.Read(b)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		log.Fatal(err)
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
	err = f.Close()
	return err
}

func (c *Collection) AddDirectoriesToCollection(dirs []Dir, f *os.File) error {
	for _, dir := range dirs {
		c.Directories[dir.Path] = &dir
	}
	c.WriteCollection(f)
	return nil
}
