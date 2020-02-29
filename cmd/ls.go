package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/lhopki01/dirin/internal/color"
	"github.com/lhopki01/dirin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func registerLsCmd(rootCmd *cobra.Command) {
	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "List all directories in a collection",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runLsCmd(args)
		},
	}
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().String("collection", "", "The collection to add directories too")
	viper.BindPFlag("collectionLs", lsCmd.Flags().Lookup("collection"))
	viper.AutomaticEnv()
}

func runLsCmd(args []string) {
	collection, err := config.GetCollection("collectionLs")
	if err != nil {
		log.Fatal(err)
	}

	c, err := config.LoadCollectionRead(collection)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Collection %s does not exit.  Please choose from: %v\n", collection, getCollections())
			fmt.Println("Or create the collection using dirin create <collection name>")
			os.Exit(1)
		}
		log.Fatal(err)
	}

	dirs := getDirectories(c)
	for _, dir := range dirs {
		fmt.Println(color.ColorDirectory(c.Directories[dir]))
	}
}

func getDirectories(c *config.Collection) []string {
	dirs := []string{}
	for _, dir := range c.Directories {
		dirs = append(dirs, dir.Path)
	}
	sort.Strings(dirs)
	return dirs
}
