# goroutines-async-channel-receive-no-waiting

_A goroutine asynchronously sending data (via a channel buffer) and a goroutine
receiving that data, using the latest (if available) and does not wait._

These are my 5 main example of using goroutines,

* **goroutines-async-channel-receive-no-waiting** <- You are here
* [goroutines-async-channel-send-receive-waiting](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-async-channel-send-receive-waiting)
* [goroutines-multi-core](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core)
* [goroutines-waitgroup](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-waitgroup)
* [goroutines-worker-pools](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-worker-pools)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## ASYNCHRONOUS CHANNEL

`sendData()` will,

* Send data every x seconds to the channel buffer

`rcvData()` will,

* Receive the latest data every x seconds from the channel buffer (if available)
* If there are no data, just use the previous data

![IMAGE - goroutines-async-channel-receive-no-waiting - IMAGE](../../docs/pics/goroutines/goroutines-async-channel-receive-no-waiting.jpg)

## rcvData() FUNCTION - NO WAITING

Since the channel is asynchronous, the channel buffer could be empty or full
depending on how much faster or slower the `sendData()` goroutine is.
Hence, **we don't want the `rvcData()` function to wait.**

To solve the problem of the `rvcData()` function waiting around,

* If there is something in channel
  * Read and continue reading until empty channel (hence, get latest one)
  * break out of loop
* If there is nothing in channel
  * default and break out of the loop

```go
for {
    select {
    case newVal := <-rcvCh:
        data = newVal
        fmt.Printf("%40d - Received data %v\n", counter, data)
        continue
    default:
    }
    break
}
```

## RUN

```bash
go run goroutines-async-channel-receive-no-waiting.go
```

Simply press return to exit.

Example output when receiver faster than sender,

```txt
sendData() will send data every 5 seconds
rcvData() will receive data every 2 seconds

                                       1 - START RECEIVING DATA
                                       1 - Using data empty
   1 - START SENDING DATA
   1 - Sent data 1
                                       2 - START RECEIVING DATA
                                       2 - Received data 1
                                       2 - Using data 1
                                       3 - START RECEIVING DATA
                                       3 - Using data 1
   2 - START SENDING DATA
   2 - Sent data 2
                                       4 - START RECEIVING DATA
                                       4 - Received data 2
                                       4 - Using data 2
                                       5 - START RECEIVING DATA
                                       5 - Using data 2
                                       6 - START RECEIVING DATA
                                       6 - Using data 2
```

Try slowing down or speeding up the goroutines by changing,

```go
const sendSpeedSeconds = X
const rcvSpeedSeconds = X
```
