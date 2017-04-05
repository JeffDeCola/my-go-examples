# goroutines-channels-select example

`goroutines-channels-select`  _is an example of concurrency and message passing 
via channels in go._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## GOROUTINE (CONCURRENCY)

Start a thread using a goroutine,

```go
go function()
go foo.bar()
```

## CHANNELS (MESAGE PASSING)

You can pass a message over a channel to a goroutine,

```go
hi := Hello{}
ch := make(chan string)
go hi.Say(ch)
time.Sleep(5 * time.Second)
ch <- "Jeff"
```

## SELECT (WAIT ON CHANNELS)

The select statement lets a goroutine wait on
multiple channels. A select blocks until one
of its cases can run, then it executes that case.

```go
for {
    select {
    case foo := <-ch1:
        fmt.Printf("do something")
    case bar := <-ch2:
         fmt.Printf("do something else")
    }
}
```