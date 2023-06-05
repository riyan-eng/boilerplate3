package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnRedis() {
	addr := fmt.Sprintf("%s:%v", "localhost", 6379)
	redisDB := 0
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       redisDB,
	})
	ctx := context.Background()
	if err := Redis.Ping(ctx).Err(); err != nil {
		log.Fatalf("redis: can't ping to redis: %v", err)
	}
	log.Println("redis: connection opened")
}
