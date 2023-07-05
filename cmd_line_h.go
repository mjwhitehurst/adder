package main

import (
	"errors"
	"fmt"
	"os"
)

/*#######################################################*/
/* Variables / Declarations                              */
/*#######################################################*/

/*#######################################################*/
/* Functions                                             */
/*#######################################################*/

/**
 *  Find a definitions file from a command line argument.
 */
func findFilePathCmdLine() string {

	//PASSED IN BY DOCKER USING -v ARGUMENT
	sourceDir := "/app/sourcedir"
	stringArg1 := os.Args[2]

	//check definitions file
	definitionsFile := findDbDefinitionsFileInDir(stringArg1, sourceDir)

	if definitionsFile == "" {
		fmt.Println("Couldn't find definitions file from ", stringArg1)
	} else {
		fmt.Println("Found file ", definitionsFile, "from ", stringArg1)
	}

	return definitionsFile
} /* findFilePathCmdLine */

/**
 *	Go through the expected command line arguments and run their validation
 *		functions.
 */
func validateCmdLineArgs(dataStruct *DbFieldAddition) error {

	if len(os.Args) < 4 {
		return errors.New("Invalid number of arguments - use --help to find out more")
	}

	//dont worry about os.Args[2] - it is checked as a db name later
	stringArg2 := os.Args[3]
	stringArg3 := os.Args[4]
	stringArg4 := ""

	//have we got a comment?
	if len(os.Args) == 6 {
		stringArg4 = os.Args[5]
	}

	//Check arg 2 - field name
	fieldName := validateFieldName(stringArg2)
	if fieldName == "" {

	}

	//Check arg 3 - type
	fieldType := validateFieldType(stringArg3)
	if fieldType == "" {

	}

	//Check arg 4, if necessary
	commentStr := ""
	if stringArg4 != "" {
		commentStr = validateComment(stringArg4)
	}

	//Success
	fmt.Println("field name: ", fieldName, " field type: ", fieldType, " comment: ", commentStr)

	dataStruct.fieldName = fieldName
	dataStruct.fieldType = fieldType
	dataStruct.comment = commentStr

	return nil
} /* validateCmdLineArgs */
