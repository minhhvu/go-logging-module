package go_logging_module

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

/**
 * A record contains all fields and values of a log that will be printed out.
 */
type Record struct {
	ProjectName    string    `json:"projectName"`
	Environment    string    `json:"Environment"`
	Module         string    `json:"module"`
	FunctionName   string    `json:"FunctionName"`
	SequenceNumber uuid.UUID `json:"sequenceNumber"`
	XRayId         string    `json:"XRayId"`
	ObjectId       string    `json:"ObjectId"`
	Error          error     `json:"Error"`
	Message        string    `json:"message"`
}

/**
 * OutputType determines how a log will be outputted. It can be either json or flat.
 */
type OutputType string

const (
	jsonType OutputType = "json"
	flatType OutputType = "flat"
)

var logRecord Record
var outputLogType OutputType

/**
 * Initialization Phase:
 * Initializes parameters that do not change across the Lambda's execution.
 * @param projectName The main project name (e.g. Rails/IST, Travellers/INS, etc.)
 * @param environment The current deployment environment (e.g. dev, test, uat, prod, etc.)
 * @param module The name of the module (e.g. JPM Connector, JPM Treasury Service, etc.)
 * @param functionName The name of the function (e.g. AddCardHolderData, SaveCardHolderData, etc.)
 */
func Initialize(projectName string, env string, module string, functionName string, outputType OutputType, header ...*http.Header) {
	logRecord.ProjectName = projectName
	logRecord.Environment = env
	logRecord.Module = module
	logRecord.FunctionName = functionName
	logRecord.SequenceNumber = uuid.New()
	if len(header) > 0 {
		logRecord.XRayId = header[0].Get("X-Amzn-Trace-Id")
	} else {
		logRecord.XRayId = "N/A"
	}
	outputLogType = outputType
}

/**
 * Generate a INFO log record
 * @param sequenceNum Can be used to represent object id, reference id
 * @message The actual message itself (e.g. Invalid user ID, payment does not exist, etc.).
 */
func Info(objectId string, message string) {
	setLogMessageParameters(objectId, message)
	PrintOutLog()
}

/**
 * Generate a ERROR log record
 * @param sequenceNum Can be used to represent object id, reference id
 * @message The actual message itself (e.g. Invalid user ID, payment does not exist, etc.).
 */
func Error(objectId string, err error, message string) {
	setLogErrorParameters(objectId, err, message)
	PrintOutLog()
}

/**
 * Change the log output type to FLAT
 */
func SetOutputTypeToFlat() {
	outputLogType = flatType
}

/**
 * Change the log output type to JSON
 */
func SetOutputTypeToJson() {
	outputLogType = jsonType
}

/**
 * Print out the log
 */
func PrintOutLog() {
	if outputLogType == jsonType {
		logRecordJson, err := json.Marshal(logRecord)
		if err != nil {
			log.Println("ERROR: function = json.Marshal, object =", logRecord, ", err= ", err)
		}
		log.Print(string(logRecordJson))
	} else {
		// Prints out the appropriate log depending on if an error was passed to the logs or not.
		if logRecord.Error != nil {
			log.Printf("<ProjectName = %s, Environment = %s, Module = %s, FunctionName = %s, X-RayID = %s>, [SequenceNumber = %s] [ObjectId: %s] %s",
				logRecord.ProjectName, logRecord.Environment, logRecord.Module, logRecord.FunctionName,
				logRecord.XRayId, logRecord.SequenceNumber, logRecord.ObjectId, logRecord.Message)
		} else {
			log.Printf("<ProjectName = %s, Environment = %s, Module = %s, FunctionName = %s, X-RayID = %s>, [SequenceNumber = %s] [ObjectId: %s] %s [Error = %v]",
				logRecord.ProjectName, logRecord.Environment, logRecord.Module, logRecord.FunctionName,
				logRecord.XRayId, logRecord.SequenceNumber, logRecord.ObjectId, logRecord.Message, logRecord.Error)
		}
	}
}

func setLogMessageParameters(objectId string, message string) {
	logRecord.ObjectId = objectId
	logRecord.Message = message
	logRecord.Error = nil
}

func setLogErrorParameters(objectId string, err error, message string) {
	logRecord.ObjectId = objectId
	logRecord.Message = message
	logRecord.Error = err
}
