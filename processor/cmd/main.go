package main

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rschio/logprocess/processor"
	"github.com/rschio/logprocess/processor/storage/mysql"
)

var (
	usrenv  = os.Getenv("MYSQL_USER")
	passenv = os.Getenv("MYSQL_PASSWORD")
	dbenv   = os.Getenv("MYSQL_DATABASE")
)

func connectDB() processor.Storage {
	info := fmt.Sprintf("%s:%s@/%s", usrenv, passenv, dbenv)
	conn, err := sql.Open("mysql", info)
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}
	return mysql.NewMySQL(conn)
}

func main() {
	db := connectDB()
	sc := bufio.NewScanner(os.Stdin)
	logs := make([]processor.Record, 0, 200)
	for sc.Scan() {
		jsonlog := sc.Bytes()
		l := processor.Record{}
		json.Unmarshal(jsonlog, &l)
		if err := processor.ValidRecord(&l); err != nil {
			log.Println(err)
		} else {
			logs = append(logs, l)
		}
	}
	if err := sc.Err(); err != nil {
		log.Println(err)
	}
	ctx := context.Background()
	for _, l := range logs {
		err := db.InsertRecord(ctx, &l)
		if err != nil {
			log.Fatal(err)
		}
	}
}
