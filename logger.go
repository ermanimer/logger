package log

import (
	"fmt"
	"os"
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
	checkInitialization()
	if instance.traceLevel <= DebugTraceLevel {
		instance.log(debugPrefix, message)
	}
}

func Debugf(messageFormat string, values ...interface{}) {
	if instance.traceLevel <= DebugTraceLevel {
		message := fmt.Sprintf(messageFormat, values...)
		instance.log(debugPrefix, message)
	}
}

func Info(message string) {
	if instance.traceLevel <= InfoTraceLevel {
		instance.log(infoPrefix, message)
	}
}

func Infof(messageFormat string, values ...interface{}) {
	if instance.traceLevel <= InfoTraceLevel {
		message := fmt.Sprintf(messageFormat, values...)
		instance.log(infoPrefix, message)
	}
}

func Warning(message string) {
	if instance.traceLevel <= WarningTraceLevel {
		instance.log(warningPrefix, message)
	}
}

func Warningf(messageFormat string, values ...interface{}) {
	if instance.traceLevel <= WarningTraceLevel {
		message := fmt.Sprintf(messageFormat, values...)
		instance.log(warningPrefix, message)
	}
}

func Error(message string) {
	if instance.traceLevel <= ErrorTraceLevel {
		instance.log(errorPrefix, message)
	}
}

func Errorf(messageFormat string, values ...interface{}) {
	if instance.traceLevel <= ErrorTraceLevel {
		message := fmt.Sprintf(messageFormat, values...)
		instance.log(errorPrefix, message)
	}
}

func Fatal(message string) {
	if instance.traceLevel <= FatalTraceLevel {
		instance.log(fatalPrefix, message)
		os.Exit(1)
	}
}

func Fatalf(messageFormat string, values ...interface{}) {
	if instance.traceLevel <= FatalTraceLevel {
		message := fmt.Sprintf(messageFormat, values...)
		instance.log(fatalPrefix, message)
		os.Exit(1)
	}
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

func (instance *logger) log(prefix string, message string) {
	//synchronization
	instance.mutex.Lock()
	defer instance.mutex.Unlock()

	//create formatted message
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
