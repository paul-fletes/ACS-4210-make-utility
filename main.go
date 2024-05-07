package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// LoadDictionary loads words from the system's dictionary into a map
func LoadDictionary() (map[string]bool, error) {
	dictionary := make(map[string]bool)

	cmd := exec.Command("cat", "/usr/share/dict/words") // Path to system dictionary file
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		word := scanner.Text()
		dictionary[word] = true
	}

	return dictionary, scanner.Err()
}

func main() {
	// Load dictionary
	dictionary, err := LoadDictionary()
	if err != nil {
		fmt.Println("Error loading dictionary:", err)
		return
	}

	// Read the file
	file, err := os.Open("sample-input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Loop through each line in the file
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		words := strings.Fields(line) // Split the line into words

		// Check each word for misspellings
		for _, word := range words {
			if !dictionary[word] {
				fmt.Printf("Misspelled word '%s' found at line %d\n", word, lineNumber)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
