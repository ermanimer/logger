# logger
Go logger

![Go](https://github.com/ermanimer/logger/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ermanimer/logger)](https://goreportcard.com/report/github.com/ermanimer/logger)

## Features
logger writes logs to the specified log file. Use [log_viewer](https://github.com/ermanimer/log_viewer) for the best viewing experince.

## Installation
```bash
go get -u github.com/ermanimer/logger
```

## Prefixes
|Constant     |Value  |
|:------------|:-----:|
|debugPrefix  |Debug  |
|infoPrefix   |Info   |
|warningPrefix|Warning|
|errorPrefix  |Error  |
|fatalPrefix  |Fatal  |

## Trace Levels
|Constant         |Value |
|:----------------|:----:|
|DebugTraceLevel  |1     |
|InfoTraceLevel   |2     |
|WarningTraceLevel|3     |
|ErrorTraceLevel  |4     |
|FatalTraceLevel  |5     |

## Time Format
|Constant  |Value       |
|:---------|:----------:|
|timeFormat|time.RFC3339|

## FileMode
|Constant|Value           |
|:-------|:--------------:|
|fileMode|0644 (rw-r--r--)|

## Default Parameters
|Parameter        |Value          |
|:----------------|:-------------:|
|defaultFilename  |default.log    |
|defaultTraceLevel|DebugTraceLevel|

## Usage
```go
package main

import (
	"github.com/ermanimer/logger/v2"
)

func main() {
	//create a logger with DefaultLogger function
	l := logger.DefaultLogger()
	
	//or with NewLoggerFunction
	//l := logger.NewLogger("default.log", logger.DebugTraceLevel)

	//log debug message
	l.Debug("this is a debug message")

	//log formatted debug message
	l.Debugf("this is a %v debug message", "formatted")

	//log info message
	l.Info("this is an info message")

	//log formatted info message
	l.Infof("this is a %v info message", "formatted")

	//log warning message
	l.Warning("this is a warning message")

	//log formatted warning message
	l.Warningf("this is a %v warning message", "formatted")

	//log error message
	l.Error("this is an error message")

	//log formatted error message
	l.Errorf("this is a %v error message", "formatted")

	//log fatal message and call os.Exit(1)
	l.Fatal("this is a fatal message")

	//log formatted fatal message and call os.Exit(1)
	l.Fatalf("this is a %v fatal message", "formatted")
}
```

## Terminal Output With [log-viewer](https://github.com/ermanimer/log-viewer)
![Terminal Output](/images/terminal_output.png)
