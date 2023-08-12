package asciiArtOutPut

import (
	"fmt"
	"bufio"
	"os"
)

// Print the full outcome
func PrintBanners(outputFileName string, banners, arr []string) {
	file, errs := os.Create(outputFileName)
		if errs != nil {
			fmt.Println("Failed to create file:", errs)
			os.Exit(2)
		}
		defer file.Close()
	writer := bufio.NewWriter(file)
	num := 0
	for _, ch := range banners {
		num = num + 1
		if ch == "" {
			if num < len(banners) {
				fmt.Fprintln(writer, "")
				// fmt.Println()
				continue
			} else {
				continue
			}
		}
		for i := 0; i < 8; i++ {
			for _, j := range ch {
				n := (j-32)*9 + 1
				fmt.Fprint(writer, arr[int(n)+i])
				// fmt.Print(arr[int(n)+i])

			}
			fmt.Fprintln(writer, "")
			// fmt.Println()
		}
		fmt.Fprintln(writer, "")
		// fmt.Println()
	}
	writer.Flush()
	fmt.Println("Wrote to file: " + outputFileName + ".")
}
