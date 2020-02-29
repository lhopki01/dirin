package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/lhopki01/dirin/internal/color"
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
	addCmd.Flags().String("collection", "", "The collection to add directories too")
	viper.BindPFlag("collectionAdd", addCmd.Flags().Lookup("collection"))
}

func runAddCmd(args []string) {
	collection, err := config.GetCollection("collectionAdd")
	if err != nil {
		log.Fatal(err)
	}

	c, f, _ := config.LoadCollection(collection)
	usedColors := c.GetUsedColors()

	dirs := []*config.Dir{}
	for _, dir := range args {
		if stat, err := os.Stat(dir); err == nil && stat.IsDir() {
			absoluteFilePath, err := filepath.Abs(dir)
			if err != nil {
				fmt.Printf("Can't find absolute filepath for %s\n", dir)
			} else {
				newColor := 15
				usedColors, newColor = color.NewColor(usedColors)
				newDir := &config.Dir{
					Path:  absoluteFilePath,
					Color: newColor,
					Name:  filepath.Base(dir),
				}
				dirs = append(dirs, newDir)
			}
		} else {
			fmt.Printf("%s is not a dir\n", dir)
		}
	}
	c.AddDirectoriesToCollection(dirs, f)
}
