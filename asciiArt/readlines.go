package asciiArt

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(fileName string, startLine int, linesCount int) []string {
	if startLine < 0 {
		return nil
	}
	result := []string{}
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Define the start and end line numbers
	endLineNumber := startLine + linesCount

	// Read the file line by line until you reach the end line
	currentLineNumber := 0
	for scanner.Scan() {
		currentLineNumber++

		// Check if the current line is within the specified range
		if currentLineNumber >= startLine && currentLineNumber < endLineNumber {
			// fmt.Printf("Line %d: %s\n", currentLineNumber, scanner.Text())
			result = append(result, scanner.Text())
		}

		// Break the loop if we have reached the end line
		if currentLineNumber == endLineNumber {
			break
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return result
}
