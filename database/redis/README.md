# redis example

`redis` _is an example of
a non-relational (NoSQL) database.
SET/GET a key/value pair._

Documentation and reference,

* My
  [redis-cheat-sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/redis-cheat-sheet)
  on how to get redis running on your machine.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## CREATE A REDIS DATABASE

Obviously you need a redis server running,

Check you have one running,

```bash
redis-cli ping
```

## RUN

First, make sure you have the go library

```bash
go get -u github.com/go-redis/redis
```

Then run,

```bash
go run redis.go
```

## CREATE A CONNECTION TO REDIS SERVER

```go
client := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
})

pong, err := client.Ping().Result()
fmt.Println(pong, err)
// Output: PONG <nil>
```

## SET A KEY/VALUE

```go
err = client.Set("jeff", "monkey", 0).Err()
if err != nil {
    panic(err)
}
```

## GET A KEY/VALUE

```go
val, err := client.Get("jeff").Result()
if err != nil {
    panic(err)
}
fmt.Println("jeff", val)
// Output: jeff monkey
```

## CHECK KEY/VALUE USING REDIS-CLI

Check your key value using command line,

```bash
redis-cli
get jeff
```