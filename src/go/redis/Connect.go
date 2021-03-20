package redis

import "github.com/go-redis/redis/v8"

// Connect is exported
func Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:3306",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
