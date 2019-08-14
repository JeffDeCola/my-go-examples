# logrus example

`logrus` _package. An example of logging._

Refer to the
[logrus](https://github.com/sirupsen/logrus)
package for more info.  This is not a standard package.

* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus#run)
* [LOG LEVELS](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus#log-levels)
* [USING LOGRUS](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus#using-logrus)
  * [SET LOG LEVEL](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus#set-log-level)
  * [SET FORMAT](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus#set-format)
  * [SET OUTPUT (DEFAULT stderr)](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus#set-output-default-stderr)
  * [LOGGING TO FILE (APPEND) (io.Writer)](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus#logging-to-file-append-iowriter)
  * [NORMAL LOGGING](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus#normal-logging)
  * [NORMAL LOGGING WITH FORMATTING](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus#normal-logging-with-formatting)
  * [USING FIELDS](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus#using-fields)
  * [REUSING FIELDS](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus#reusing-fields)
  
[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

```go
go run logrus.go
```

## LOG LEVELS

Logrus has seven levels in this order,

* log.**Panic**("I'm bailing. Calls panic() after logging.")
* log.**Fatal**("Bye. Calls os.Exit(1) after logging.")  
* log.**Error**("Something failed but I'm not quitting.")
* log.**Warn**("You should probably take a look at this.")
* log.**Info**("Something noteworthy happened!")
* log.**Debug**("Useful debugging information.")
* log.**Trace**("Something very low level.")

## USING LOGRUS

This is a very simple package.

### SET LOG LEVEL

```go
// SET LOG LEVEL
// log.SetLevel(log.InfoLevel)
log.SetLevel(log.TraceLevel)
```

### SET FORMAT

```go
// SET FORMAT
log.SetFormatter(&log.TextFormatter{})
// log.SetFormatter(&log.JSONFormatter{})
```

### SET OUTPUT (DEFAULT stderr)

Output to stdout instead of the default stderr,
  
```go
// SET OUTPUT (DEFAULT stderr)
log.SetOutput(os.Stdout)
```

### LOGGING TO FILE (APPEND) (io.Writer)

```go
// LOGGING TO FILE (APPEND) (io.Writer)
myfile, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err == nil {
    log.SetOutput(myfile)
} else {
    log.Info("Failed to log to file, using default stderr")
}
// don't forget to close it
defer myfile.Close()
```

### NORMAL LOGGING

```go
// NORMAL LOGGING
//log.Panic("I'm bailing. Calls panic() after logging.")
//log.Fatal("Bye. Calls os.Exit(1) after logging.")
log.Error("Something failed but I'm not quitting.")
log.Warn("You should probably take a look at this.")
log.Info("Something noteworthy happened!")
log.Debug("Useful debugging information.")
log.Trace("Something very low level.")
```

### NORMAL LOGGING WITH FORMATTING

```go
// NORMAL LOGGING WITH FORMATTING
name := "jeff"
s := fmt.Sprintf("This is from %s", name)
log.Info(s)
```

### USING FIELDS

```go
// USING FIELDS
log.WithFields(log.Fields{
    "animal": "cat",
}).Info("A cat appears")
```

### REUSING FIELDS

```go
// REUSING FIELDS
jeffLogger := log.WithFields(log.Fields{
    "animal": "cat",
    "other": "I also should be logged always",
})
jeffLogger.Info("I'll be logged with common and other field")
jeffLogger.Info("Me too")
```
