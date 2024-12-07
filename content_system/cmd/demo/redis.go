package demo

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main1(){
	rdb := connRdb()
	ctx := context.Background()
	err := rdb.Set(ctx, "session_id:admin", "session_id", 5*time.Second).Err()
	if err != nil{
		panic(err)
	}

	session_id, err := rdb.Get(ctx, "session_id:admin").Result()
	if err != nil && err != redis.Nil{
		panic(err)
	}
	fmt.Println(session_id)
}

func connRdb() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil{
		panic(err)
	}

	

	return rdb
}
