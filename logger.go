package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

//prefixes
const (
	debugPrefix   = "Debug"
	infoPrefix    = "Info"
	warningPrefix = "Warning"
	errorPrefix   = "Error"
	fatalPrefix   = "Fatal"
)

//trace levels
const (
	DebugTraceLevel = iota
	InfoTraceLevel
	WarningTraceLevel
	ErrorTraceLevel
	FatalTraceLevel
)

//time format
const (
	timeFormat = time.RFC3339
)

//new line character for Linux
const (
	newLine = "\n"
)

//file mode
const (
	fileMode = 0644
)

//default filename and default trace level
const (
	filename   = "default.log"
	traceLevel = DebugTraceLevel
)

type logger struct {
	filename   string
	traceLevel int
	mutex      *sync.Mutex
}

var instance *logger

func Initialize(filename string, traceLevel int) {
	//create instance with parameters
	instance = &logger{
		filename:   filename,
		traceLevel: traceLevel,
		mutex:      &sync.Mutex{},
	}
}

func Debug(values ...interface{}) {
	messageFormat := createMessageFormat(values...)
	log(DebugTraceLevel, debugPrefix, messageFormat, values...)
}

func Debugf(messageFormat string, values ...interface{}) {
	log(DebugTraceLevel, debugPrefix, messageFormat, values...)
}

func Info(values ...interface{}) {
	messageFormat := createMessageFormat(values...)
	log(InfoTraceLevel, infoPrefix, messageFormat, values...)
}

func Infof(messageFormat string, values ...interface{}) {
	log(InfoTraceLevel, infoPrefix, messageFormat, values...)
}

func Warning(values ...interface{}) {
	messageFormat := createMessageFormat(values...)
	log(WarningTraceLevel, warningPrefix, messageFormat, values...)
}

func Warningf(messageFormat string, values ...interface{}) {
	log(WarningTraceLevel, warningPrefix, messageFormat, values...)
}

func Error(values ...interface{}) {
	messageFormat := createMessageFormat(values...)
	log(ErrorTraceLevel, errorPrefix, messageFormat, values...)
}

func Errorf(messageFormat string, values ...interface{}) {
	log(ErrorTraceLevel, errorPrefix, messageFormat, values...)
}

func Fatal(values ...interface{}) {
	messageFormat := createMessageFormat(values...)
	log(FatalTraceLevel, fatalPrefix, messageFormat, values...)
	os.Exit(1)
}

func Fatalf(messageFormat string, values ...interface{}) {
	log(FatalTraceLevel, fatalPrefix, messageFormat, values...)
	os.Exit(1)
}

func checkInitialization() {
	if instance == nil {
		//create instance with default parameters
		instance = &logger{
			filename:   filename,
			traceLevel: traceLevel,
			mutex:      &sync.Mutex{},
		}
	}
}

func createMessageFormat(values ...interface{}) string {
	//create message format
	messageFormat := strings.Repeat("%v, ", len(values))
	//trim last coma and white space
	messageFormat = strings.Trim(messageFormat, ", ")
	return messageFormat
}

func getCallerFunction() string {
	callerFunctionPointer, _, _, _ := runtime.Caller(3)
	callerFunction := runtime.FuncForPC(callerFunctionPointer).Name()
	//trim package path
	callerFunctionParts := strings.Split(callerFunction, "/")
	callerFunction = callerFunctionParts[len(callerFunctionParts)-1]
	return callerFunction
}

func log(traceLevel int, prefix string, messageFormat string, values ...interface{}) {
	//check initialization
	checkInitialization()
	//check trace level
	if instance.traceLevel > traceLevel {
		return
	}
	//synchronization
	instance.mutex.Lock()
	defer instance.mutex.Unlock()
	//create message
	message := fmt.Sprintf(messageFormat, values...)
	//replace new line characters with white spaces
	message = strings.Replace(message, newLine, " ", -1)
	//create formatted message
	message = fmt.Sprintf("[%s][%s][%s][%s]%s", time.Now().Format(timeFormat), prefix, getCallerFunction(), message, newLine)
	//open log file
	logFile, err := os.OpenFile(instance.filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, fileMode)
	if err != nil {
		printError("Opening log file failed!")
		return
	}
	defer func() {
		err := logFile.Close()
		if err != nil {
			printError("Closing log file failed!")
		}
	}()
	//write to log file
	_, err = logFile.WriteString(message)
	if err != nil {
		printError("Writing to log file failed!")
	}
}

func printError(message string) {
	fmt.Printf("logger: %s%s", message, newLine)
}
