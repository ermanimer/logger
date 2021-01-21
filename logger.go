package logger

import (
	"fmt"
	"os"
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
	DebugTraceLevel = iota + 1
	InfoTraceLevel
	WarningTraceLevel
	ErrorTraceLevel
	FatalTraceLevel
)

//time format
const (
	timeFormat = time.RFC3339
)

//file mode
const (
	fileMode = 0644
)

//default parameters
const (
	defaultFilename   = "default.log"
	defaultTraceLevel = DebugTraceLevel
)

type Logger struct {
	Filename   string
	TraceLevel int
	mutex      *sync.Mutex
}

func DefaultLogger() *Logger {
	return &Logger{
		Filename:   defaultFilename,
		TraceLevel: defaultTraceLevel,
		mutex:      &sync.Mutex{},
	}
}

func NewLogger(filename string, traceLevel int) *Logger {
	l := &Logger{
		Filename:   filename,
		TraceLevel: traceLevel,
		mutex:      &sync.Mutex{},
	}
	return l
}

func (l *Logger) Debug(values ...interface{}) {
	messageFormat := createMessageFormat(values...)
	l.log(DebugTraceLevel, debugPrefix, messageFormat, values...)
}

func (l *Logger) Debugf(messageFormat string, values ...interface{}) {
	l.log(DebugTraceLevel, debugPrefix, messageFormat, values...)
}

func (l *Logger) Info(values ...interface{}) {
	messageFormat := createMessageFormat(values...)
	l.log(InfoTraceLevel, infoPrefix, messageFormat, values...)
}

func (l *Logger) Infof(messageFormat string, values ...interface{}) {
	l.log(InfoTraceLevel, infoPrefix, messageFormat, values...)
}

func (l *Logger) Warning(values ...interface{}) {
	messageFormat := createMessageFormat(values...)
	l.log(WarningTraceLevel, warningPrefix, messageFormat, values...)
}

func (l *Logger) Warningf(messageFormat string, values ...interface{}) {
	l.log(WarningTraceLevel, warningPrefix, messageFormat, values...)
}

func (l *Logger) Error(values ...interface{}) {
	messageFormat := createMessageFormat(values...)
	l.log(ErrorTraceLevel, errorPrefix, messageFormat, values...)
}

func (l *Logger) Errorf(messageFormat string, values ...interface{}) {
	l.log(ErrorTraceLevel, errorPrefix, messageFormat, values...)
}

func (l *Logger) Fatal(values ...interface{}) {
	messageFormat := createMessageFormat(values...)
	l.log(FatalTraceLevel, fatalPrefix, messageFormat, values...)
	os.Exit(1)
}

func (l *Logger) Fatalf(messageFormat string, values ...interface{}) {
	l.log(FatalTraceLevel, fatalPrefix, messageFormat, values...)
	os.Exit(1)
}

func createMessageFormat(values ...interface{}) string {
	messageFormat := strings.Repeat("%v, ", len(values))
	messageFormat = strings.Trim(messageFormat, ", ")
	return messageFormat
}

func (l *Logger) log(traceLevel int, prefix string, messageFormat string, values ...interface{}) {
	//check trace level
	if l.TraceLevel > traceLevel {
		return
	}
	//synchronization
	l.mutex.Lock()
	defer l.mutex.Unlock()
	//create message
	message := fmt.Sprintf(messageFormat, values...)
	//replace new line characters with white spaces
	message = strings.Replace(message, "\n", " ", -1)
	//create formatted message
	message = fmt.Sprintf("[%s][%s][%s]\n", time.Now().Format(timeFormat), prefix, message)
	//open log file
	logFile, err := os.OpenFile(l.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, fileMode)
	if err != nil {
		fmt.Println("opening log file failed")
		return
	}
	defer func() {
		err := logFile.Close()
		if err != nil {
			fmt.Println("closing log file failed")
		}
	}()
	//write to log file
	_, err = logFile.WriteString(message)
	if err != nil {
		fmt.Println("writing log failed")
	}
}
