# goroutines-async-channel-send-receive-waiting example

_Sending data to a goroutine via an async channel.
Both SEND and RCV will block/wait for data._

These are my 5 main example of using goroutines,

* [goroutines-async-channel-receive-no-waiting](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-async-channel-receive-no-waiting)
* **goroutines-async-channel-send-receive-waiting** <- You are here
* [goroutines-multi-core](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core)
* [goroutines-waitgroup](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-waitgroup)
* [goroutines-worker-pools](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-worker-pools)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## CHANNELS (MESSAGE PASSING)

Giving a channel `msgCh`,

Send,

```go
msgCh <- "Jeff"                             // WAITS/BLOCKS
```

Receive,

```go
say = <-msgCh                               // ALSO WAITS/BLOCKS
```

![IMAGE - goroutines-async-channel-send-receive-waiting - IMAGE](../../docs/pics/goroutines-async-channel-send-receive-waiting.jpg)

## RUN

```bash
go run goroutines-async-channel-send-receive-waiting.go
```

Simply press return to exit.

Output is always,

```txt
Sent message Jeff
                             0 Received message Jeff
Sent message Clif
                             1 Received message Clif
Sent message Jack
                             2 Received message Jack
Sent message Jill
                             3 Received message Jill
```

Try slowing down or speeding up the SEND or RCV and you will always
get the same result because both sides block/wait.
