package cmd

import (
	"fmt"
	"log"
	"sort"

	"github.com/lhopki01/dirin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func registerLsCmd(rootCmd *cobra.Command) {
	lsCmd := &cobra.Command{
		Use:   "ls [options]",
		Short: "List all directories in a collection",
		Run: func(cmd *cobra.Command, args []string) {
			runLsCmd(args)
		},
	}
	rootCmd.AddCommand(lsCmd)
	err := viper.BindPFlags(lsCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func runLsCmd(args []string) {
	c, _ := config.LoadCollectionRead(viper.GetString("collection"))
	dirs := []string{}
	for _, dir := range c.Directories {
		dirs = append(dirs, dir.Path)
	}
	sort.Strings(dirs)
	for _, dir := range dirs {
		fmt.Println(dir)
	}
}
