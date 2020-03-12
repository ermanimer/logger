package log

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

//prefixes
const (
	debugPrefix   = "Debug  "
	infoPrefix    = "Info   "
	warningPrefix = "Warning"
	errorPrefix   = "Error  "
	fatalPrefix   = "Fatal  "
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
	defaultFileName   = "logs"
	defaultTraceLevel = DebugTraceLevel
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

func Debug(message string) {
	log(DebugTraceLevel, debugPrefix, message)
}

func Debugf(messageFormat string, values ...interface{}) {
	log(DebugTraceLevel, debugPrefix, messageFormat, values...)
}

func Info(message string) {
	log(InfoTraceLevel, infoPrefix, message)
}

func Infof(messageFormat string, values ...interface{}) {
	log(InfoTraceLevel, infoPrefix, messageFormat, values...)
}

func Warning(message string) {
	log(WarningTraceLevel, warningPrefix, message)
}

func Warningf(messageFormat string, values ...interface{}) {
	log(WarningTraceLevel, warningPrefix, messageFormat, values...)
}

func Error(message string) {
	log(ErrorTraceLevel, errorPrefix, message)
}

func Errorf(messageFormat string, values ...interface{}) {
	log(ErrorTraceLevel, errorPrefix, messageFormat, values...)
}

func Fatal(message string) {
	log(FatalTraceLevel, fatalPrefix, message)
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
			filename:   defaultFileName,
			traceLevel: defaultTraceLevel,
			mutex:      &sync.Mutex{},
		}
	}
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

	//create formatted message
	message := fmt.Sprintf(messageFormat, values...)
	message = strings.Replace(message, newLine, " ", -1)
	formattedMessage := fmt.Sprintf("[%s][%s]: %s%s", time.Now().Format(timeFormat), prefix, message, newLine)

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
	_, err = logFile.WriteString(formattedMessage)
	if err != nil {
		printError("Writing to log file failed!")
	}
}

func printError(message string) {
	fmt.Println(fmt.Sprintf("go_logger: %s%s", message, newLine))
}
