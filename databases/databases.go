package databases

import (
	"database/sql"
	"log"
	"os"

	"github.com/Art0r/psychic-invention/models"
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

	if err := db.CreateTables(); err != nil {
		log.Fatal("Erro ao fazer setup do banco de dados: ", err)
	}

	u1 := models.User{Name: "Art", Email: "art@art.com"}
	u2 := models.User{Name: "Art0r", Email: "art0r@art0r.com"}
	
	u1.Create(db.PsqlClient)
	u2.Create(db.PsqlClient)
}
