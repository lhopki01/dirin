package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/lhopki01/dirin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func registerRmCmd(rootCmd *cobra.Command) {
	rmCmd := &cobra.Command{
		Use:   "rm [options] <list of directories>",
		Short: "Remove directories from a collection",
		Run: func(cmd *cobra.Command, args []string) {
			rmRmCmd(args)
		},
	}
	rootCmd.AddCommand(rmCmd)
	err := viper.BindPFlags(rmCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func rmRmCmd(args []string) {
	if len(args) == 1 {
		fmt.Printf("Removing directory %s\n", strings.Join(args, " "))
	} else if len(args) > 1 {
		fmt.Printf("Removing directories %s\n", strings.Join(args, " "))
	} else {
		log.Fatal("Please specify a list of directories to remove")
	}
	c, f, _ := config.LoadCollection(viper.GetString("collection"))
	for _, dir := range args {
		absoluteFilePath, err := filepath.Abs(dir)
		if err != nil {
			fmt.Printf("Can't find absolute path for %s\n", dir)
		}
		delete(c.Directories, absoluteFilePath)
	}
	c.WriteCollection((f))
}
