package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/lhopki01/dirin/internal/color"
	"github.com/lhopki01/dirin/internal/config"
	"github.com/remeh/sizedwaitgroup"
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
	runCmd.Flags().String("collection", "", "The collection to add directories too")
	runCmd.Flags().Int("parallelism", 25, "The number of processes to run in parallel (default 25)")
	viper.BindPFlag("collectionRun", runCmd.Flags().Lookup("collection"))
	viper.BindPFlag("parallelism", runCmd.Flags().Lookup("parallelism"))
	viper.AutomaticEnv()
}

func runRunCmd(args []string) {
	collection, err := config.GetCollection("collectionRun")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Running %s\n", strings.Join(args, " "))
	c, f, _ := config.LoadCollection(collection)
	swg := sizedwaitgroup.New(viper.GetInt("parallelism"))
	for _, dir := range c.Directories {
		swg.Add()
		go runCommand(c, dir, args, &swg)
	}
	swg.Wait()
	c.WriteCollection(f)
}

func runCommand(c *config.Collection, dir *config.Dir, commands []string, swg *sizedwaitgroup.SizedWaitGroup) {
	defer swg.Done()
	cmd := exec.Command("sh", append([]string{"-c"}, commands...)...)
	cmd.Dir = dir.Path
	combinedOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	color.PrintDirectory(dir)
	fmt.Printf(string(combinedOutput))
	command := config.Command{
		Command: commands,
		Output:  string(combinedOutput),
	}
	c.Directories[dir.Path].Commands = append(c.Directories[dir.Path].Commands, command)
}
