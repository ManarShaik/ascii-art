package main

import (
	"asciiArt/Split"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 || len(args[0]) == 0 {
		return
	}

	standardFile, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fileLines := strings.Split(string(standardFile), "\n")
	var start int = 0
	str := ""
	for i := 0; i < len(args[0]); i++ {
		str += string((args[0])[i])
	}
	resultStr := strings.ReplaceAll(str, "\\n", "\n")
	argLettersDivided := Split.SplitWithNewlines(resultStr)
	for i := 0; i < len(argLettersDivided); i++ {
		if argLettersDivided[i] == "\\n" {
			fmt.Println()
			continue
		}
		for j := 0; j < 8; j++ {
			for k := 0; k < len(argLettersDivided[i]); k++ {
				start = ((int((argLettersDivided[i])[k]) - 32) * 9) + (j + 1)
				if (argLettersDivided[i])[k] == ' ' {
					fmt.Print((fileLines[start]))
					continue
				}
				fmt.Print(fileLines[start])
				start = 0
			}
			fmt.Print("\n")
		}
	}

}