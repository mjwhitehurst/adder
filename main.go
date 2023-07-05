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
	printMode := printModeNone
	action := -1

	printMode = printModeScreen

	printOrLog(printMode, "-- Starting Process -- ")

	// Check the number of command-line arguments
	if len(os.Args) < 2 {
		runMode = runModeServer
		printOrLog(printMode, "run mode: Server")
	} else {
		runMode = runModeCmdLine
		printOrLog(printMode, "run mode: Cmd Line")
	}

	var newDbFieldAddition DbFieldAddition
	var err error

	/* If we're command-line run, then work out what we're doing. */
	if runMode == runModeCmdLine {
		newDbFieldAddition, err = dbFieldAdditionFromCmdLine()

		if err != nil {
			printOrLog(printMode, "ERROR: ", err.Error())
			printHelp(printMode)
			return
		}
		printOrLog(printMode, "got field addition from command line")

	} else if runMode == runModeServer { // we are running as a server
		// TODO: get the action with other data from an HTTP request
		action = actionAddRecField
		/* BLOCK HERE TO CHECK SERVER*/

	} else { //I HAVE NO IDEA WHAT IM DOING, JUST PRINT HELP!
		printHelp(printMode)
		return
	}

	switch action {
	case actionAddRecField:
		printOrLog(printMode, "Adding REC field. File: ", newDbFieldAddition.filePath,
			"Field Name: ", newDbFieldAddition.fieldName,
			"Field Type: ", newDbFieldAddition.fieldType,
			"Comment:", newDbFieldAddition.comment)
		addRecField(newDbFieldAddition.filePath,
			newDbFieldAddition.fieldName,
			newDbFieldAddition.fieldType,
			newDbFieldAddition.comment)
		break
	case actionAddMemField:
		printOrLog(printMode, "Adding MEM field. File: ", newDbFieldAddition.filePath,
			"Field Name: ", newDbFieldAddition.fieldName,
			"Field Type: ", newDbFieldAddition.fieldType,
			"Comment:", newDbFieldAddition.comment)
		addMemField(newDbFieldAddition.filePath,
			newDbFieldAddition.fieldName,
			newDbFieldAddition.fieldType,
			newDbFieldAddition.comment)
		break
	case actionAddNonDbField:
		printOrLog(printMode, "Adding NONDB field. File: ", newDbFieldAddition.filePath,
			"Field Name: ", newDbFieldAddition.fieldName,
			"Field Type: ", newDbFieldAddition.fieldType,
			"Comment:", newDbFieldAddition.comment)
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

func printHelp(printMode int) {

	if printMode == printModeScreen {
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
}
