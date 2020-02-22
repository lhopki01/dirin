package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/lhopki01/dirin/internal/config"
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

func registerSwitchCmd(rootCmd *cobra.Command) {
	switchCmd := &cobra.Command{
		Use:   "switch [collection name]",
		Short: "Switch to an existing collection for subsequent commands to run against",
		Run: func(cmd *cobra.Command, args []string) {
			runSwitchCmd(args)
		},
	}
	rootCmd.AddCommand(switchCmd)
	err := viper.BindPFlags(switchCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func runSwitchCmd(args []string) {
	fmt.Println("Not currently implemented.  Please use --collection for now")
	fmt.Printf("Activating %s\n", args[0])
	config.LoadCollection(args[0])
}
