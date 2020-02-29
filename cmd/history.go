package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/lhopki01/dirin/internal/color"
	"github.com/lhopki01/dirin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func registerHistoryCmd(rootCmd *cobra.Command) {
	historyCmd := &cobra.Command{
		Use:   "history [options]",
		Short: "Show the results of previous commands run in a collection",
		Run: func(cmd *cobra.Command, args []string) {
			historyHistoryCmd(args)
		},
	}
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags().String("collection", "", "The collection to add directories too")
	viper.BindPFlag("collectionHistory", historyCmd.Flags().Lookup("collection"))
	viper.AutomaticEnv()
}

func historyHistoryCmd(args []string) {
	collection, err := config.GetCollection("collectionHistory")
	if err != nil {
		log.Fatal(err)
	}
	c, _ := config.LoadCollectionRead(collection)
	for _, dir := range c.Directories {
		color.PrintDirectory(dir)
		for _, cmd := range dir.Commands {
			fmt.Printf("%s\n", strings.Join(cmd.Command, " "))
			fmt.Printf(cmd.Output)
		}
	}
}
