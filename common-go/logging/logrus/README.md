# logrus

_Logging using `logrus` package._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/logrus#overview)
  * [SET LOG LEVEL, FORMAT AND OUTPUT](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/logrus#set-log-level-format-and-output)
  * [LOGGING](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/logrus#logging)
  * [FIELDS](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/logrus#fields)
  * [LOGGING TO A FILE](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/logrus#logging-to-a-file)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/logrus#prerequisites)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/logrus#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/logrus#test)
  
Documentation and references,

* Refer to the
  [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
  package for more info
* Refer to my
  [flag](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flag)
  example
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Logrus has seven levels in this order,

* log.**Panic**("I'm bailing. Calls panic() after logging.")
* log.**Error**("Something failed but I'm not quitting.")
* log.**Warn**("You should probably take a look at this.")
* log.**Info**("Something noteworthy happened!")
* log.**Debug**("Useful debugging information.")
* log.**Trace**("Something very low level.")

But I like to keep it simple,

* **Fatal**
* **Error**
* **Info**
* **Trace**

### SET LOG LEVEL, FORMAT AND OUTPUT

You got three decisions to make,

```go
// SET LOG LEVEL TO TRACE
log.SetLevel(log.TraceLevel)
```

```go
// SET FORMAT
log.SetFormatter(&log.TextFormatter{})
```

Set output to stdout instead of the default stderr,
  
```go
// SET OUTPUT (DEFAULT stderr)
log.SetOutput(os.Stdout)
```

### LOGGING

```go
// LOGGING
log.Error("Something failed but I'm not quitting.")
log.Info("Something noteworthy happened!")
log.Trace("Something very low level.")
```

```go
// WITH FORMATTING
name := "jeff"
log.Infof("This is from %s", name)
```

### FIELDS

```go
log.WithFields(log.Fields{
    "animal": "cat",
}).Trace("What animal is it?")
```

Reusing fields,

```go
// REUSING FIELDS
jeffLogger := log.WithFields(log.Fields{
    "animal": "cat",
    "color":  "grey",
})

jeffLogger.Trace("Using the animal and color field")
jeffLogger.Trace("Me too")
```

### LOGGING TO A FILE

```go
// LOGGING TO FILE (APPEND) (io.Writer)
myfile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err == nil {
    log.SetOutput(myfile)
} else {
    return nil, fmt.Errorf("failed to log to file, using default stderr: %w", err)

}
// don't forget to close it
defer myfile.Close()
```

## PREREQUISITES

Need,

```txt
github.com/sirupsen/logrus
```

Run,

```bash
go mod tidy
```

## RUN

Run with various log levels,

```bash
go run logrus.go
go run logrus.go -loglevel error
go run logrus.go -loglevel info
go run logrus.go -loglevel trace
```

For trace, the output should be,

```bash
ERRO[0000] Something failed but I'm not quitting.       
INFO[0000] Something noteworthy happened!               
TRAC[0000] Something very low level.                    
INFO[0000] This is from jeff                            
TRAC[0000] What animal is it?                            animal=cat
TRAC[0000] Using the animal and color field              animal=cat color=grey
TRAC[0000] Me too                                        animal=cat color=grey
```

An example of an error,

```bash
go run logrus.go -loglevel badinput
```

Log to a logfile instead of stdout,

```bash
go run logrus.go -loglevel trace -logfile logfile.log
cat logfile.log
```

## TEST

To create _test files,

```bash
gotests -w -all logrus.go
```

To unit test the code,

```bash
go test -cover ./... 
```
