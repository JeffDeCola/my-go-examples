package main

import (
	"fmt"

	redis "github.com/go-redis/redis"
)

func main() {

    // CREATE A CONNECTION
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

    // SET key value
	err = client.Set("jeff", "monkey", 0).Err()
	if err != nil {
		panic(err)
	}

    // GET key value
	val, err := client.Get("jeff").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("jeff", val)
	// Output: jeff monkey

    // Get key value with bad key - Does not exist
	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key2 does not exist

}
