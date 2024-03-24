package databases

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func (db *Databases) InitPsqlClient() *sql.DB {
	username, password, dbName, server, sslMode := db.SetPsqlEnvVars()

	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		username, password, server, dbName, sslMode)

	postgres, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Erro ao conectar-se ao postgresql: ", err)
	}

	if err := postgres.Ping(); err != nil {
		log.Fatal("Erro ao conectar-se ao postgresql: ", err)
	}

	return postgres
}

func (db *Databases) SetPsqlEnvVars() (string, string, string, string, string){
	var password string
	var server string
	var username string
	var dbName string

	server = func() string {
		if db.Env == 0 {
			return "localhost"
		}
		return "psql"
	}()

	username = func() string {
		if db.Env == 0 {
			return "art0r"
		}
		return os.Getenv("DB_USER")
	}()

	password = func() string {
		if db.Env == 0 {
			return "1329"
		}
		return os.Getenv("DB_PASSWORD")
	}()

	dbName = func() string {
		if db.Env == 0 {
			return "myapp"
		}
		return os.Getenv("DB_NAME")
	}()
	
	sslMode := "disable"

	return username, password, dbName, server, sslMode
}