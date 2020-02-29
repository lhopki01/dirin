package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/lhopki01/dirin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func registerRmCmd(rootCmd *cobra.Command) {
	rmCmd := &cobra.Command{
		Use:   "rm <list of directories>",
		Short: "Remove directories from a collection",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runRmCmd(args)
		},
	}
	rootCmd.AddCommand(rmCmd)
	rmCmd.Flags().String("collection", "", "The collection to add directories too")
	viper.BindPFlag("collectionRm", rmCmd.Flags().Lookup("collection"))
}

func runRmCmd(args []string) {
	collection, err := config.GetCollection("collectionRm")
	if err != nil {
		log.Fatal(err)
	}

	c, f, err := config.LoadCollection(collection)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Collection %s does not exit.  Please choose from: %v\n", collection, getCollections())
			fmt.Println("Or create the collection using dirin create <collection name>")
			os.Exit(1)
		}
		log.Fatal(err)
	}

	for _, dir := range args {
		absoluteFilePath, err := filepath.Abs(dir)
		if err != nil {
			fmt.Printf("Can't find absolute path for %s\n", dir)
		}
		fmt.Printf("Removing %s\n", dir)
		delete(c.Directories, absoluteFilePath)
	}

	c.WriteCollection((f))
}
