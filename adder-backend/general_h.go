package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Define all the actions here
// adding one here requires adding  handling later
const (
	displayHelp          = -1
	actionAddRecField    = 1
	actionAddMemField    = 2
	actionAddNonDbField  = 3
	actionChainDatabases = 4
	actionGetMemFields   = 5
	actionGetRecFields   = 6
	actionGetNondbFields = 7
	actionGetAllFields   = 8
)

const (
	printModeNone   = -1
	printModeScreen = 0
	printModeLog    = 1
	noOfPrintModes
)

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
	case "GET_MEM_FIELDS":
		return actionGetMemFields
	case "GET_REC_FIELDS":
		return actionGetRecFields
	case "GET_NONDB_FIELDS":
		return actionGetNondbFields
	case "GET_ALL_FIELDS":
		return actionGetAllFields
	default:
		return -1 //default
	}
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
 *	Set up log file(s)
 */
func setUpLogFiles(printMode int) error {

	// Sometimes we don't need log files...
	if printMode == printModeNone ||
		printMode == printModeScreen {
		return nil
	}

	zeroLogFile, err := os.OpenFile(
		"adder_backend.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664)

	if err != nil {
		return err
	}

	multi := zerolog.MultiLevelWriter(os.Stdout, zeroLogFile)

	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	log.Info().Msg("Log file created successfully")
	return nil
}

/**
 * Basic info logging for whole program
 */
func printOrLog(printMode int, msg ...interface{}) {

	switch printMode {
	case printModeNone:
		return
	case printModeScreen:
		fmt.Println(msg...)
		return
	case printModeLog:
		log.Info().Msg(fmt.Sprint(msg...))
		return

	default:
		return
	}
} /* printOrLog */
