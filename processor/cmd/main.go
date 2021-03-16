package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dnlo/struct2csv"
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
	info := fmt.Sprintf("%s:%s@/%s?parseTime=true", usrenv, passenv, dbenv)
	conn, err := sql.Open("mysql", info)
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}
	return mysql.NewMySQL(conn)
}

func genCSV(ls []processor.ServiceLatencies) {
	w := struct2csv.NewWriter(os.Stdout)
	err := w.WriteStructs(ls)
	if err != nil {
		log.Fatal(err)
	}
}
func genReportCSV(ls []processor.ReportRow) {
	w := struct2csv.NewWriter(os.Stdout)
	err := w.WriteStructs(ls)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db := connectDB()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	err := processor.InsertBatch(ctx, db, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	//genCSV(as)
	//genReportCSV(as)
	//for _, a := range as {
	//	fmt.Printf("%+v\n", a)
	//}
	//for _, l := range logs {
	//	err := db.InsertRecord(ctx, &l)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
}
