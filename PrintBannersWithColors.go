package asciiArtColor

import (
	"fmt"
	// "golang.org/x/image/colornames"
	// "image/color"
)

// Print the full outcome in the triminal
func PrintBannersWithColors(colors string, banners, arr []string) {
	// colors = color.colors
	num := 0
	color := WordColors(colors)
	for _, ch := range banners {
		num = num + 1
		if ch == "" {
			if num < len(banners) {
				fmt.Println()
				continue
			} else {
				continue
			}
		}
		for i := 0; i < 8; i++ {
			for _, j := range ch {
				n := (j-32)*9 + 1
				fmt.Println(color[1], arr[int(n)+i])

			}
			fmt.Println()

		}
	}
}

//     colorReset := "\033[0m"

// colorRed := "\033[31m"
// colorGreen := "\033[32m"
// colorYellow := "\033[33m"
// colorBlue := "\033[34m"
// colorPurple := "\033[35m"
// colorCyan := "\033[36m"
// colorWhite := "\033[37m"

// fmt.Println(string(colorRed), "test")
// fmt.Println(string(colorGreen), "test")
// fmt.Println(string(colorYellow), "test")
// fmt.Println(string(colorBlue), "test")
// fmt.Println(string(colorPurple), "test")
// fmt.Println(string(colorWhite), "test")
// fmt.Println(string(colorCyan), "test", string(colorReset))
// fmt.Println("next")
// }
