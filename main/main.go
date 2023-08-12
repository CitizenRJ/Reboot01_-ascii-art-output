package main

import (
	"asciiArtOutPut"
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
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println(len(os.Args), "is Not a valid amount of arguments. \n" )
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}
	args := os.Args[1:]
	// bol := asciiArtFs.IsValid(args[0])
	if !(asciiArtOutPut.IsValid(args[0])) {
		fmt.Println("This's Not a valid character.")
		return
	}

	text := args[1] // "hello" == [0]
	output := args[0]
	font := "standard" //base font
	if len(args) == 3 {
		switch args[2] {
		case "shadow":
			font = "shadow"
		case "thinkertoy":
			font = "thinkertoy"
		case "standard":
			font = "standard"
		default:
			fmt.Println(args[2] ,"is Not a valid font.")
			return
		}
	}

	outputFile := strings.Split(output, "--output=")
	outputFileName := ""
	for _, name := range outputFile {
		outputFileName = outputFileName + name
	}
	NameLen := len(outputFileName) 
	if NameLen < 5 {
		fmt.Println(outputFileName ,"is Not a valid output File Name.")
		return
	} else if !(outputFileName[NameLen-1] == 't' && outputFileName[NameLen-2] == 'x' && outputFileName[NameLen-3] == 't' && outputFileName[NameLen-4] == '.') {
	// 	for i := len(outputFileName)-1; i >= 0; i-- {
		fmt.Println("output File Name should end with <.txt> .")
		return
	// }
	}
	
	// Read the content of the file
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
	asciiArtOutPut.PrintBanners(outputFileName, argsArr, arr)
}
