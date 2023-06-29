package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"syscall"
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

	} else { // we are running as a server
		//TODO: get the action with other data from a http request
		action = actionAddRecField
	}
	/* BLOCK HERE TO CHECK SERVER*/

	switch action {
	case actionAddRecField:
		addRecField()
		break
	case actionAddMemField:
		addMemField()
		break
	case actionAddNonDbField:
		addNonDbField()
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
func addDatabaseField(tag string) {

	if len(os.Args) < 4 {
		fmt.Println("Invalid number of arguments - use --help to find out more")
		return
	}

	//PASSED IN BY DOCKER USING -v ARGUMENT
	sourceDir := "/app/sourcedir"

	stringArg1 := os.Args[2]
	stringArg2 := os.Args[3]
	stringArg3 := os.Args[4]
	stringArg4 := ""

	//have we got a comment?
	if len(os.Args) == 6 {
		stringArg4 = os.Args[5]
	}

	//check definitions file
	definitionsFile := findDbDefinitionsFileInDir(stringArg1, sourceDir)

	if definitionsFile == "" {
		fmt.Println("Couldn't find definitions file from ", stringArg1)
	} else {
		fmt.Println("Found file ", definitionsFile, "from ", stringArg1)
	}

	//check arg 2 - field name
	fieldName := validateFieldName(stringArg2)
	if fieldName == "" {

	}

	//check arg 3 - type
	fieldType := validateFieldType(stringArg3)
	if fieldType == "" {

	}

	//check arg 4, if necessary
	commentStr := ""
	if stringArg4 != "" {
		commentStr = validateComment(stringArg4)
	}

	fmt.Println("file: ", definitionsFile, " field name: ", fieldName, " field type: ", fieldType, " comment: ", commentStr)

	//At this point we have a file, and we know what we want to put into it.
	err := addFieldAfterTag(definitionsFile, tag, fieldName, fieldType, commentStr)
	if err != nil {
		fmt.Println("Error adding field: ", err)
	}

} /* addDatabaseField */

func addRecField() {
	addDatabaseField("ADDER REC START")
}

func addMemField() {
	addDatabaseField("ADDER MEM START")
	return
}

func addNonDbField() {
	//TODO: like addDatabaseField, but add the whole thing as a comment
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
 *  Looks for a db file in a directory. will try multiple ways
 *  input:
 *          DB_NAME/db_name/db_name_definitions/db_name_definitions.h
 *
 *  looks for (and returns):
 *          db_name_definitions.h
 *
 *  on fail, returns ""
 *
 */
func findDbDefinitionsFileInDir(dbName string, dirName string) string {

	// Check nulls
	if dbName == "" || dirName == "" {
		return ""
	}

	dbFileName := ""

	// Best case scenario, we are given db_name_definitions.h
	if strings.HasSuffix(dbName, "_definitions.h") {
		//hooray
		dbFileName = strings.TrimSuffix(dbName, "_definitions.h")

	} else if strings.HasSuffix(dbName, "_definitions") {
		//hooray
		dbFileName = strings.TrimSuffix(dbName, "_definitions")

	} else {
		//guess! lower case it.
		dbFileName = strings.ToLower(dbName)
	}

	definitionsFile := dbFileName + "_definitions.h"

	filePath := filepath.Join(dirName, definitionsFile)
	fmt.Println("Searching for file: ", filePath)

	if _, err := os.Stat(filePath); err == nil {
		return filePath
	} else {
		return ""
	}
} /* findDbDefinitionsFileInDir */

func addFieldAfterTag(filePath, tag, fieldName, fieldType, commentStr string) error {
	// Create a temporary file
	tempFilePath := filePath + ".tmp"
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

	// Search for the tag and insert the new line after it
	foundTag := false
	for scanner.Scan() {
		line := scanner.Text()
		writer.WriteString(line + "\n")

		if !foundTag && strings.Contains(line, tag) {
			foundTag = true

			// Insert the new line with field details after the tag
			newLine := fmt.Sprintf("    %s %s // %s\n", fieldType, fieldName, commentStr)
			writer.WriteString(newLine)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
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

	// Get the original file info (owner and permissions)
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	// Get the original owner
	originalOwnerUint := fileInfo.Sys().(*syscall.Stat_t).Gid

	fmt.Println(" uint og owner: ", originalOwnerUint)

	// -1 will tell chown to do nothing
	originalOwner := -1

	//quick check
	if originalOwnerUint <= math.MaxInt32 {
		originalOwner = int(originalOwnerUint)
	} else {
		return errors.New("int conversion would lead to loss of data! this is bad!")
	}

	// Restore the original owner to the new file
	fmt.Println("chowning ", filePath, "to ", originalOwner)
	if err := os.Chown(filePath, originalOwner, originalOwner); err != nil {
		fmt.Println("bad chown")
		return err
	}

	return nil
}
