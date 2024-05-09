package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/logrusorgru/aurora"
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
		word := strings.ToLower(scanner.Text())
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
	file, err := os.Open("sample-correct-text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Instantiate colorizer
	colorizer := aurora.NewAurora(true)

	// Loop through each line in the file
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		words := strings.Fields(line) // Split the line into words

		for i, word := range words {
			words[i] = strings.ToLower(word)
		}

		// Track misspellings
		misspellingsFound := false

		// Check each word for misspellings
		for _, word := range words {
			if !dictionary[word] {
				fmt.Printf("%s word '%s' found at line %d\n", colorizer.BgBrightRed("Misspelling:"), colorizer.Red(word), lineNumber)
				misspellingsFound = true
			}
		}

		// If no misspellings were found in the current line, print a message
		if !misspellingsFound {
			fmt.Printf("%s No misspellings detected in line %d\n", colorizer.BgBrightGreen("Clear:"), lineNumber)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
