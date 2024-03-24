package main

import (
	"log"

	"github.com/Art0r/psychic-invention/databases"
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
}