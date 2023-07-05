package main

import "fmt"

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
