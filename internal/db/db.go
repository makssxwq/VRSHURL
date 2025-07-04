package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"shortURL/internal/shortener"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB: 0,
})

func AddLink(url string) (string,error) {
	short := shortener.MakeShort(GetNextID())
	err := rdb.Set(ctx, short, url, 0).Err()
	return short, err
}

func GetLink(id string) (string, error) {
	return rdb.Get(ctx, id).Result()
}

func GetNextID() int64 {
	id,err := rdb.Incr(ctx, "last_id").Result()
	if err != nil {
		log.Fatal(err)
	}
	return id
}