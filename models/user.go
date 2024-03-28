package models

import (
	"reflect"

	"github.com/Art0r/psychic-invention/databases"
	utils "github.com/Art0r/psychic-invention/utils"
)

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

type UserModel struct {
	Dbs *databases.Databases
}

func (u *UserModel) SeedUsers() {
	dbPsql := u.Dbs.InitPsqlClient()
	defer dbPsql.Close()
	u.CreateUser(&User{Name: "Art0r", Email: "art0r@art0r.com"})
	u.CreateUser(&User{Name: "Lucas", Email: "lucas@lucas.com"})
	u.CreateUser(&User{Name: "Simone", Email: "simone@simone.com"})
}

func (u *UserModel) GetUserById(id string) (*User, error)       { return u.GetOne("id", id) }
func (u *UserModel) GetUserByName(name string) (*User, error)   { return u.GetOne("name", name) }
func (u *UserModel) GetUserByEmail(email string) (*User, error) { return u.GetOne("email", email) }

func (u *UserModel) UpdateUserName(id, name string) error   { return u.Update("name", id, name) }
func (u *UserModel) UpdateUserEmail(id, email string) error { return u.Update("email", id, email) }

func (u *UserModel) CreateUser(user *User) error {
	db := u.Dbs.InitPsqlClient()
	defer db.Close()

	query := utils.GetQueryAsString("user/create")

	stmt, err := db.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

	userValues := reflect.ValueOf(user).Elem()
	numberOfFields := userValues.NumField()

	values := make([]any, numberOfFields - 1)

	for i := 1; i < numberOfFields; i++ {
		values[i - 1] = userValues.Field(i).Interface()
	}

	if err := stmt.QueryRow(values...).Scan(); err != nil {
		return err
	}

	return nil
}

func (u *UserModel) DeleteUser(id string) error {
	db := u.Dbs.InitPsqlClient()
	defer db.Close()

	query := utils.GetQueryAsString("user/delete")

	if _, err := db.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (u *UserModel) GetAllUsers() ([]User, error) {
	db := u.Dbs.InitPsqlClient()
	defer db.Close()

	var users []User

	query := utils.GetQueryAsString("all")

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, err
}
