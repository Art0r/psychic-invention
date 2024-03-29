package main

import (
	"github.com/Art0r/psychic-invention/databases"
	"github.com/Art0r/psychic-invention/models"
	views "github.com/Art0r/psychic-invention/views"
	"github.com/gin-gonic/gin"
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
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	
	dbs := databases.Databases{}
	dbs.InitDatabases()

	userModel := models.UserModel{
		Dbs: &dbs,
	}

	userModel.SeedUsers()
	
	views.SetUsersRoutes(r, &userModel)

	r.Run(":8000")
}