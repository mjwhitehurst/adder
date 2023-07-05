package main

// Define all the actions here
// adding one here requires adding  handling later
const (
	displayHelp          = -1
	actionAddRecField    = 1
	actionAddMemField    = 2
	actionAddNonDbField  = 3
	actionChainDatabases = 4
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
