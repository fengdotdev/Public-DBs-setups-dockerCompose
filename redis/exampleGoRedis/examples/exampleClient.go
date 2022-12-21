package examples

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)


func ExampleClient() {
	err1 := godotenv.Load()
	if err1 != nil {
	log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	addr := os.Getenv("REDIS_ADDRESS")
	password := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,  // use default DB
	})

	somekey:= "marco"
	somevalue:= "polo"

	err := rdb.Set(ctx, somekey, somevalue, 0).Err()
	if err != nil {
		panic(err)
	}


	val, err := rdb.Get(ctx, somekey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1: "+ somekey, val)


	somekey2:= "flash"
	val2, err := rdb.Get(ctx, somekey2).Result()
	if err == redis.Nil {
		fmt.Println("key2 "+ somekey2 +"does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2: "+ somekey2, val2)
	}

}