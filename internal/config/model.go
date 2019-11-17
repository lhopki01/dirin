package config

type Dir struct {
	Name string `yaml:name`
	Path string `yaml:path`
}

type Collection struct {
	Name        string `yaml:name`
	Directories []Dir  `yaml:directories`
}
