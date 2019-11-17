package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

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

func registerAddCmd(rootCmd *cobra.Command) {
	addCmd := &cobra.Command{
		Use:   "add [list of directories]",
		Short: "Add a list directories to a collection",
		Run: func(cmd *cobra.Command, args []string) {
			runAddCmd(args)
		},
	}
	rootCmd.AddCommand(addCmd)
	err := viper.BindPFlags(addCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func runAddCmd(args []string) {
	fmt.Printf("Adding %s\n to current collection", strings.Join(args, " "))
}

func registerRunCmd(rootCmd *cobra.Command) {
	runCmd := &cobra.Command{
		Use:   "run [options] cmd",
		Short: "Execute a command on all directories in project",
		Run: func(cmd *cobra.Command, args []string) {
			runRunCmd(args)
		},
	}
	rootCmd.AddCommand(runCmd)
	err := viper.BindPFlags(runCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func runRunCmd(args []string) {
	fmt.Printf("Running %s\n", args[0])
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
