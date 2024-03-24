package databases

import (
	"database/sql"
	"os"

	"github.com/go-redis/redis"
)

type Databases struct {
	Env         rune
	RedisClient *redis.Client
	PsqlClient  *sql.DB
}

func (db *Databases) Init() {

	env := os.Getenv("ENV")

	switch env {
	case "0":
		db.Env = 0
	case "1":
		db.Env = 1
	case "2":
		db.Env = 2
	default:
		db.Env = 0
	}

	db.RedisClient = db.InitRedisClient()
	db.PsqlClient = db.InitPsqlClient()
}
