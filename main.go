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
	dbs.Init()

	switch dbs.Env {
	case 0:
		log.Print("Running on development")
	case 1:
		log.Print("Running on homologation")
	case 2:
		log.Print("Running on production")
	}

	user, _ := models.GetUserById(dbs.PsqlClient, "1")

	fmt.Println(user)
	fmt.Println("-------------------------")

	users, _ := models.GetAllUsers(dbs.PsqlClient)

	fmt.Println(users)
	fmt.Println("-------------------------")

	models.UpdateUser(dbs.PsqlClient, "2", "asf@asf.com", "Asf")
	user, _ = models.GetUserById(dbs.PsqlClient, "2")

	fmt.Println(user)
	fmt.Println("-------------------------")

	models.DeleteUser(dbs.PsqlClient, "2")
	users, _ = models.GetAllUsers(dbs.PsqlClient)

	fmt.Println(users)
}
