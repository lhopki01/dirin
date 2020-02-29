package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func AddCommands() {
	rootCmd := &cobra.Command{
		Use:   "dirin <subcommand>",
		Short: "A tool to run commands across multiple directories",
	}

	// Collection
	registerCreateCmd(rootCmd)
	registerSwitchCmd(rootCmd)
	registerDestroyCmd(rootCmd)
	registerListCmd(rootCmd)

	// Commands
	registerAddCmd(rootCmd)
	registerRunCmd(rootCmd)
	registerHistoryCmd(rootCmd)
	registerRmCmd(rootCmd)
	registerLsCmd(rootCmd)

	viper.AutomaticEnv()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
