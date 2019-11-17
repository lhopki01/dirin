package main

import (
	"log"

	"github.com/lhopki01/dirin/cmd"
	"github.com/lhopki01/dirin/internal/config"
)

func main() {
	err := config.EnsureConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	cmd.AddCommands()
}
