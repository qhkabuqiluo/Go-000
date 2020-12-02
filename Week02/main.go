package main

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

type User struct {
	ID   string
	NAME string
}

func main() {
	user, err := Biz()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("query no found: %+v\n", err)
			returngo
		}
		log.Printf("query failed: %+v\n", err)
		return
	}
	log.Printf("user info: %+v\n", user)
}
func Biz() (*User, error) {
	return Service()
}
func Service() (*User, error) {
	user, err := Dao("1")
	return user, errors.Wrap(err, "query error")
}
func Dao(id string) (*User, error) {
	err := sql.ErrNoRows
	sqlStr := "select * from user where id = ?"
	if err == sql.ErrNoRows {
		return nil, errors.WithMessagef(err, "query sql is %s, query params is %s", sqlStr, id)
	}
	return &User{}, err
}
