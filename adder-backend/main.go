package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	var err error

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Error().Msg("Error message")
	log.Warn().Msg("Warning message")
	log.Info().Msg("Info message")
	log.Debug().Msg("Debug message")
	log.Trace().Msg("Trace message")

	// Run mode
	runMode := runModeUnknown
	printMode := printModeNone
	action := -1

	printMode = printModeScreen

	err = setUpLogFiles(printMode)
	if err != nil {
		//PANIC
	}

	printOrLog(printMode, "-- Starting Process -- ")
	printOrLog(printMode, "args: ", os.Args)

	// Check the args - lots of args, none of which are 'dlv', mean that we're runnning as command line
	dlvArgFound := false
	for _, arg := range os.Args {
		if arg == "dlv" {
			dlvArgFound = true
			break
		}
	}

	printOrLog(printMode, "arguments: ", len(os.Args), " delve found? [", dlvArgFound, "]")

	if len(os.Args) < 2 || dlvArgFound {
		runMode = runModeServer
		printOrLog(printMode, "run mode: Server")
	} else {
		runMode = runModeCmdLine
		printOrLog(printMode, "run mode: Cmd Line")
	}

	var newDbFieldAddition DbFieldAddition

	/* If we're command-line run, then work out what we're doing. */
	if runMode == runModeCmdLine {

		//determine what we're doing
		firstArg := os.Args[1]

		// Process the argument and display the result
		action = actionFromString(firstArg)

		switch action {
		/* Field additions are all handled by the same functions */
		case actionAddRecField:
			fallthrough
		case actionAddMemField:
			fallthrough
		case actionAddNonDbField:
			newDbFieldAddition, err = dbFieldAdditionFromCmdLine()
		case actionGetMemFields:
			printOrLog(printMode, "MEM Fields:")
			memStr, err := dbMemInfoFromCmdLine()
			if err == nil {
				printOrLog(printMode, memStr)
			} else {
				printOrLog(printMode, "failed to get mem info :( ")
			}
			return
		case actionGetRecFields:
			printOrLog(printMode, "REC Fields:")
			recStr, err := dbRecInfoFromCmdLine()
			if err == nil {
				printOrLog(printMode, recStr)
			} else {
				printOrLog(printMode, "failed to get rec info :( ")
			}
			return
		case actionGetNondbFields:
			printOrLog(printMode, "\nNONDB Fields:")
			nondbStr, err := dbNondbInfoFromCmdLine()
			if err == nil {
				printOrLog(printMode, nondbStr)
			} else {
				printOrLog(printMode, "failed to get nondb info :(")
			}
			return
		case actionGetAllFields:
			printOrLog(printMode, "MEM Fields:")
			memStr, err := dbMemInfoFromCmdLine()
			if err == nil {
				printOrLog(printMode, memStr)
			} else {
				printOrLog(printMode, "failed to get mem info :( ")
			}
			printOrLog(printMode, "\nREC Fields:")
			recStr, err := dbRecInfoFromCmdLine()
			if err == nil {
				printOrLog(printMode, recStr)
			} else {
				printOrLog(printMode, "failed to get rec info :( ")
			}
			printOrLog(printMode, "\nNONDB Fields:")
			nondbStr, err := dbNondbInfoFromCmdLine()
			if err == nil {
				printOrLog(printMode, nondbStr)
			} else {
				printOrLog(printMode, "failed to get nondb info :(")
			}
			return
		//TODO: add more actions here if necessary

		case -1:
			fallthrough
		default:
			printHelp(printMode)
			return
		}

		if err != nil {
			printOrLog(printMode, "ERROR: ", err.Error())
			return
		}

	} else if runMode == runModeServer { // we are running as a server

		runServer()

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

	case actionAddMemField:
		printOrLog(printMode, "Adding MEM field. File: ", newDbFieldAddition.filePath,
			"Field Name: ", newDbFieldAddition.fieldName,
			"Field Type: ", newDbFieldAddition.fieldType,
			"Comment:", newDbFieldAddition.comment)
		addMemField(newDbFieldAddition.filePath,
			newDbFieldAddition.fieldName,
			newDbFieldAddition.fieldType,
			newDbFieldAddition.comment)

	case actionAddNonDbField:
		printOrLog(printMode, "Adding NONDB field. File: ", newDbFieldAddition.filePath,
			"Field Name: ", newDbFieldAddition.fieldName,
			"Field Type: ", newDbFieldAddition.fieldType,
			"Comment:", newDbFieldAddition.comment)
		addNonDbField(newDbFieldAddition.filePath,
			newDbFieldAddition.fieldName,
			newDbFieldAddition.fieldType,
			newDbFieldAddition.comment)

	default:
		printOrLog(printMode, "not given an action - exiting")

	}

	printOrLog(printMode, "-- Finished Process -- ")
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
		fmt.Println("    GET_MEM_FIELDS                     ")
		fmt.Println("    GET_REC_FIELDS                     ")
		fmt.Println("    GET_NONDB_FIELDS                   ")
		fmt.Println("    GET_ALL_FIELDS                     ")
		fmt.Println("        Options: <1> Database (TM)     ")
		fmt.Println("                                       ")
		fmt.Println("                                       ")
		fmt.Println("***************************************")
	}
}
