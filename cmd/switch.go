package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/lhopki01/dirin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
	_, err := config.LoadCollectionRead(args[0])
	if err != nil {
		fmt.Printf("Collection %s does not exist. Please choose from:\n", args[0])
		runListCmd([]string{})
		fmt.Println("Or create one using dirin create <collection name>")
		os.Exit(1)
	}
	fmt.Printf("Activating %s\n", args[0])
	config, f, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	config.ActiveCollection = args[0]
	err = config.WriteConfig(f)
	if err != nil {
		fmt.Printf("Failed to switch to collection %s with err: %v", args[0], err)
	}
}
