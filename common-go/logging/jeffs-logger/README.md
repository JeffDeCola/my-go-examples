# JEFFS LOGGER EXAMPLE

_Logging using my
[github.com/JeffDeCola/my-go-packages/golang/logger](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger)
which uses the standard structured logging
[slog](https://pkg.go.dev/log/slog)
library._

tl;r,

```go
logger "github.com/JeffDeCola/my-go-packages/golang/logger"

log := logger.CreateLogger(logger.Trace, "jeffs_noTime", output)

log.Debug("This is a debug message")
log.Info(fmt.Sprintf("Formatted Info Message a=%.2f", a), "a", a, "user", "jeff")

log.ChangeLogLevel(logger.Warning)
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/jeffs-logger/README.md#overview)
* [SET LOG LEVEL, FORMAT AND OUTPUT](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/jeffs-logger/README.md#set-log-level-format-and-output)
* [CHANGE LEVEL](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/jeffs-logger/README.md#change-level)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/jeffs-logger/README.md#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/jeffs-logger/README.md#test)

Documentation and Reference

* [github.com/JeffDeCola/golang/logger](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger)

## OVERVIEW

My logger package uses six levels,

* log.**Fatal**("I'm bailing")
* log.**Error**("Something failed but I'm not quitting.")
* log.**Warning**("You should probably take a look at this.")
* log.**Info**("Something noteworthy happened!")
* log.**Debug**("Useful debugging information.")
* log.**Trace**("Something very low level.")

## SET LOG LEVEL, FORMAT AND OUTPUT

When you create a logger, you have three choices,

* myLevel (int)
  * Trace
  * Debug
  * Info **(default)**
  * Warning
  * Error
  * Fatal
* format (string)
  * "text" - Uses slog text
  * "json" - Uses slog json
  * "jeffs" - My own formatting **(default)**
  * "jeffs_noTime" - My own formatting without time stamp
* output (*os.File)
  * os.Stdout
  * os.Stderr
  * file handler

For example,

```go
    log := logger.CreateLogger(logger.Debug, "json", os.Stdout)
```

## CHANGE LEVEL

You may change level at any time,

```go
log.ChangeLogLevel(logger.Warning)
```

## RUN

Run,

```bash
go run main.go
```

As an example, you can use logger to write to a file.

```go
    // Open or create the log file, output is filename
    output, err := os.OpenFile("jefflog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic("Failed to open/create log file" + err.Error())
    }

    log := logger.CreateLogger(logger.Trace, "jeffs_noTime", output)

    a := 4.54534

    log.Trace("This is a trace message")
    log.Debug("This is a debug message")
    log.Info(fmt.Sprintf("Formatted Info Message a=%.2f", a), "a", a, "user", "jeff")
    log.Warning("This is a Warning Message", "user", "jeff")
    log.Error("This is an Error message")
    // log.Fatal("Fatal Error")

    // Dynamically change log level
    fmt.Printf("\nCHANGE LEVEL\n\n")
    log.ChangeLogLevel(logger.Warning)

    log.Trace("This is a trace message")
    log.Debug("This is a debug message")
    log.Info(fmt.Sprintf("Formatted Info Message a=%.2f", a), "a", a, "user", "jeff")
    log.Warning("This is a Warning Message", "user", "jeff")
    log.Error("This is an Error message")
```

The logfile "jefflog.log" would look like,

```text
[TRACE] This is a trace message
[DEBUG] This is a debug message
[INFO ] Formatted Info Message a=4.55 (a: 4.54534), (user: jeff)
[WARN ] This is a Warning Message (user: jeff)
[ERROR] This is an Error message
[WARN ] This is a Warning Message (user: jeff)
[ERROR] This is an Error message
```

## TEST

To create _test files,

```bash
gotests -w -all main.go
```

To unit test the code,

```bash
go test -cover ./...
```
