package cmd

import (
	"fmt"
	"log"

	"github.com/lhopki01/dirin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v3"
)

func registerCreateCmd(rootCmd *cobra.Command) {
	createCmd := &cobra.Command{
		Use:   "create <collection name>",
		Short: "Create a collection of directories for subsequent commands to run against",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runCreateCmd(args[0])
		},
	}
	rootCmd.AddCommand(createCmd)
	err := viper.BindPFlags(createCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func runCreateCmd(collection string) {
	fmt.Printf("Creating collection %s\n", collection)
	c := &config.Collection{
		Name: collection,
	}
	_, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	err = c.CreateCollection()
	if err != nil {
		log.Fatal(err)
	}
	runSwitchCmd(collection)
}
