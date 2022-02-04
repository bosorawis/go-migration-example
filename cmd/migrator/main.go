package main

import (
	"context"
	"fmt"
	"github.com/dihmuzikien/go-migration-example/db"
	"github.com/dihmuzikien/go-migration-example/migration"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
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
	m, err := migration.New(conn.DB)
	if err != nil {
		panic(err)
	}
	go func() {
		<-c
		m.Stop()
		fmt.Println("sent stop signal")
		time.Sleep(10 * time.Second)
		os.Exit(1)
	}()

	err = m.Run()
	if err != nil {
		panic(err)
	}
}
