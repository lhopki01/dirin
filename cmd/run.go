package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/lhopki01/dirin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func registerRunCmd(rootCmd *cobra.Command) {
	runCmd := &cobra.Command{
		Use:   "run [options] cmd",
		Short: "Execute a command in all directories in a collection",
		Run: func(cmd *cobra.Command, args []string) {
			runRunCmd(args)
		},
	}
	rootCmd.AddCommand(runCmd)
	//runCmd.Flags().String("collection", "", "The collection to add directories too")
	//err := viper.BindPFlags(runCmd.Flags())
	//if err != nil {
	//	log.Fatalf("Binding flags failed: %s", err)
	//}
	viper.AutomaticEnv()
}

func runRunCmd(args []string) {
	fmt.Printf("Running %s\n", strings.Join(args, " "))
	c, f, _ := config.LoadCollection(viper.GetString("collection"))
	var wg sync.WaitGroup
	for _, dir := range c.Directories {
		wg.Add(1)
		go runCommand(c, dir, args, &wg)
	}
	wg.Wait()
	c.WriteCollection(f)
}

func runCommand(c *config.Collection, dir *config.Dir, commands []string, wg *sync.WaitGroup) {
	cmd := exec.Command("sh", append([]string{"-c"}, commands...)...)
	cmd.Dir = dir.Path
	combinedOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s:\n", dir.Path)
	fmt.Printf(string(combinedOutput))
	command := config.Command{
		Command: commands,
		Output:  string(combinedOutput),
	}
	c.Directories[dir.Path].Commands = append(c.Directories[dir.Path].Commands, command)
	wg.Done()
}
