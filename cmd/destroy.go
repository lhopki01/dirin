package cmd

import (
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
		Use:   "destroy [collection name]",
		Short: "Delete a collection",
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
	} else if len(args) > 1 {
		fmt.Printf("Deleting collections %s\n", strings.Join(args, " "))
	} else {
		log.Fatal("Please specify a list of collections to delete")
	}
	collectionsDir := config.CollectionsDir()
	for _, collection := range args {
		err := os.Remove(filepath.Join(collectionsDir, collection))
		if err != nil {
			println(err.Error())
		}
	}
}
