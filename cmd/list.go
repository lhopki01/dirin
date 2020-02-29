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
		Use:   "list",
		Short: "List all collections",
		Args:  cobra.NoArgs,
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
	for _, c := range getCollections() {
		fmt.Println(c)
	}
}

func getCollections() []string {
	fileNames := []string{}
	files, err := ioutil.ReadDir(config.CollectionsDir())
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if !f.IsDir() {
			c, err := config.LoadCollectionRead(f.Name())
			if err == nil && c != nil {
				fileNames = append(fileNames, f.Name())
			}
		}
	}
	return fileNames
}
