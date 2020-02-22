package cmd

import (
	"fmt"
	"sort"

	"github.com/lhopki01/dirin/internal/color"
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
	lsCmd.Flags().String("collection", "", "The collection to add directories too")
	viper.BindPFlag("collectionLs", lsCmd.Flags().Lookup("collection"))
	viper.AutomaticEnv()
}

func runLsCmd(args []string) {
	c, _ := config.LoadCollectionRead(viper.GetString("collectionLs"))
	dirs := []string{}
	for _, dir := range c.Directories {
		dirs = append(dirs, dir.Path)
	}
	sort.Strings(dirs)
	for _, dir := range dirs {
		fmt.Println(color.ColorDirectory(c.Directories[dir]))
	}
}
