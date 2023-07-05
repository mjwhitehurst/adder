package main

import "fmt"

// Define all the actions here
// adding one here requires adding  handling later
const (
	displayHelp          = -1
	actionAddRecField    = 1
	actionAddMemField    = 2
	actionAddNonDbField  = 3
	actionChainDatabases = 4
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
	default:
		return -1 //default
	}
}

/**
 * TODO: make this log somewhere for when used as a server
 */
func printOrLog(printMode int, msg ...interface{}) {

	switch printMode {
	case printModeNone:
		return
	case printModeScreen:
		fmt.Println(msg...)
		return
	case printModeLog:
		return

	default:
		return
	}
} /* printOrLog */
