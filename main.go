package main

import (
	"fmt"
	"log"

	"github.com/Art0r/psychic-invention/databases"
	"github.com/Art0r/psychic-invention/models"
)

/*
Ideais

- Recomendador de Mangás que funciona
- App de Relacionamentos
- planos de fundo
*/

/*

sudo -u postgres psql
CREATE USER art0r WITH PASSWORD '1329';
CREATE DATABASE myapp WITH OWNER art0r;
GRANT ALL PRIVILEGES ON DATABASE myapp TO art0r;

*/

func main() {
	dbs := databases.Databases{}
	dbs.InitDatabases()

	dbs.SeedPsql()

	switch dbs.Env {
	case 0:
		log.Print("Running on development")
	case 1:
		log.Print("Running on homologation")
	case 2:
		log.Print("Running on production")
	}

	db := dbs.InitPsqlClient()
	user, _ := models.GetUserById(db, "1")

	fmt.Println(user)
	db.Close()
	fmt.Println("-------------------------")

	db = dbs.InitPsqlClient()
	users, _ := models.GetAllUsers(db)

	fmt.Println(users)
	db.Close()
	fmt.Println("-------------------------")

	db = dbs.InitPsqlClient()
	models.UpdateUser(db, "2", "asf@asf.com", "Asf")
	user, _ = models.GetUserById(db, "2")

	fmt.Println(user)
	db.Close()
	fmt.Println("-------------------------")

	db = dbs.InitPsqlClient()
	models.DeleteUser(db, "2")
	users, _ = models.GetAllUsers(db)

	fmt.Println(users)
	db.Close()
}
