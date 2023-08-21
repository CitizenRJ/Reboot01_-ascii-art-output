package main

import (
	"asciiArtColor"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	fileLen = 855
)

// check amount of arguments
func main() {
	if len(os.Args) < 2 || len(os.Args) > 5 {
		fmt.Println(len(os.Args), "is Not a valid amount of arguments.")
		return
	}
	substr := "--output="
	substr2 := "--color="
	str := os.Args[1]
	for i := 0; i < len(str)-len(substr)+1; i++ {
		if str[i:i+len(substr)] == substr && len(os.Args) == 3 {
			os.Args = append(os.Args, "standard")
		} else {
			continue
		}
	}
	for i := 0; i < len(str)-len(substr2)+1; i++ {
		if str[i:i+len(substr2)] == substr2 && len(os.Args) == 3 {
			os.Args = append(os.Args, "standard")
		} else {
			continue
		}
	}

	args := os.Args[1:]
	ArgsLen := len(args)
	if !(asciiArtColor.IsValid(args[0])) {
		fmt.Println("This's Not a valid character.")
		return
	}

	font := "standard" //base font
	outputFileName := ""
	colortextToP := ""
	text := ""
	text = args[0] // "hello" == [0]

	if ArgsLen <= 2 {
		if ArgsLen == 1 {
			os.Args = append(os.Args, "standard")
			args = os.Args[1:]
			ArgsLen = len(args)
		}
		switch args[ArgsLen-1] {
		case "shadow":
			font = "shadow"
		case "thinkertoy":
			font = "thinkertoy"
		case "standard":
			font = "standard"
		default:
			fmt.Println(args[ArgsLen-1], "is Not a valid font.", args, ArgsLen)
			return
		}
	}
	if len(os.Args) >= 4 {
		text = args[0+1] // "hello" == [0]
		output := ""
		color := ""
		for i := 0; i < len(args)-2; i++ {
			word := args[i]
			/////////////////////////////////////
			if word[i:i+len(substr2)] == substr2 {
				color = word
				/////////////////////////////////////
			} else if word[i:i+len(substr)] == substr {
				output = word
			} else {
				continue
			}
		}
		//////////////////////////////////////////////////////
		outputFile := strings.Split(output, "--output=")
		for i := 0; i < len(outputFile); i++ {
			if len(outputFile) == 1 && outputFile[i] != "" {
				fmt.Println("you should print the run like this example: ")
				fmt.Println("EX: go run . --output=<fileName.txt> something standard")
				fmt.Println(outputFile, "guiothjerikhgmjopr")
				return
			} else if i+1 != len(outputFile) {
				outputFileName = outputFile[1]
				if outputFile[0] != "" {
					fmt.Println("tf wrong w u.")
					return
				}
				break
			}
			if len(outputFile) > 2 || outputFile[0] != "" {
				fmt.Println("Error: smt wrong w the output string.")
				return
			}
		}
		if len(outputFile) == 2 && outputFileName != "" {
			NameLen := len(outputFileName)
			if NameLen < 5 {
				fmt.Println(outputFileName, "is Not a valid output File Name.")
				return
			} else if !(outputFileName[NameLen-1] == 't' && outputFileName[NameLen-2] == 'x' && outputFileName[NameLen-3] == 't' && outputFileName[NameLen-4] == '.') {
				fmt.Println("output File Name should end with <.txt> .")
				return
			}
		}
		////////////////////////////////////////////////////////////////////////////
		textColor := strings.Split(color, "--color=")
		for i := 0; i < len(textColor); i++ {
			if len(textColor) == 1 && textColor[i] != "" {
				fmt.Println("you should print the run like this example: ")
				fmt.Println("EX: go run . --output=<fileName.txt> something standard")
				fmt.Println(textColor, "gr")
				return
			} else if i+1 != len(textColor) {
				colortextToP = textColor[1]
				if textColor[0] != "" {
					fmt.Println("tf wrong w u.")
					return
				}
				break
			}
			if len(textColor) > 2 || textColor[0] != "" {
				fmt.Println("Error: smt wrong w the color string.")
				return
			}
		}
	}
	// else {
	// 	text = args[0] // "hello" == [0]
	// 	if len(args) == 2 {
	// 		switch args[1] {
	// 		case "shadow":
	// 			font = "shadow"
	// 		case "thinkertoy":
	// 			font = "thinkertoy"
	// 		case "standard":
	// 			font = "standard"
	// 		default:
	// 			fmt.Println("Not a valid font")
	// 			return
	// 		}
	// 	}

	// }

	// Read the content of the file
	text = strings.ReplaceAll(text, "\\t", "   ")
	argsArr := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")
	arr := []string{}
	readFile, err := os.Open("fonts/" + font + ".txt")
	defer readFile.Close()

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}

	if len(arr) != fileLen {
		fmt.Println("File is corrupted.")
		return
	}
	larg := len(argsArr)
	if larg >= 2 {
		if argsArr[larg-1] == "" && argsArr[larg-2] != "" {
			argsArr = argsArr[:larg-1]
		}
	}
	// fmt.Println(colortextToP, argsArr)

	if outputFileName != "" {
		asciiArtColor.PrintBannersInFile(outputFileName, argsArr, arr)
	} else if colortextToP != "" {
		asciiArtColor.PrintBannersWithColors(colortextToP, argsArr, arr)
	} else {
		asciiArtColor.PrintBanners(argsArr, arr)
	}
}
