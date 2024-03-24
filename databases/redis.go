package databases

import (
	"log"

	"github.com/go-redis/redis"
)

func (db *Databases) InitRedisClient() *redis.Client {
	addr := db.GetRedisAddr()

	options := &redis.Options{
		Addr: addr,
		DB:   0,
	}
	client := redis.NewClient(options)

	if _, err := client.Ping().Result(); err != nil {
		log.Fatal("Erro ao conectar-se ao Redis: ", err)
		return nil
	}

	return client
}

func (db *Databases) GetRedisAddr() string {
	var addr string

	addr = "redis:6379"

	if db.Env == 0 {
		addr = "localhost:6379"
	}

	return addr
}
