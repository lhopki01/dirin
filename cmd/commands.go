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

	// Project
	registerCreateCmd(rootCmd)
	registerSwitchCmd(rootCmd)
	registerDestroyCmd(rootCmd)
	registerListCmd(rootCmd)
	registerAddCmd(rootCmd)

	// Commands
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
	fmt.Printf("Activating %s\n", args[0])
	config.LoadCollection(args[0])
}

func registerDestroyCmd(rootCmd *cobra.Command) {
	destroyCmd := &cobra.Command{
		Use:   "destroy [collection name]",
		Short: "Delete a collection of folders",
		Run: func(cmd *cobra.Command, args []string) {
			runDestroyCmd(args)
		},
	}
	rootCmd.AddCommand(destroyCmd)
	err := viper.BindPFlags(destroyCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func runDestroyCmd(args []string) {
	fmt.Printf("Deleting collection %s\n", args[0])
}

func registerListCmd(rootCmd *cobra.Command) {
	listCmd := &cobra.Command{
		Use:   "list [project name]",
		Short: "List all collections",
		Run: func(cmd *cobra.Command, args []string) {
			runListCmd(args)
		},
	}
	rootCmd.AddCommand(listCmd)
	err := viper.BindPFlags(listCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func runListCmd(args []string) {
	fmt.Printf("Listing collections\n")
}

func registerHistoryCmd(rootCmd *cobra.Command) {
	historyCmd := &cobra.Command{
		Use:   "history [options]",
		Short: "Show the results of previous commands run in this collection",
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
}

func registerLsCmd(rootCmd *cobra.Command) {
	lsCmd := &cobra.Command{
		Use:   "ls [options]",
		Short: "List all directories in collection",
		Run: func(cmd *cobra.Command, args []string) {
			lsLsCmd(args)
		},
	}
	rootCmd.AddCommand(lsCmd)
	err := viper.BindPFlags(lsCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func lsLsCmd(args []string) {
	fmt.Printf("directories %s\n", args[0])
}
