package cmd

import (
	"fmt"
	"log"
	"strings"

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
	fmt.Printf("Removing directories %s\n", strings.Join(args, " "))
}
