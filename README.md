# logger
Go logger

![Go](https://github.com/ermanimer/logger/workflows/Go/badge.svg)

## Features
logger writes logs to a file.

## Installation
```bash
go get -u github.com/ermanimer/logger
```

## Prefixes
| Constant      | Value   | Description                 |
| :------------ | :------ | :-------------------------- |
| debugPrefix   | Debug   | Prefix for debug messages   |
| infoPrefix    | Info    | Prefix for info messages    |
| warningPrefix | Warning | Prefix for warning messages |
| errorPrefix   | Error   | Prefix for error messages   |
| fatalPrefix   | Fatal   | Prefix for fatal messages   |

## Trace Levels
| Constant          | Value  | Description                      |
| :------------     | :----: | :------------------------------- |
| DebugTraceLevel   | 0      | Trace level for debug messages   |
| InfoTraceLevel    | 1      | Trace level for info messages    |
| WarningTraceLevel | 2      | Trace level for warning messages |
| ErrorTraceLevel   | 3      | Trace level for error messages   |
| FatalTraceLevel   | 4      | Trace level for fatal messages   |

## Time Format
| Constant   | Value        | Description |
| :--------- | :----------: | :---------- |
| timeFormat | time.RFC3339 | Time format |

## New Line
| Constant | Value | Description                  |
| :------- | :-----| :--------------------------- |
| newLine  | \n    | New line character for Linux |

## FileMode
| Constant | Value | Description           |
| :------- | :-----| :-------------------- |
| fileMode | 0644  | File mode (rw-r--r--) |

## Default Filename and Trace Level
| Constant          | Value           | Description         |
| :---------------- | :-------------- | :------------------ |
| defaultFilename   | logs            | Default file name   |
| defaultTraceLevel | DebugTraceLevel | Default trace level |

## Usage
```go
package main

import (
	log "github.com/ermanimer/logger"
)

func main() {
	//optional: initialize logger
	log.Initialize("new_log_file", log.DebugTraceLevel)

	//log debug message
	log.Debug("This is a debug message.")

	//log formatted debug message
	log.Debugf("This is a %v debug message", "formatted")

	//log info message
	log.Info("This is an info message.")

	//log formatted info message
	log.Infof("This is a %v info message.", "formatted")

	//log warning message
	log.Warning("This is a warning message.")

	//log formatted warning message
	log.Warningf("This is a %v warning message.", "formatted")

	//log error message
	log.Error("This is an error message!")

	//log formatted error message
	log.Errorf("This is a %v error message!", "formatted")

	//log fatal message and call os.Exit(1)
	log.Fatal("This is a fatal message!")

	//log formatted fatal message and call os.Exit(1)
	log.Fatalf("This is a %v fatal message!", "formatted")
}

```

## Terminal Output
![Terminal Output](/images/terminal_output.png)

## Terminal Output With [log-viewer](https://github.com/ermanimer/log-viewer)
![Terminal Output](/images/terminal_output_with_log-viewer.png)

