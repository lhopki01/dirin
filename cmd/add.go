package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/lhopki01/dirin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func registerAddCmd(rootCmd *cobra.Command) {
	addCmd := &cobra.Command{
		Use:   "add [list of directories]",
		Short: "Add a list directories to a collection",
		Run: func(cmd *cobra.Command, args []string) {
			runAddCmd(args)
		},
	}
	rootCmd.AddCommand(addCmd)
	//addCmd.Flags().String("collectionA", "", "The collection to add directories too")
	//err := viper.BindPFlags(addCmd.Flags())
	//if err != nil {
	//	log.Fatalf("Binding flags failed: %s", err)
	//}
	viper.AutomaticEnv()
}

func runAddCmd(args []string) {
	dirs := []*config.Dir{}
	for _, dir := range args {
		if stat, err := os.Stat(dir); err == nil && stat.IsDir() {
			newDir := &config.Dir{
				Path: dir,
				Name: filepath.Base(dir),
			}
			dirs = append(dirs, newDir)
		} else {
			fmt.Printf("%s is not a dir\n", dir)
		}
	}
	spew.Dump(dirs)
	c, f, _ := config.LoadCollection(viper.GetString("collection"))
	c.AddDirectoriesToCollection(dirs, f)
}
