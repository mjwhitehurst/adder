package main

import (
	"errors"
	"os"
)

/*#######################################################*/
/* Functions                                             */
/*#######################################################*/

func dbFieldAdditionFromCmdLine() (DbFieldAddition, error) {
	returnStruct := DbFieldAddition{}
	action := -1
	var err error

	// Access the command-line argument
	firstArg := os.Args[1]

	// Process the argument and display the result
	action = actionFromString(firstArg)

	if action == -1 {
		return returnStruct, errors.New("unable to determine action from argument")
	}

	// we are running as a command line action - look for our file from arguments.
	returnStruct.filePath, err = findFilePathCmdLine()
	if err != nil {
		return returnStruct, err
	}

	err = validateCmdLineArgsFieldAddition(&returnStruct)
	if err != nil {
		return returnStruct, err
	}

	if err != nil || returnStruct.filePath == "" {
		return returnStruct, errors.New("unable to build db addition struct from command line")
	}

	return returnStruct, nil
} /* dbFieldAdditionFromCmdLine */

func dbMemInfoFromCmdLine() (string, error) {

	filePath, err := findFilePathCmdLine()
	if err != nil {
		return "", err
	}

	memFields, err := findMemFields(filePath)

	if err != nil {
		return "", errors.New("failed to find mem fields")
	}

	fieldsStr := fieldsToString(memFields)
	return fieldsStr, nil
}

func dbRecInfoFromCmdLine() (string, error) {

	filePath, err := findFilePathCmdLine()
	if err != nil {
		return "", err
	}

	recFields, err := findRecFields(filePath)

	if err != nil {
		return "", errors.New("failed to find rec fields")
	}

	fieldsStr := fieldsToString(recFields)
	return fieldsStr, nil
}

func dbNondbInfoFromCmdLine() (string, error) {

	filePath, err := findFilePathCmdLine()
	if err != nil {
		return "", err
	}

	nondbFields, err := findNondbFields(filePath)

	if err != nil {
		return "", errors.New("failed to find nondb fields")
	}

	fieldsStr := fieldsToString(nondbFields)
	return fieldsStr, nil
}

/**
 *  Find a definitions file from a command line argument.
 */
func findFilePathCmdLine() (string, error) {

	//ACTUAL $SRC PASSED IN BY DOCKER USING -v ARGUMENT
	//Defined by using "-v $SRC:/app/sourcedir"
	sourceDir := "/app/sourcedir"
	stringArg1 := os.Args[2]
	var err error

	//check definitions file
	definitionsFile := findDbDefinitionsFileInDir(stringArg1, sourceDir)

	if definitionsFile == "" {
		err = errors.New("couldn't find definitions file from arg 2")
	}

	return definitionsFile, err
} /* findFilePathCmdLine */

/**
 *	Go through the expected command line arguments and run their validation
 *		functions.
 */
func validateCmdLineArgsFieldAddition(dataStruct *DbFieldAddition) error {
	if len(os.Args) < 4 {
		return errors.New("invalid number of arguments - use --help to find out more")
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
	dataStruct.fieldName = fieldName
	dataStruct.fieldType = fieldType
	dataStruct.comment = commentStr

	return nil
} /* validateCmdLineArgsFieldAddition */

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
