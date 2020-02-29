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
		Use:   "switch <collection name>",
		Short: "Switch to an existing collection for subsequent commands to run against",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runSwitchCmd(args[0])
		},
	}
	rootCmd.AddCommand(switchCmd)
	err := viper.BindPFlags(switchCmd.Flags())
	if err != nil {
		log.Fatalf("Binding flags failed: %s", err)
	}
	viper.AutomaticEnv()
}

func runSwitchCmd(collection string) {
	_, err := config.LoadCollectionRead(collection)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Collection %s does not exit.  Please choose from: %v\n", collection, getCollections())
			fmt.Println("Or create the collection using dirin create <collection name>")
			os.Exit(1)
		}
		log.Fatal(err)
	}
	fmt.Printf("Activating %s\n", collection)
	config, f, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	config.ActiveCollection = collection
	err = config.WriteConfig(f)
	if err != nil {
		fmt.Printf("Failed to switch to collection %s with err: %v", collection, err)
	}
}
