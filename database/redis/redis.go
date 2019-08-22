package main

import (
	"fmt"
	"log"

	redis "github.com/go-redis/redis"
)

// Check your error
func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

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
	checkErr(err)

	// GET key value
	val, err := client.Get("jeff").Result()
	checkErr(err)
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
