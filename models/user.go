package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Art0r/psychic-invention/databases"
	utils "github.com/Art0r/psychic-invention/utils"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

type User struct {
	ID    string `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

type UserModel struct {
	Dbs *databases.Databases
}

func (u *UserModel) SeedUsers() {
	dbPsql := u.Dbs.InitPsqlClient()
	defer dbPsql.Close()

	for i := 0; i < 10; i++ {
		id := uuid.NewString()
		email := faker.Email()
		name := faker.Name()

		u.CreateUser(&User{ID: id, Name: name, Email: email})
	}
}

func (u *UserModel) GetUserById(id string) (*User, error) { return u.GetOne("id", id) }
func (u *UserModel) GetUserByName(name string) (*User, error)   { return u.GetOne("name", name) }
func (u *UserModel) GetUserByEmail(email string) (*User, error) { return u.GetOne("email", email) }

func (u *UserModel) UpdateUser(id string, columns map[string]string) error {
	db := u.Dbs.InitPsqlClient()
	defer db.Close()

	var fields []string
	var values []interface{}
	values = append(values, id)
	i := 2
	for index, value := range columns {
		v := fmt.Sprintf("%s = $%d", index, i)
		fields = append(fields, v)
		values = append(values, value)
		i++
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d", strings.Join(fields, ", "), 1)

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(values...); err != nil {
		return err
	}

	return nil
}

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

	values := make([]any, numberOfFields)

	for i := 0; i < numberOfFields; i++ {
		values[i] = userValues.Field(i).Interface()
	}

	if _, err := stmt.Exec(values...); err != nil {
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

func (u *UserModel) GetAllUsers() ([]*User, error) {
	db := u.Dbs.InitPsqlClient()
	defer db.Close()

	var users []*User

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

		users = append(users, &user)
	}

	return users, err
}
