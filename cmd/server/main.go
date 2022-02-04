package main

import (
	"context"
	"github.com/dihmuzikien/go-migration-example/db"
	"github.com/dihmuzikien/go-migration-example/migration"
)

func main() {
	conn, err := db.Connect(context.TODO(), db.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Name:     "helloworld",
	})
	if err != nil {
		panic(err)
	}
	_, err = migration.New(conn.DB)
	if err != nil {
		panic(err)
	}

}
