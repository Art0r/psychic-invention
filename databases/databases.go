package databases

import (
	"log"
	"os"

	"github.com/Art0r/psychic-invention/models"
)

type Databases struct {
	Env         rune
}

func (db *Databases) InitDatabases() {

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

	if err := db.CreateTables(); err != nil {
		log.Fatal("Erro ao fazer setup do banco de dados: ", err)
	}
}

func (db *Databases) SeedPsql() {
	dbPsql := db.InitPsqlClient()
	defer dbPsql.Close()
	models.CreateUser(dbPsql, models.User{Name: "Art0r", Email: "art0r@art0r.com"})
	models.CreateUser(dbPsql, models.User{Name: "Lucas", Email: "lucas@lucas.com"})
	models.CreateUser(dbPsql, models.User{Name: "Simone", Email: "simone@simone.com"})
}