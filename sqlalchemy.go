package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/slicebit/qb"
	_ "github.com/slicebit/qb/dialects/sqlite"
)

type User struct {
	ID       string `db:"id"`
	Email    string `db:"email"`
	FullName string `db:"full_name"`
	Oscars   int    `db:"oscars"`
}

func main() {

	users := qb.Table(
		"users",
		qb.Column("id", qb.Varchar().Size(40)),
		qb.Column("email", qb.Varchar()).NotNull().Unique(),
		qb.Column("full_name", qb.Varchar()).NotNull(),
		qb.Column("oscars", qb.Int()).NotNull().Default(0),
		qb.PrimaryKey("id"),
	)

	db, err := qb.New("sqlite3", "./qb_test.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	metadata := qb.MetaData()

	// add table to metadata
	metadata.AddTable(users)

	// create all tables registered to metadata
	metadata.CreateAll(db)
	defer metadata.DropAll(db) // drops all tables

	ins := qb.Insert(users).Values(map[string]interface{}{
		"id":        "b6f8bfe3-a830-441a-a097-1777e6bfae95",
		"email":     "jack@nicholson.com",
		"full_name": "Jack Nicholson",
	})

	_, err = db.Exec(ins)
	if err != nil {
		panic(err)
	}

	// find user
	var user User

	sel := qb.Select(users.C("id"), users.C("email"), users.C("full_name")).
		From(users).
		Where(users.C("id").Eq("b6f8bfe3-a830-441a-a097-1777e6bfae95"))

	err = db.Get(sel, &user)
	fmt.Printf("%+v\n", user)
}
