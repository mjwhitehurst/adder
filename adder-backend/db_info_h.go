package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

/**
 * Takes in a file (presumed to exist)
 * and returns 3 arrays - mem rec fields
 *                        rec fields
 *											  nondb fields
 */
func matfloDatabaseInfo(
	filePath string) {

}

type Field struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Comment string `json:"comment,omitempty"`
}

func findMemFields(filePath string) ([]Field, error) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Ensure the file is closed even if an error occurs

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// inStruct is a flag indicating whether the current line is within a struct definition
	inStruct := false

	// ignoreStrings are lines we'll skip over when encountered
	ignoreStrings := []string{"ifdef", "endif", "PLUGINSTART", "PLUGINEND"}

	// Fields will hold the struct fields we find
	var fields []Field

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()                 // Get the current line
		trimmedLine := strings.TrimSpace(line) // Remove leading/trailing whitespace

		// Check whether the line contains any of the strings we want to ignore
		shouldIgnore := false
		for _, ignore := range ignoreStrings {
			if strings.Contains(trimmedLine, ignore) {
				shouldIgnore = true
				break
			}
		}

		// If the line should be ignored, skip to the next line
		if shouldIgnore {
			continue
		}

		// Remove any /*pi*/ comments from the line
		if strings.Contains(trimmedLine, "/*pi*/") {
			line = strings.ReplaceAll(line, "/*pi*/", "")
		}

		// If the line starts a struct definition, set the inStruct flag to true
		if strings.HasPrefix(trimmedLine, "typedef struct") {
			inStruct = true
			continue
			// If the line ends the struct definition, break from the loop
		} else if strings.Contains(trimmedLine, "}") && strings.Contains(trimmedLine, "_MEM_REC_TYPE;") {
			break
		}

		// If the current line is within a struct definition, try to parse it as a field definition
		if inStruct {
			// This regular expression matches lines of the form "Type Name; // Comment" or "Type Name; /* Comment */"
			fieldPattern := regexp.MustCompile(`^\s*([\w\[\]]+)\s+([\w\[\]]+);\s*(?://\s*(.*)|\s*/\*\s*(.*)\s*\*/)?$`)
			matches := fieldPattern.FindStringSubmatch(line)

			// If the line matches the regular expression, extract the field's type, name, and comment
			if matches != nil {
				field := Field{
					Type:    matches[1],
					Name:    matches[2],
					Comment: matches[3],
				}

				if matches[4] != "" {
					field.Comment = matches[4]
				}

				// Add the parsed field to the list of fields
				fields = append(fields, field)
			}
		}
	}

	// If an error occurred while reading the file, return the error
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	// Return the list of fields
	return fields, nil
}

func findRecFields(filePath string) ([]Field, error) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Ensure the file is closed even if an error occurs

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Booleans to check which structure we're in.
	inMemStruct := false
	inRecStruct := false
	seenMemStruct := false

	// ignoreStrings are lines we'll skip over when encountered
	ignoreStrings := []string{"ifdef", "endif", "PLUGINSTART", "PLUGINEND", "DEFNONDBFIELD"}

	// Fields will hold the struct fields we find
	var fields []Field

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()                 // Get the current line
		trimmedLine := strings.TrimSpace(line) // Remove leading/trailing whitespace

		// Check whether the line contains any of the strings we want to ignore
		shouldIgnore := false
		for _, ignore := range ignoreStrings {
			if strings.Contains(trimmedLine, ignore) {
				shouldIgnore = true
				break
			}
		}

		// If the line should be ignored, skip to the next line
		if shouldIgnore {
			continue
		}

		// Remove any /*pi*/ comments from the line
		if strings.Contains(trimmedLine, "/*pi*/") {
			line = strings.ReplaceAll(line, "/*pi*/", "")
		}

		// First time we see 'typedef struct', it'll be the mem.
		if !inMemStruct && strings.HasPrefix(trimmedLine, "typedef struct") {
			inMemStruct = true
			seenMemStruct = true
			continue
			// If the line is the end of the mem struct, flag that we're out of it.
		} else if strings.Contains(trimmedLine, "}") && strings.Contains(trimmedLine, "_MEM_REC_TYPE;") {
			inMemStruct = false
			continue
		} else if seenMemStruct && strings.HasPrefix(trimmedLine, "typedef struct") {
			inRecStruct = true
			continue
			//end of processing hopefully.
		} else if strings.HasPrefix(trimmedLine, "typedef struct") {
			break
		}

		// If the current line is within a struct definition, try to parse it as a field definition
		if inRecStruct {
			// This regular expression matches lines of the form "Type Name; // Comment" or "Type Name; /* Comment */"
			fieldPattern := regexp.MustCompile(`^\s*([\w\[\]]+)\s+([\w\[\]]+);\s*(?://\s*(.*)|\s*/\*\s*(.*)\s*\*/)?$`)
			matches := fieldPattern.FindStringSubmatch(line)

			// If the line matches the regular expression, extract the field's type, name, and comment
			if matches != nil {
				field := Field{
					Type:    matches[1],
					Name:    matches[2],
					Comment: matches[3],
				}

				if matches[4] != "" {
					field.Comment = matches[4]
				}

				// Add the parsed field to the list of fields
				fields = append(fields, field)
			}
		}
	}

	// If an error occurred while reading the file, return the error
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	// Return the list of fields
	return fields, nil
}
