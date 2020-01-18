package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/lhopki01/dirin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func registerListCmd(rootCmd *cobra.Command) {
	listCmd := &cobra.Command{
		Use:   "list [project name]",
		Short: "List all collections",
		Run: func(cmd *cobra.Command, args []string) {
			runListCmd(args)
		},
	}
	rootCmd.AddCommand(listCmd)
	err := viper.BindPFlags(listCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func runListCmd(args []string) {
	files, err := ioutil.ReadDir(config.CollectionsDir())
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if !f.IsDir() {
			c, err := config.LoadCollectionRead(f.Name())
			if err == nil && c != nil {
				fmt.Println(f.Name())
			}
		}
	}
}
