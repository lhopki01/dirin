package color

import (
	"fmt"
	"math/rand"

	"github.com/lhopki01/dirin/internal/config"
)

func PrintDirectory(dir *config.Dir) {
	fmt.Printf("%s:\n", ColorDirectory(dir))
}

func ColorDirectory(dir *config.Dir) string {
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", dir.Color, dir.Path)
}

func NewColor(usedColors map[int]bool) (map[int]bool, int) {
	i := 9
	for i < 16 {
		if !usedColors[i] {
			usedColors[i] = true
			fmt.Printf("Chose1 %d\n", i)
			return usedColors, i
		}
		i++
	}

	i = 1
	for i < 7 {
		if !usedColors[i] {
			usedColors[i] = true
			fmt.Printf("Chose2 %d\n", i)
			return usedColors, i
		}
		i++
	}

	// 16-18 are too dark
	for _, i := range []int{19, 20, 21, 22, 23, 24, 25, 26, 27, 28} {
		// Avoid greyscale above 231
		for i < 231 {
			if !usedColors[i] {
				usedColors[i] = true
				fmt.Printf("Chose3 %d\n", i)
				return usedColors, i
			}
			i += 10
		}

	}

	i = 0
	for i == 0 || i == 7 || i == 16 || i == 17 || i == 18 {
		i = rand.Intn(231)
	}

	fmt.Printf("Chose %d\n", i)
	return usedColors, i
}
