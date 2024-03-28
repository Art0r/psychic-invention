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
	userModel := models.UserModel{
		Dbs: &dbs,
	}

	userModel.SeedUsers()

	switch dbs.Env {
	case 0:
		log.Print("Running on development")
	case 1:
		log.Print("Running on homologation")
	case 2:
		log.Print("Running on production")
	}

	user, _ := userModel.GetUserById("1")
	fmt.Println(user)

	fmt.Println("-------------------------")

	users, _ := userModel.GetAllUsers()
	fmt.Println(users)

	fmt.Println("-------------------------")

	userModel.UpdateUserEmail("2", "asf@asf.com")
	user, _ = userModel.GetUserById("2")
	fmt.Println(user)

	fmt.Println("-------------------------")

	userModel.UpdateUserName("2", "Asf")
	user, _ = userModel.GetUserById("2")
	fmt.Println(user)

	fmt.Println("-------------------------")

	userModel.DeleteUser("2")
	users, _ = userModel.GetAllUsers()
	fmt.Println(users)

	fmt.Println("-------------------------")

	user, _ = userModel.GetUserByName("Art0r")
	fmt.Println(user)

	fmt.Println("-------------------------")

	user, _ = userModel.GetUserByEmail("art0r@art0r.com")
	fmt.Println(user)

}
