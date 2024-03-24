package models

import (
	"database/sql" 	
utils "github.com/Art0r/psychic-invention/utils"
)

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func (user *User) Create(db *sql.DB) error {

	query := utils.GetQueryAsString("create_user")

	if _, err := db.Exec(query, user.Name, user.Email); err != nil {
		return err
	}

	return nil
}
