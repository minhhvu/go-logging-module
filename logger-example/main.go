package main

import (
	"errors"
	logger "go-logging-module"
)

func main() {
	// Initialize logHandler at beginning
	logger.Initialize("test", "dev", "moduleName", "lambdaFunctionName", logger.OutputType("json"))

	// Print out a INFO log with an object id and message
	logger.Info("12123123dfs", "Info Log")

	// Print out a ERROR log with an object id and message
	logger.Error("32434234", errors.New("New error"), "Error log")

	//Change the output of log to FLAT type
	logger.SetOutputTypeToFlat()

	//Print out the current log record
	logger.PrintOutLog()
}
