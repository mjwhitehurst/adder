package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

/*#######################################################*/
/* Definitions                                           */
/*#######################################################*/

type DbFieldAddition struct {
	filePath  string
	fieldName string
	fieldType string
	comment   string
}

/*#######################################################*/
/* Functions                                             */
/*#######################################################*/

/**
 *	adds a field to a database.
 *		db file should already have been found
 *		looks for tag
 *			depending on nonDb it wil place
 *			(nondb)
 *    DEF NONDBFIELD FIELDTYPE FIELDNAME //COMMENT
 *		TAG
 *
 *			(!nondb)
 *
 *		FIELDTYPE FIELDNAME //COMMENT
 *		TAG
 *
 */
func addDatabaseField(
	filePath string,
	tag string,
	fieldName string,
	fieldType string,
	commentStr string,
	nonDb bool) {

	//At this point we have a file, and we know what we want to put into it.
	err := addFieldBeforeTag(filePath, nonDb, tag, fieldName, fieldType, commentStr)
	if err != nil {
		fmt.Println("Error adding field: ", err)
	}

} /* addDatabaseField */

func addRecField(filePath string, fieldName string, fieldType string, commentStr string) {
	addDatabaseField(filePath, "ADDER REC END", fieldName, fieldType, commentStr, false)
}

func addMemField(filePath string, fieldName string, fieldType string, commentStr string) {
	addDatabaseField(filePath, "ADDER MEM END", fieldName, fieldType, commentStr, false)
	return
}

func addNonDbField(filePath string, fieldName string, fieldType string, commentStr string) {
	//TODO: like addDatabaseField, but add the whole thing as a comment
	addDatabaseField(filePath, "ADDER NONDB END", fieldName, fieldType, commentStr, true)
	return
}

/**
 *  Generates a filename from a given string, hoping we have enough information.
 *          DB_NAME/db_name/db_name_definitions/db_name_definitions.h
 *
 *  looks for (and returns):
 *          db_name_definitions.h
 *
 *  on fail, returns ""
 *
 */
func definitionsFileNameFromStr(searchStr string) string {

	//Null Check
	if searchStr == "" {
		return ""
	}

	dbFileName := ""

	// Best case scenario, we are given db_name_definitions.h
	if strings.HasSuffix(searchStr, "_definitions.h") {
		//hooray
		dbFileName = strings.TrimSuffix(searchStr, "_definitions.h")

	} else if strings.HasSuffix(searchStr, "_definitions") {
		//hooray
		dbFileName = strings.TrimSuffix(searchStr, "_definitions")

	} else {
		//guess! lower case it.
		dbFileName = strings.ToLower(searchStr)
	}

	definitionsFile := dbFileName + "_definitions.h"

	//Give up
	return definitionsFile
}

/**
 * Tries to find a matching xxx_definitions.h from a string, in a directory given also
 * by a string. returns filePath if ok, nil/"" if not
 */
func findDbDefinitionsFileInDir(dbName string, dirName string) string {

	// Check nulls
	if dbName == "" || dirName == "" {
		return ""
	}

	definitionsFile := definitionsFileNameFromStr(dbName)

	filePath := filepath.Join(dirName, definitionsFile)

	if _, err := os.Stat(filePath); err == nil {
		return filePath
	} else {
		return ""
	}
} /* findDbDefinitionsFileInDir */

/**
 *	opens a file, looks for a tag, and places a string above it.
 *
 */
func addFieldBeforeTag(
	filePath string,
	nonDb bool,
	tag string,
	fieldName string,
	fieldType string,
	commentStr string) error {

	// Create a temporary file
	tempFilePath := filePath + ".addertmp"

	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	// Open the original file in read mode
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(tempFile)

	// Search for the tag and insert the new line before it
	foundTag := false
	for scanner.Scan() {
		line := scanner.Text()

		if !foundTag && strings.Contains(line, tag) {
			foundTag = true

			// Insert the new line with field details before the tag
			newLine := ""
			if nonDb {
				newLine = fmt.Sprintf("  /* DEFNONDBFLD %s %s; */ // %s", fieldType, fieldName, commentStr)
			} else {
				newLine = fmt.Sprintf("  %s %s // %s", fieldType, fieldName, commentStr)
			}

			// Write the new line with field details
			writer.WriteString(newLine)
			writer.WriteString("\n")
		}

		// Write the current line
		writer.WriteString(line)
		writer.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if !foundTag {
		return errors.New("unable to find tag in file")
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	// Close the original file
	if err := file.Close(); err != nil {
		return err
	}

	// Close the temporary file
	if err := tempFile.Close(); err != nil {
		return err
	}

	// Replace the original file with the modified temporary file
	if err := os.Rename(tempFilePath, filePath); err != nil {
		return err
	}

	// Get the original owner
	uid, gid := os.Getuid(), os.Getgid()

	// Restore the original owner to the new file
	if err := os.Chown(filePath, uid, gid); err != nil {
		return err
	}

	// Set the file permissions to -rw-rw-r--
	if err := os.Chmod(filePath, os.FileMode(0664)); err != nil {
		return err
	}

	return nil
}

func checkStringsInFile(filePath string, str1 string, str2 string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	foundStr1 := false
	foundStr2 := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, str1) {
			foundStr1 = true
		}
		if strings.Contains(line, str2) {
			foundStr2 = true
		}

		// Exit the loop early if both strings are found
		if foundStr1 && foundStr2 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Check if both strings are found in the file
	if foundStr1 && foundStr2 {
		return nil // Both strings are present
	}

	return errors.New("strings not found in file")
}

/**
 *	Returns a list of all the definitions files a system has
 */
func findDefinitionFiles() (string, int) {
	srcDir := "/app/sourcedir"
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	var definitionFiles []string
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, "_definitions.h") {
			definitionName := strings.TrimSuffix(filename, "_definitions.h")
			definitionFiles = append(definitionFiles, definitionName)
		}
	}

	if len(definitionFiles) == 0 {
		return "No definition files found", http.StatusNotFound
	}

	jsonOutput, err := json.Marshal(definitionFiles)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	return string(jsonOutput), http.StatusOK
}
