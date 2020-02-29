package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lhopki01/dirin/internal/color"
	"github.com/lhopki01/dirin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func registerHistoryCmd(rootCmd *cobra.Command) {
	historyCmd := &cobra.Command{
		Use:   "history",
		Short: "Show the results of previous commands run in a collection",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			historyHistoryCmd()
		},
	}
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags().String("collection", "", "The collection to add directories too")
	viper.BindPFlag("collectionHistory", historyCmd.Flags().Lookup("collection"))
	viper.AutomaticEnv()
}

func historyHistoryCmd() {
	collection, err := config.GetCollection("collectionHistory")
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

	for _, dir := range c.Directories {
		color.PrintDirectory(dir)
		for _, cmd := range dir.Commands {
			fmt.Printf("%s\n", strings.Join(cmd.Command, " "))
			fmt.Printf(cmd.Output)
		}
	}
}
