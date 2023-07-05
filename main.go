package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Define all the actions here
// adding one here requires adding  handling later
const (
	displayHelp          = -1
	actionAddRecField    = 1
	actionAddMemField    = 2
	actionAddNonDbField  = 3
	actionChainDatabases = 4
)

/*#######################################################*/
/* Main function                                         */
/*#######################################################*/
func main() {

	// Run mode
	runModeServer := false
	action := -1

	filePath := ""
	newDbFieldAddition := DbFieldAddition{}

	// Check the number of command-line arguments
	if len(os.Args) < 2 {
		runModeServer = true
	}

	/* If we're command-line run, then work out what we're doing. */
	if !runModeServer {
		// Access the command-line argument
		firstArg := os.Args[1]

		// Process the argument and display the result
		action = actionFromString(firstArg)

		if action == -1 {
			printHelp()
			return
		}

		// we are running as a command line action - look for our file from arguments.
		newDbFieldAddition.filePath = findFilePathCmdLine()
		err := validateCmdLineArgs(&newDbFieldAddition)

		if err != nil || filePath == "" {
			printHelp()
			return
		}

	} else { // we are running as a server
		//TODO: get the action with other data from a http request
		action = actionAddRecField
	}
	/* BLOCK HERE TO CHECK SERVER*/

	switch action {
	case actionAddRecField:
		addRecField(newDbFieldAddition.filePath,
			newDbFieldAddition.fieldName,
			newDbFieldAddition.fieldType,
			newDbFieldAddition.comment)
		break
	case actionAddMemField:
		addMemField(newDbFieldAddition.filePath,
			newDbFieldAddition.fieldName,
			newDbFieldAddition.fieldType,
			newDbFieldAddition.comment)
		break
	case actionAddNonDbField:
		addNonDbField(newDbFieldAddition.filePath,
			newDbFieldAddition.fieldName,
			newDbFieldAddition.fieldType,
			newDbFieldAddition.comment)
		break
	default:
		break
	}

} /* main */

/*#######################################################*/
/* Functions                                             */
/*#######################################################*/

func printHelp() {

	fmt.Println("***************************************")
	fmt.Println("          Adder for MATFLO             ")
	fmt.Println("                                       ")
	fmt.Println("Used for easily adding database fields ")
	fmt.Println(" and more to matflo systems.           ")
	fmt.Println("                                       ")
	fmt.Println("                                       ")
	fmt.Println("USAGE:                                 ")
	fmt.Println("   adder <Action> <Option1> <Option2...")
	fmt.Println("                                       ")
	fmt.Println("ACTIONS                                ")
	fmt.Println("                                       ")
	fmt.Println("   --help                              ")
	fmt.Println("       Displays this message           ")
	fmt.Println("                                       ")
	fmt.Println("                                       ")
	fmt.Println("   ADD_REC_FIELD                       ")
	fmt.Println("       Options: <1> Database (TM)      ")
	fmt.Println("                <2> Field Name (Flag1) ")
	fmt.Println("                <3> Field Type (int)   ")
	fmt.Println("                <4> Optional Comment   ")
	fmt.Println("                                       ")
	fmt.Println("   ADD_MEM_FIELD                       ")
	fmt.Println("       Options: <1> Database (TM)      ")
	fmt.Println("                <2> Field Name (Flag1) ")
	fmt.Println("                <3> Field Type (int)   ")
	fmt.Println("                <4> Optional Comment   ")
	fmt.Println("                                       ")
	fmt.Println("   ADD_NONDB_FIELD                     ")
	fmt.Println("       Options: <1> Database (TM)      ")
	fmt.Println("                <2> Field Name (Flag1) ")
	fmt.Println("                <3> Field Type (int)   ")
	fmt.Println("                <4> Optional Comment   ")
	fmt.Println("                                       ")
	fmt.Println("                                       ")
	fmt.Println("***************************************")

}

/**
 *  function to determine program action from a given string argument
 */
func actionFromString(arg string) int {

	switch arg {
	case "ADD_REC_FIELD":
		return actionAddRecField
	case "ADD_MEM_FIELD":
		return actionAddMemField
	case "ADD_NONDB_FIELD":
		return actionAddNonDbField
	case "CHAIN_DBS":
		return actionChainDatabases
	default:
		return -1 //default
	}
}

/**
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
 *  TODO: validation on comments
 */
func validateComment(comment string) string {
	return comment
}

/**
 *  TODO: validation on field type?
 */
func validateFieldType(fieldType string) string {
	return fieldType

	//TODO: generate a list here of all types in the matflo system... how?
	// in the meantime, just allow anything.. it's just text, after all
}

/**
 *  TODO: validation on field names? make them consistent?
 *
 */
func validateFieldName(fieldName string) string {
	return fieldName
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
	fmt.Println("Searching for file: ", filePath)

	if _, err := os.Stat(filePath); err == nil {
		return filePath
	} else {
		return ""
	}
} /* findDbDefinitionsFileInDir */

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
				newLine = fmt.Sprintf("  /* DEFNONDBFIELD %s %s; */ // %s", fieldType, fieldName, commentStr)
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
	uid := os.Getuid()
	gid := os.Getgid()

	// Restore the original owner to the new file
	fmt.Println("Chowning", filePath, "to UID:", uid, "GID:", gid)
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
