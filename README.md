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

## Usage
```go
package main

import (
	"github.com/ermanimer/logger"
)

func main() {
	//initialize logger
	l := logger.NewLogger("filename.log", logger.DebugTraceLevel)

	//log debug message
	l.Debug("This is a debug message.")

	//log formatted debug message
	l.Debugf("This is a %v debug message.", "formatted")

	//log info message
	l.Info("This is an info message.")

	//log formatted info message
	l.Infof("This is a %v info message.", "formatted")

	//log warning message
	l.Warning("This is a warning message.")

	//log formatted warning message
	l.Warningf("This is a %v warning message.", "formatted")

	//log error message
	l.Error("This is an error message!")

	//log formatted error message
	l.Errorf("This is a %v error message!", "formatted")

	//log fatal message and call os.Exit(1)
	l.Fatal("This is a fatal message!")

	//log formatted fatal message and call os.Exit(1)
	l.Fatalf("This is a %v fatal message!", "formatted")
}

```

## Terminal Output With [log-viewer](https://github.com/ermanimer/log-viewer)
![Terminal Output](/images/terminal_output.png)
