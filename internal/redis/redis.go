package redis

import "github.com/redis/go-redis/v9"

func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{})
	if(client==nil){
		panic("Cannot connect to Redis")
	}
	return client
}