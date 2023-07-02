package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	initDbConnOnce sync.Once
	dbConn         *DBConn
)

type DBConn struct {
	dialect string
	url     string
	user    string
	pass    string
}

func initDatabaseConnection() {
	fmt.Println("DB connection initiation process started")
	time.Sleep(time.Second)
	dbConn = &DBConn{
		dialect: "localhost:3066",
		url:     "db_url",
		user:    "root",
		pass:    "pass",
	}
	fmt.Println("DB connection initiation process finished successfully")
}

func getDBConn() *DBConn {
	initDbConnOnce.Do(initDatabaseConnection)
	return dbConn
}

func main() {
	conn := getDBConn()
	conn = getDBConn()
	conn = getDBConn()
	_ = conn
}
