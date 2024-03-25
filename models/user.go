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

func CreateUser(db *sql.DB, user User) error {

	query := utils.GetQueryAsString("user/create")

	if _, err := db.Exec(query, user.Name, user.Email); err != nil {
		return err
	}

	return nil
}

func UpdateUser(db *sql.DB, id string, email string, name string) error {
	
	query := utils.GetQueryAsString("user/update")

	if _, err := db.Exec(query, id, email, name); err != nil {
		return err
	}

	return nil
}

func DeleteUser(db *sql.DB, id string) error {
		
	query := utils.GetQueryAsString("user/delete")

	if _, err := db.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func GetUserById(db *sql.DB, id string) (User, error) {
	var user User

	query := utils.GetQueryAsString("user/get_by_id")

	row := db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)

	return user, err
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	var users []User

	query := utils.GetQueryAsString("user/get_all")

	rows, err := db.Query(query)
	
	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, err
}
