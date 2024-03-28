package models

import (
	"fmt"
	utils "github.com/Art0r/psychic-invention/utils"
)


func (u *UserModel) GetOne(sql, attr string) (*User, error) {
	db := u.Dbs.InitPsqlClient()
	defer db.Close()

	var user User

	s := fmt.Sprintf("user/read/%s", sql)
	query := utils.GetQueryAsString(s)

	stmt, err := db.Prepare(query)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    err = stmt.QueryRow(attr).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, err
    }


	return &user, err
}

func (u *UserModel) Update(sql, id, attr string) error {
	db := u.Dbs.InitPsqlClient()
	defer db.Close()

	s := fmt.Sprintf("user/update/%s", sql)
	query := utils.GetQueryAsString(s)

	stmt, err := db.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

	if err := stmt.QueryRow(id, attr).Scan(); err != nil {
		return err
	}

	return nil
}
