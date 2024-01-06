package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run verify.go <your_wordlist.txt>")
		return
	}

	filename := os.Args[1]
	folder := "fuzzylists/"

	// Read the content of the given file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}
	searchString := string(content)

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		fmt.Println("Error reading the directory:", err)
		return
	}

	for _, file := range files {
		filePath := folder + file.Name()
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNumber := 1
		found := false
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), searchString) {
				fmt.Printf("Match found in file: %s at line: %d\n", filePath, lineNumber)
				found = true
				break
			}
			lineNumber++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error scanning the file:", err)
		}

		if found {
			break
		}
	}
}
