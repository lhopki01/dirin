package config

type Config struct {
	ActiveProject string `yaml:"activeProject"`
}

type Collection struct {
	Name        string          `yaml:"name"`
	Directories map[string]*Dir `yaml:"directories"`
}

type Dir struct {
	Name     string    `yaml:"name"`
	Path     string    `yaml:"path"`
	Commands []Command `yaml:"commands"`
}

type Command struct {
	Command  []string `yaml:"command"`
	ExitCode int      `yaml:"exitCode"`
	Output   string   `yaml:"output"`
}
