package cmd

import (
	"fmt"
	"log"
	"strings"

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
	err := viper.BindPFlags(historyCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func historyHistoryCmd(args []string) {
	fmt.Printf("History for current project\n")
	c, _ := config.LoadCollectionRead(viper.GetString("collection"))
	for _, dir := range c.Directories {
		fmt.Printf("%s:\n", dir.Path)
		for _, cmd := range dir.Commands {
			fmt.Printf("%s\n", strings.Join(cmd.Command, " "))
			fmt.Printf(cmd.Output)
		}
	}

}
