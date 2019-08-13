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
		Use:   "dirin",
		Short: "A tool to run commands across multiple directories",
	}

	registerSwitchCmd(rootCmd)
	registerAddCmd(rootCmd)
	registerRunCmd(rootCmd)
	registerCreateCmd(rootCmd)
	registerHistoryCmd(rootCmd)
	registerRmCmd(rootCmd)
	registerDestroyCmd(rootCmd)
	registerListCmd(rootCmd)
	registerLsCmd(rootCmd)

	viper.AutomaticEnv()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func registerSwitchCmd(rootCmd *cobra.Command) {
	switchCmd := &cobra.Command{
		Use:   "switch [project name]",
		Short: "Switch a collection of folders for subsequent commands to run against",
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

func registerAddCmd(rootCmd *cobra.Command) {
	addCmd := &cobra.Command{
		Use:   "add [list of directories]",
		Short: "Add a list folders to the project",
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

func runSwitchCmd(args []string) {
	fmt.Printf("Activating %s\n", args[0])
}

func runAddCmd(args []string) {
	fmt.Printf("Adding directories %s\n", args[0])
}

func runRunCmd(args []string) {
	fmt.Printf("Running %s\n", strings.Join(args, " "))
}
