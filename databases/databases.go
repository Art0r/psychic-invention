package databases

import (
	"log"
	"os")

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