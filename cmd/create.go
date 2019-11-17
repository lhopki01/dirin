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
		Use:   "create [collection name]",
		Short: "Create a collection of folders for subsequent commands to run against",
		Run: func(cmd *cobra.Command, args []string) {
			runCreateCmd(args)
		},
	}
	rootCmd.AddCommand(createCmd)
	err := viper.BindPFlags(createCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func runCreateCmd(args []string) {
	fmt.Printf("Creating collection %s\n", args[0])
	c := &config.Collection{
		Name: args[0],
	}
	bytes, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
	err = c.CreateCollection()
	if err != nil {
		log.Fatal(err)
	}
}
