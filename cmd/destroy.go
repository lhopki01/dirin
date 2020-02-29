package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lhopki01/dirin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func registerDestroyCmd(rootCmd *cobra.Command) {
	destroyCmd := &cobra.Command{
		Use:   "destroy <space separated list of collections>",
		Short: "Delete a collection",
		Args: func(cmd *cobra.Command, args []string) error {
			return validateDestroyCmdArgs(args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			runDestroyCmd(args)
		},
	}
	rootCmd.AddCommand(destroyCmd)
	err := viper.BindPFlags(destroyCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func runDestroyCmd(args []string) {
	if len(args) == 1 {
		fmt.Printf("Deleting collection %s\n", args[0])
	}
	fmt.Printf("Deleting collections %s\n", strings.Join(args, " "))

	collectionsDir := config.CollectionsDir()
	for _, collection := range args {
		err := os.Remove(filepath.Join(collectionsDir, collection))
		if err != nil {
			println(err.Error())
		}
	}
}

func validateDestroyCmdArgs(args []string) error {
	if len(args) < 1 {
		return errors.New("Please specify at least one collection to destroy")
	}
	collections := getCollections()
	for _, c := range args {
		if stringInList(c, collections) {
			continue
		}
		return errors.New(fmt.Sprintf("Collection %s does not exist.  Please choose from: %v", c, collections))
	}
	return nil
}

func stringInList(str string, list []string) bool {
	for _, v := range list {
		if str == v {
			return true
		}
	}
	return false
}
