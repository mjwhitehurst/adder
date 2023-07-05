package main

import (
	"fmt"
	"os"
)

const (
	runModeUnknown = -1
	runModeServer  = 0
	runModeCmdLine = 1
	noOfRunModes   = 2
)

/*#######################################################*/
/* Main function                                         */
/*#######################################################*/
func main() {
	// Run mode
	runMode := runModeUnknown
	action := -1

	// Check the number of command-line arguments
	if len(os.Args) < 2 {
		runMode = runModeCmdLine
	} else {
		runMode = runModeServer
	}

	var newDbFieldAddition DbFieldAddition
	var err error

	/* If we're command-line run, then work out what we're doing. */
	if runMode == runModeCmdLine {
		newDbFieldAddition, err = dbFieldAdditionFromCmdLine()

		if err != nil {
			printHelp()
			return
		}

	} else if runMode == runModeServer { // we are running as a server
		// TODO: get the action with other data from an HTTP request
		action = actionAddRecField
		/* BLOCK HERE TO CHECK SERVER*/

	} else { //I HAVE NO IDEA WHAT IM DOING, JUST PRINT HELP!
		printHelp()
		return
	}

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
