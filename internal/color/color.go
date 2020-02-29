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
	if dir.Color == 0 {
		dir.Color = 15
	}
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", dir.Color, dir.Path)
}

func NewColor(usedColors map[int]bool) (map[int]bool, int) {
	// Nicest colors: 9-15
	i := 9
	for i < 16 {
		if !usedColors[i] {
			usedColors[i] = true
			return usedColors, i
		}
		i++
	}

	// Second nicest colors: 1-6
	i = 1
	for i < 7 {
		if !usedColors[i] {
			usedColors[i] = true
			return usedColors, i
		}
		i++
	}

	// 16-18 are too dark
	// Increment by 10 to avoid too similar colors being next to each other
	for _, i := range []int{19, 20, 21, 22, 23, 24, 25, 26, 27, 28} {
		// Avoid greyscale above 231
		for i < 231 {
			if !usedColors[i] {
				usedColors[i] = true
				return usedColors, i
			}
			i += 10
		}

	}

	// When we've run out of free colors just start randomly assigning them
	i = 0
	for i == 0 || i == 7 || i == 16 || i == 17 || i == 18 {
		i = rand.Intn(231)
	}

	return usedColors, i
}
