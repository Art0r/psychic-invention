package models

import (
	"errors"
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

func (u *UserModel) UpdateOne(sql, id, attr string) error {
	db := u.Dbs.InitPsqlClient()
	defer db.Close()

	s := fmt.Sprintf("user/update/%s", sql)
	query := utils.GetQueryAsString(s)

	stmt, err := db.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

	result, err := stmt.Exec(id, attr)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were affected")
	}

	return nil
}
